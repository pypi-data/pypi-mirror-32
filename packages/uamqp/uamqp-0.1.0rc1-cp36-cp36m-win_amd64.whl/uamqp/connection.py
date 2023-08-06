#-------------------------------------------------------------------------
# Copyright (c) Microsoft Corporation. All rights reserved.
# Licensed under the MIT License. See License.txt in the project root for
# license information.
#--------------------------------------------------------------------------

import logging
import uuid
import threading

import uamqp
from uamqp import c_uamqp
from uamqp import utils


_logger = logging.getLogger(__name__)


class Connection:
    """An AMQP Connection. A single Connection can have multiple Sessions, and
    can be shared between multiple Clients.

    :ivar max_frame_size: Maximum AMQP frame size. Default is 63488 bytes.
    :vartype max_frame_size: int
    :ivar channel_max: Maximum number of Session channels in the Connection.
    :vartype channel_max: int
    :ivar idle_timeout: Timeout in milliseconds after which the Connection will close
     if there is no further activity.
    :vartype idle_timeout: int
    :ivar properties: Connection properties.
    :vartype properties: dict

    :param hostname: The hostname of the AMQP service with which to establish
     a connection.
    :type hostname: bytes or str
    :param sasl: Authentication for the connection. If none is provided SASL Annoymous
     authentication will be used.
    :type sasl: ~uamqp.authentication.AMQPAuth
    :param container_id: The name for the client, also known as the Container ID.
     If no name is provided, a random GUID will be used.
    :type container_id: str or bytes
    :param max_frame_size: Maximum AMQP frame size. Default is 63488 bytes.
    :type max_frame_size: int
    :param channel_max: Maximum number of Session channels in the Connection.
    :type channel_max: int
    :param idle_timeout: Timeout in milliseconds after which the Connection will close
     if there is no further activity.
    :type idle_timeout: int
    :param properties: Connection properties.
    :type properties: dict
    :param remote_idle_timeout_empty_frame_send_ratio: Ratio of empty frames to
     idle time for Connections with no activity. Value must be between
     0.0 and 1.0 inclusive. Default is 0.5.
    :type remote_idle_timeout_empty_frame_send_ratio: float
    :param debug: Whether to turn on network trace logs. If `True`, trace logs
     will be logged at INFO level. Default is `False`.
    :type debug: bool
    :param encoding: The encoding to use for parameters supplied as strings.
     Default is 'UTF-8'
    :type encoding: str
    """

    def __init__(self, hostname, sasl,
                 container_id=False,
                 max_frame_size=None,
                 channel_max=None,
                 idle_timeout=None,
                 properties=None,
                 remote_idle_timeout_empty_frame_send_ratio=None,
                 debug=False,
                 encoding='UTF-8'):
        uamqp._Platform.initialize()  # pylint: disable=protected-access
        self.container_id = container_id if container_id else str(uuid.uuid4())
        self.hostname = hostname
        self.auth = sasl
        self.cbs = None
        self._conn = c_uamqp.create_connection(
            sasl.sasl_client.get_client(),
            hostname.encode(encoding) if isinstance(hostname, str) else hostname,
            self.container_id.encode(encoding) if isinstance(self.container_id, str) else self.container_id,
            self)
        self._conn.set_trace(debug)
        self._sessions = []
        self._lock = threading.Lock()
        self._state = c_uamqp.ConnectionState.UNKNOWN
        self._encoding = encoding

        if max_frame_size:
            self.max_frame_size = max_frame_size
        if channel_max:
            self.channel_max = channel_max
        if idle_timeout:
            self.idle_timeout = idle_timeout
        if properties:
            self.properties = properties
        if remote_idle_timeout_empty_frame_send_ratio:
            self._conn.remote_idle_timeout_empty_frame_send_ratio = remote_idle_timeout_empty_frame_send_ratio

    def __enter__(self):
        """Open the Connection in a context manager."""
        return self

    def __exit__(self, *args):
        """Close the Connection when exiting a context manager."""
        self.destroy()

    def _state_changed(self, previous_state, new_state):
        """Callback called whenever the underlying Connection undergoes
        a change of state. This function wraps the states as Enums for logging
        purposes.
        :param previous_state: The previous Connection state.
        :type previous_state: int
        :param new_state: The new Connection state.
        :type new_state: int
        """
        try:
            _previous_state = c_uamqp.ConnectionState(previous_state)
        except ValueError:
            _previous_state = c_uamqp.ConnectionState.UNKNOWN
        try:
            _new_state = c_uamqp.ConnectionState(new_state)
        except ValueError:
            _new_state = c_uamqp.ConnectionState.UNKNOWN
        self._state = _new_state
        _logger.debug("Connection state changed from {} to {}".format(_previous_state, _new_state))

    def destroy(self):
        """Close the connection, and close any associated
        CBS authentication session.
        """
        if self.cbs:
            self.auth.close_authenticator()
        self._conn.destroy()
        self.auth.close()
        uamqp._Platform.deinitialize()  # pylint: disable=protected-access

    def work(self):
        """Perform a single Connection iteration."""
        self._lock.acquire()
        self._conn.do_work()
        self._lock.release()

    @property
    def max_frame_size(self):
        return self._conn.max_frame_size

    @max_frame_size.setter
    def max_frame_size(self, value):
        self._conn.max_frame_size = int(value)

    @property
    def channel_max(self):
        return self._conn.channel_max

    @channel_max.setter
    def channel_max(self, value):
        self._conn.channel_max = int(value)

    @property
    def idle_timeout(self):
        return self._conn.idle_timeout

    @idle_timeout.setter
    def idle_timeout(self, value):
        self._conn.idle_timeout = int(value)

    @property
    def properties(self):
        return self._conn.properties

    @properties.setter
    def properties(self, value):
        if not isinstance(value, dict):
            raise TypeError("Connection properties must be a dictionary.")
        self._conn.properties = utils.data_factory(value, encoding=self._encoding)

    @property
    def remote_max_frame_size(self):
        return self._conn.remote_max_frame_size
