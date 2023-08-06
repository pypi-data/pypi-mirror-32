#-------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#--------------------------------------------------------------------------

import logging

from uamqp import c_uamqp
from uamqp import utils
from uamqp import constants


_logger = logging.getLogger(__name__)


class Message:
    """An AMQP message.

    When sending, depending on the nature of the data,
    different body encoding will be used. If the data is str or bytes,
    a single part DataBody will be sent. If the data is a list or str/bytes,
    a multipart DataBody will be sent. Any other type of list will be sent
    as a SequenceBody, where as any other type of data will be sent as
    a ValueBody. An empty payload will also be sent as a ValueBody.

    :ivar on_send_complete: A custom callback to be run on completion of
     the send operation of this message. The callback must take two parameters,
     a result (of type ~uamqp.constants.MessageSendResult) and an error (of type
     Exception). The error parameter may be None if no error ocurred or the error
     information was undetermined.
    :vartype on_send_complete: callable[~uamqp.constants.MessageSendResult, Exception]

    :param body: The data to send in the message.
    :type body: Any Python data type.
    :param properties: Properties to add to the message.
    :type properties: ~uamqp.message.MessageProperties
    :param application_properties: Service specific application properties.
    :type application_properties: dict
    :param annotations: Service specific message annotations. Keys in the dictionary
     must be ~uamqp.types.AMQPSymbol or ~uamqp.types.AMQPuLong.
    :type annotations: dict
    :param header: The message header.
    :type header: ~uamqp.message.MessageHeader
    :param msg_format: A custom message format. Default is 0.
    :type msg_format: int
    :param message: Internal only. This is used to wrap an existing message
     that has been received from an AMQP service. If specified, all other
     parameters will be ignored.
    :type message: ~uamqp.c_uamqp.cMessage
    :param encoding: The encoding to use for parameters supplied as strings.
     Default is 'UTF-8'
    :type encoding: str
    """

    def __init__(self,
                 body=None,
                 properties=None,
                 application_properties=None,
                 annotations=None,
                 header=None,
                 msg_format=None,
                 message=None,
                 encoding='UTF-8'):
        self.state = constants.MessageState.WaitingToBeSent
        self.idle_time = 0
        self._retries = 0
        self._encoding = encoding
        self.on_send_complete = None
        self.properties = None
        self.application_properties = None
        self.annotations = None
        self.header = None
        self.footer = None
        self.delivery_annotations = None

        if message:
            self._parse_message(message)
        else:
            self._message = c_uamqp.create_message()
            if isinstance(body, (bytes, str)):
                self._body = DataBody(self._message)
                self._body.append(body)
            elif isinstance(body, list) and all([isinstance(b, (bytes, str)) for b in body]):
                self._body = DataBody(self._message)
                for value in body:
                    self._body.append(value)
            elif isinstance(body, list):
                self._body = SequenceBody(self._message)
                for value in body:
                    self._body.append(value)
            else:
                self._body = ValueBody(self._message)
                self._body.set(body)
            if msg_format:
                self._message.message_format = msg_format
            self.properties = properties
            self.application_properties = application_properties
            self.annotations = annotations
            self.header = header

    def __str__(self):
        if not self._message:
            return ""
        return str(self._body)

    def _parse_message(self, message):
        """Parse a message received from an AMQP service.
        :param message: The received C message.
        :type message: ~uamqp.c_uamqp.cMessage
        """
        self._message = message
        body_type = message.body_type
        if body_type == c_uamqp.MessageBodyType.NoneType:
            self._body = None
        elif body_type == c_uamqp.MessageBodyType.DataType:
            self._body = DataBody(self._message)
        elif body_type == c_uamqp.MessageBodyType.SequenceType:
            self._body = SequenceBody(self._message)
        else:
            self._body = ValueBody(self._message)
        _props = self._message.properties
        if _props:
            self.properties = MessageProperties(properties=_props, encoding=self._encoding)
        _header = self._message.header
        if _header:
            self.header = MessageHeader(header=_header)
        _footer = self._message.footer
        if _footer:
            self.footer = _footer.map
        _app_props = self._message.application_properties
        if _app_props:
            self.application_properties = _app_props.map
        _ann = self._message.message_annotations
        if _ann:
            self.annotations = _ann.map
        _delivery_ann = self._message.delivery_annotations
        if _delivery_ann:
            self.delivery_annotations = _delivery_ann.map

    def _on_message_sent(self, result, error=None):
        """Callback run on a message send operation. If message
        has a user defined callback, it will be called here. If the result
        of the operation is failure, the message state will be reverted
        to 'pending' up to the maximum retry count.

        :param result: The result of the send operation.
        :type result: int
        :param error: An Exception if an error ocurred during the send operation.
        :type error: ~Exception
        """
        result = constants.MessageSendResult(result)
        if not error and result == constants.MessageSendResult.Error and self._retries < constants.MESSAGE_SEND_RETRIES:
            self._retries += 1
            _logger.debug("Message error, retrying. Attempts: {}".format(self._retries))
            self.state = constants.MessageState.WaitingToBeSent
        elif result == constants.MessageSendResult.Error:
            _logger.error("Message error, {} retries exhausted ({})".format(constants.MESSAGE_SEND_RETRIES, error))
            self.state = constants.MessageState.Failed
            if self.on_send_complete:
                self.on_send_complete(result, error)
        else:
            _logger.debug("Message sent: {}, {}".format(result, error))
            self.state = constants.MessageState.Complete
            if self.on_send_complete:
                self.on_send_complete(result, error)

    def get_message_encoded_size(self):
        """Pre-emptively get the size of the message once it has been encoded
        to go over the wire so we can raise an error if the message will be
        rejected for being to large.
        :returns: int
        """
        # TODO: This no longer calculates the metadata accurately.
        return c_uamqp.get_encoded_message_size(self._message)

    def get_data(self):
        """Get the body data of the message. The format may vary depending
        on the body type.
        :returns: generator
        """
        if not self._message or not self._body:
            return None
        return self._body.data

    def gather(self):
        """Return all the messages represented by this object.
        This will always be a list of a single message.
        :returns: list[~uamqp.Message]
        """
        return [self]

    def get_message(self):
        """Get the underlying C message from this object.
        :returns: ~uamqp.c_uamqp.cMessage
        """
        if not self._message:
            return None
        if self.properties:
            self._message.properties = self.properties._properties  # pylint: disable=protected-access
        if self.application_properties:
            if not isinstance(self.application_properties, dict):
                raise TypeError("Application properties must be a dictionary.")
            amqp_props = utils.data_factory(self.application_properties, encoding=self._encoding)
            self._message.application_properties = amqp_props
        if self.annotations:
            if not isinstance(self.annotations, dict):
                raise TypeError("Message annotations must be a dictionary.")
            ann_props = c_uamqp.create_message_annotations(
                utils.data_factory(self.annotations, encoding=self._encoding))
            self._message.message_annotations = ann_props
        if self.header:
            self._message.header = self.header._header  # pylint: disable=protected-access
        return self._message


class BatchMessage(Message):
    """A Batched AMQP message.

    This batch message encodes multiple message bodies into a single message
    to increase through-put over the wire. It requires server-side support
    to unpackage the batched messages and so will not be universally supported.

    :ivar on_send_complete: A custom callback to be run on completion of
     the send operation of this message. The callback must take two parameters,
     a result (of type ~uamqp.constants.MessageSendResult) and an error (of type
     Exception). The error parameter may be None if no error ocurred or the error
     information was undetermined.
    :vartype on_send_complete: callable[~uamqp.constants.MessageSendResult, Exception]
    :ivar batch_format: The is the specific message format to inform the service the
     the body should be interpreted as multiple messages. The value is 0x80013700.
    :vartype batch_format: int
    :ivar max_message_length: The maximum data size in bytes to allow in a single message.
     By default this is 256kb. If sending a single batch message, an error will be raised
     if the supplied data exceeds this maximum. If sending multiple batch messages, this
     value will be used to divide the supplied data between messages.
    :vartype max_message_length: int

    :param data: An iterable source of data, where each value will be considered the
     body of a single message in the batch.
    :type data: iterable
    :param properties: Properties to add to the message. If multiple messages are created
     these properties will be applied to each message.
    :type properties: ~uamqp.message.MessageProperties
    :param application_properties: Service specific application properties. If multiple messages
     are created these properties will be applied to each message.
    :type application_properties: dict
    :param annotations: Service specific message annotations. If multiple messages are created
     these properties will be applied to each message. Keys in the dictionary
     must be ~uamqp.types.AMQPSymbol or ~uamqp.types.AMQPuLong.
    :type annotations: dict
    :param header: The message header. This header will be applied to each message in the batch.
    :type header: ~uamqp.message.MessageHeader
    :param multi_messages: Whether to send the supplied data across multiple messages. If set to
     `False`, all the data will be sent in a single message, and an error raised if the message
     is too large. If set to `True`, the data will automatically be divided across multiple messages
     of an appropriate size. The default is `False`.
    :type multi_messages: bool
    :param encoding: The encoding to use for parameters supplied as strings.
     Default is 'UTF-8'
    :type encoding: str
    :raises: ValueError if data is sent in a single message and that message exceeds the max size.
    """

    batch_format = 0x80013700
    max_message_length = constants.MAX_MESSAGE_LENGTH_BYTES
    _size_buffer = 65000

    def __init__(self,
                 data=None,
                 properties=None,
                 application_properties=None,
                 annotations=None,
                 header=None,
                 multi_messages=False,
                 encoding='UTF-8'):
        # pylint: disable=super-init-not-called
        self._multi_messages = multi_messages
        self._body_gen = data
        self._encoding = encoding
        self.on_send_complete = None
        self.properties = properties
        self.application_properties = application_properties
        self.annotations = annotations
        self.header = header

    def _create_batch_message(self):
        """Create a ~uamqp.Message for a value supplied by the data
        generator. Applies all properties and annotations to the message.

        :returns: ~uamqp.Message
        """
        return Message(body=[],
                       properties=self.properties,
                       annotations=self.annotations,
                       msg_format=self.batch_format,
                       header=self.header,
                       encoding=self._encoding)

    def _multi_message_generator(self):
        """Generate multiple ~uamqp.Message objects from a single data
        stream that in total may exceed the maximum individual message size.
        Data will be continuously added to a single message until that message
        reaches a max allowable size, at which point it will be yielded and
        a new message will be started.

        :returns: generator[~uamqp.Message]
        """
        while True:
            new_message = self._create_batch_message()
            message_size = new_message.get_message_encoded_size() + self._size_buffer
            body_size = 0
            try:
                for data in self._body_gen:
                    message_segment = []
                    if isinstance(data, str):
                        data = data.encode(self._encoding)
                    batch_data = c_uamqp.create_data(data)
                    c_uamqp.enocde_batch_value(batch_data, message_segment)
                    combined = b"".join(message_segment)
                    body_size += len(combined)
                    if (body_size + message_size) > self.max_message_length:
                        new_message.on_send_complete = self.on_send_complete
                        yield new_message
                        raise StopIteration()
                    else:
                        new_message._body.append(combined)  # pylint: disable=protected-access
            except StopIteration:
                _logger.debug("Sent partial message.")
                continue
            else:
                new_message.on_send_complete = self.on_send_complete
                yield new_message
                _logger.debug("Sent all batched data.")
                break

    def gather(self):
        """Return all the messages represented by this object. This will convert
        the batch data into individual ~uamqp.Message objects, which may be one
        or more if multi_messages is set to `True`.

        :returns: list[~uamqp.Message]
        """
        if self._multi_messages:
            return self._multi_message_generator()

        new_message = self._create_batch_message()
        message_size = new_message.get_message_encoded_size() + self._size_buffer
        body_size = 0

        for data in self._body_gen:
            message_segment = []
            if isinstance(data, str):
                data = data.encode(self._encoding)
            batch_data = c_uamqp.create_data(data)
            c_uamqp.enocde_batch_value(batch_data, message_segment)
            combined = b"".join(message_segment)
            body_size += len(combined)
            if (body_size + message_size) > self.max_message_length:
                raise ValueError(
                    "Data set too large for a single message."
                    "Set multi_messages to True to split data across multiple messages.")
            new_message._body.append(combined)  # pylint: disable=protected-access
        new_message.on_send_complete = self.on_send_complete
        return [new_message]


class MessageProperties:
    """Message properties.
    The properties that are actually used will depend on the service implementation.
    Not all received messages will have all properties, and not all properties
    will be utilized on a sent message.

    :ivar message_id: Message-id, if set, uniquely identifies a message within the message system.
     The message producer is usually responsible for setting the message-id in such a way that it
     is assured to be globally unique. A broker MAY discard a message as a duplicate if the value
     of the message-id matches that of a previously received message sent to the same node.
    :vartype message_id: str or bytes, ~uuid.UUID, ~uamqp.types.AMQPType
    :ivar user_id: The identity of the user responsible for producing the message. The client sets
     this value, and it MAY be authenticated by intermediaries.
    :vartype user_id: str or bytes
    :ivar to: The to field identifies the node that is the intended destination of the message.
     On any given transfer this might not be the node at the receiving end of the link.
    :vartype to: str or bytes
    :ivar subject:
    :vartype subject:
    :ivar reply_to:
    :vartype reply_to:
    :ivar correlation_id:
    :vartype correlation_id:
    :ivar content_type:
    :vartype content_type:
    :ivar content_encoding:
    :vartype content_encoding:
    :ivar absolute_expiry_time:
    :vartype absolute_expiry_time:
    :ivar creation_time:
    :vartype creation_time:
    :ivar group_id:
    :vartype group_id:
    :ivar group_sequence:
    :vartype group_sequence:
    :ivar reply_to_group_id:
    :vartype reply_to_group_id:
    """

    def __init__(self,
                 message_id=None,
                 user_id=None,
                 to=None,
                 subject=None,
                 reply_to=None,
                 correlation_id=None,
                 content_type=None,
                 content_encoding=None,
                 absolute_expiry_time=None,
                 creation_time=None,
                 group_id=None,
                 group_sequence=None,
                 reply_to_group_id=None,
                 properties=None,
                 encoding='UTF-8'):
        self._properties = properties if properties else c_uamqp.cProperties()
        self._encoding = encoding
        if message_id:
            self.message_id = message_id
        if user_id:
            self.user_id = user_id
        if to:
            self.to = to
        if subject:
            self.subject = subject
        if reply_to:
            self.reply_to = reply_to
        if correlation_id:
            self.correlation_id = correlation_id
        if content_type:
            self.content_type = content_type
        if content_encoding:
            self.content_encoding = content_encoding
        if absolute_expiry_time:
            self.absolute_expiry_time = absolute_expiry_time
        if creation_time:
            self.creation_time = creation_time
        if group_id:
            self.group_id = group_id
        if group_sequence:
            self.group_sequence = group_sequence
        if reply_to_group_id:
            self.reply_to_group_id = reply_to_group_id

    @property
    def message_id(self):
        _value = self._properties.message_id
        if _value:
            return _value.value
        return None

    @message_id.setter
    def message_id(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.message_id = value

    @property
    def user_id(self):
        _value = self._properties.user_id
        if _value:
            return _value.value
        return None

    @user_id.setter
    def user_id(self, value):
        if isinstance(value, str):
            value = value.encode(self._encoding)
        elif not isinstance(value, bytes):
            raise TypeError("user_id must be bytes or str.")
        self._properties.user_id = value

    @property
    def to(self):
        _value = self._properties.to
        if _value:
            return _value.value
        return None

    @to.setter
    def to(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.to = value

    @property
    def subject(self):
        _value = self._properties.subject
        if _value:
            return _value.value
        return None

    @subject.setter
    def subject(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.subject = value

    @property
    def reply_to(self):
        _value = self._properties.reply_to
        if _value:
            return _value.value
        return None

    @reply_to.setter
    def reply_to(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.reply_to = value

    @property
    def correlation_id(self):
        _value = self._properties.correlation_id
        if _value:
            return _value.value
        return None

    @correlation_id.setter
    def correlation_id(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.correlation_id = value

    @property
    def content_type(self):
        _value = self._properties.content_type
        if _value:
            return _value.value
        return None

    @content_type.setter
    def content_type(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.content_type = value

    @property
    def content_encoding(self):
        _value = self._properties.content_encoding
        if _value:
            return _value.value
        return None

    @content_encoding.setter
    def content_encoding(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.content_encoding = value

    @property
    def absolute_expiry_time(self):
        _value = self._properties.absolute_expiry_time
        if _value:
            return _value.value
        return None

    @absolute_expiry_time.setter
    def absolute_expiry_time(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.absolute_expiry_time = value

    @property
    def creation_time(self):
        _value = self._properties.creation_time
        if _value:
            return _value.value
        return None

    @creation_time.setter
    def creation_time(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.creation_time = value

    @property
    def group_id(self):
        _value = self._properties.group_id
        if _value:
            return _value
        return None

    @group_id.setter
    def group_id(self, value):
        #value = utils.data_factory(value)
        self._properties.group_id = value

    @property
    def group_sequence(self):
        _value = self._properties.group_sequence
        if _value:
            return _value
        return None

    @group_sequence.setter
    def group_sequence(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.group_sequence = value

    @property
    def reply_to_group_id(self):
        _value = self._properties.reply_to_group_id
        if _value:
            return _value.value
        return None

    @reply_to_group_id.setter
    def reply_to_group_id(self, value):
        value = utils.data_factory(value, encoding=self._encoding)
        self._properties.reply_to_group_id = value


class MessageBody:
    """Base class for an AMQP message body. This should
    not be used directly.
    """

    def __init__(self, c_message, encoding='UTF-8'):
        self._message = c_message
        self._encoding = encoding

    def __str__(self):
        if self.type == c_uamqp.MessageBodyType.NoneType:
            return ""
        return str(self.data)

    @property
    def type(self):
        return self._message.body_type

    @property
    def data(self):
        raise NotImplementedError("Only MessageBody subclasses have data.")


class DataBody(MessageBody):
    """An AMQP message body of type Data. This represents
    a list of bytes sections.

    :ivar type: The body type. This should always be DataType
    :vartype type: ~uamqp.c_uamqp.MessageBodyType
    :ivar data: The data contained in the message body. This returns
     a generator to iterate over each section in the body, where
     each section will be a byte string.
    :vartype data: generator[bytes]
    """

    def __str__(self):
        return "".join(d.decode(self._encoding) for d in self.data)

    def __len__(self):
        return self._message.count_body_data()

    def __getitem__(self, index):
        if index >= len(self):
            raise IndexError("Index is out of range.")
        data = self._message.get_body_data(index)
        return data.value

    def append(self, data):
        """Addend a section to the body.

        :param data: The data to append.
        :type data: str or bytes.
        """
        if isinstance(data, str):
            self._message.add_body_data(data.encode(self._encoding))
        elif isinstance(data, bytes):
            self._message.add_body_data(data)

    @property
    def data(self):
        for i in range(len(self)):
            yield self._message.get_body_data(i)


class SequenceBody(MessageBody):
    """An AMQP message body of type Sequence. This represents
    a list of encoded objects.

    :ivar type: The body type. This should always be SequenceType
    :vartype type: ~uamqp.c_uamqp.MessageBodyType
    :ivar data: The data contained in the message body. This returns
     a generator to iterate over each item in the body.
    :vartype data: generator
    """

    def __len__(self):
        return self._message.count_body_sequence()

    def __getitem__(self, index):
        if index >= len(self):
            raise IndexError("Index is out of range.")
        data = self._message.get_body_sequence(index)
        return data.value

    def append(self, value):
        """Addend an item to the body. This can be any
        Python data type and it will be automatically encoded
        into an AMQP type. If a specific AMQP type is required, a
        ~uamqp.types.AMQPType can be used.

        :param data: The data to append.
        :type data: ~uamqp.types.AMQPType
        """
        value = utils.data_factory(value, encoding=self._encoding)
        self._message.add_body_sequence(value)

    @property
    def data(self):
        for i in range(len(self)):  # pylint: disable=consider-using-enumerate
            yield self[i]


class ValueBody(MessageBody):
    """An AMQP message body of type Value. This represents
    a single encoded object.

    :ivar type: The body type. This should always be ValueType
    :vartype type: ~uamqp.c_uamqp.MessageBodyType
    :ivar data: The data contained in the message body. The value
     of the encoded object
    :vartype data: object
    """

    def set(self, value):
        """Set a value as the message body. This can be any
        Python data type and it will be automatically encoded
        into an AMQP type. If a specific AMQP type is required, a
        ~uamqp.types.AMQPType can be used.

        :param data: The data to send in the body.
        :type data: ~uamqp.types.AMQPType
        """
        value = utils.data_factory(value)
        self._message.set_body_value(value)

    @property
    def data(self):
        _value = self._message.get_body_value()
        if _value:
            return _value.value
        return None


class MessageHeader:
    """The Message header. This is only used on received message, and not
    set on messages being sent. The properties set on any given message
    will depend on the Service and not all messages will have all properties.

    :ivar delivery_count: The number of unsuccessful previous attempts to deliver
     this message. If this value is non-zero it can be taken as an indication that the
     delivery might be a duplicate. On first delivery, the value is zero. It is
     incremented upon an outcome being settled at the sender, according to rules
     defined for each outcome.
    :vartype delivery_count: int
    :ivar time_to_live: Duration in milliseconds for which the message is to be considered "live".
     If this is set then a message expiration time will be computed based on the time of arrival
     at an intermediary. Messages that live longer than their expiration time will be discarded
     (or dead lettered). When a message is transmitted by an intermediary that was received
     with a ttl, the transmitted message's header SHOULD contain a ttl that is computed as the
     difference between the current time and the formerly computed message expiration time,
     i.e., the reduced ttl, so that messages will eventually die if they end up in a delivery loop.
    :vartype time_to_live: int
    :ivar durable: Durable messages MUST NOT be lost even if an intermediary is unexpectedly terminated
     and restarted. A target which is not capable of fulfilling this guarantee MUST NOT accept messages
     where the durable header is set to `True`: if the source allows the rejected outcome then the
     message SHOULD be rejected with the precondition-failed error, otherwise the link MUST be detached
     by the receiver with the same error.
    :vartype durable: bool
    :ivar first_acquirer: If this value is `True`, then this message has not been acquired
     by any other link. If this value is `False`, then this message MAY have previously
     been acquired by another link or links.
    :vartype first_acquirer: bool
    :ivar priority: This field contains the relative message priority. Higher numbers indicate higher
     priority messages. Messages with higher priorities MAY be delivered before those with lower priorities.
    :vartype priority: int

    :param header: Internal only. This is used to wrap an existing message header
     that has been received from an AMQP service.
    :type header: ~uamqp.c_uamqp.cHeader
    """

    def __init__(self, header=None):
        self._header = header if header else c_uamqp.create_header()

    @property
    def delivery_count(self):
        return self._header.delivery_count

    @property
    def time_to_live(self):
        return self._header.time_to_live

    @property
    def first_acquirer(self):
        return self._header.first_acquirer

    @property
    def durable(self):
        return self._header.durable

    @durable.setter
    def durable(self, value):
        self._header.durable = bool(value)

    @property
    def priority(self):
        return self._header.priority

    @priority.setter
    def priority(self, value):
        self._header.priority = int(value)
