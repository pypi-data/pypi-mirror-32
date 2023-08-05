'''
    Copyright (c) 2016-2017 Wind River Systems, Inc.
    
    Licensed under the Apache License, Version 2.0 (the "License");
    you may not use this file except in compliance with the License.
    You may obtain a copy of the License at:
    http://www.apache.org/licenses/LICENSE-2.0
    
    Unless required by applicable law or agreed to in writing, software  distributed
    under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES
    OR CONDITIONS OF ANY KIND, either express or implied.
'''

"""
This module contains the Relay class which is a secure way to pipe data to a
local socket connection. This is useful for Telnet which is not secure by
default.
"""

import logging
import random
import select
import socket
import ssl
import threading
import time
import sys
# -------------------------------------------------------------------
# Note: when using a proxy server, the socket class is overlayed with
# pysocks class.  Keep around a local copy so that local socket
# connections don't use the proxy
# -------------------------------------------------------------------
non_proxy_socket = None

# yocto supports websockets, not websocket, so check for that
try:
    import websocket
except ImportError:
    import websockets as websocket

CONNECT_MSG = "CONNECTED-129812"

class Relay(object):
    """
    Class for establishing a secure pipe between a cloud based websocket and a
    local socket. This is useful for things like Telnet which are not secure to
    use remotely.
    """

    def __init__(self, wsock_host, sock_host, sock_port, secure=True,
                 log=None, local_socket=None, reconnect=False):
        """
        Initialize a relay object for piping data between a websocket and a
        local socket
        """

        self.wsock_host = wsock_host
        self.sock_host = sock_host
        self.sock_port = sock_port
        self.secure = secure
        self.log = log
        self.proxy = None
        self.log_name = "Relay:{}:{}({:0>5})".format(self.sock_host,
                                                    self.sock_port,
                                                    random.randint(0,99999))
        self.reconnect = reconnect
        if self.log is None:
            self.logger = logging.getLogger(self.log_name)
            log_handler = logging.StreamHandler()
            #log_formatter = logging.Formatter(constants.LOG_FORMAT, datefmt=constants.LOG_TIME_FORMAT)
            #log_handler.setFormatter(log_formatter)
            self.logger.addHandler(log_handler)
            self.logger.setLevel(logging.DEBUG)
            self.log = self.logger.log

        self.running = False
        self.thread = None
        self.ws_thread = None
        self.lsock = None
        self.wsock = None
        self.lconnect = 0

    def _connect_local(self):
        ret = False
        try:
            # check for proxy.  If not proxy, this
            # is None.
            if non_proxy_socket:
                self.lsock = non_proxy_socket(socket.AF_INET,
                                       socket.SOCK_STREAM)
            else:
                self.lsock = socket.socket(socket.AF_INET,
                                       socket.SOCK_STREAM)

            self.lsock.setsockopt(socket.IPPROTO_TCP, socket.TCP_NODELAY, True)
            self.lsock.connect((self.sock_host,
                                self.sock_port))
            self.lsock.setblocking(0)
        except socket.error as err:
            self.running = False
            ret = True
            self.log(logging.ERROR, "{} Failed to open local socket.".format(self.log_name))
            self.log(logging.ERROR, "Reason: {} ".format(str(err)))
        return ret

    def _on_local_message(self):
        """
        Main loop that pipes all data from one socket to the next. The
        websocket connection is established first and has its own
        callback, so this is where the local socket will be handled.
        """
        # ws data must be in binary format.  The websocket lib uses
        # this op code
        op_binary = 0x2
        while self.running is True:
            if self.lsock:
                socket_list = [self.lsock]
                read_sockets, write_sockets, _es = select.select(socket_list, [], [], 1)
                if len(read_sockets):
                    try:
                        data = self.lsock.recv(4096)
                    except:
                        # during a close a read might return a EBADF,
                        # that is ok, pass it don't dump an exception
                        pass
                    if data:
                        self.log(logging.DEBUG, "_on_local_message: send {} -> ws".format(len(data)))
                        try:
                            self.wsock.send(data, opcode=op_binary)
                        except websocket.WebSocketConnectionClosedException:
                            self.log(logging.ERROR, "Websocket closed")
                            break
                    else:
                        self.log(logging.INFO, "{}: Received NULL from local socket".format(self.log_name))
                        if self.reconnect and self.running:
                            self.log(logging.INFO, "Reconnecting local socket")
                            time.sleep(2)
                            self._connect_local()
                        else:
                            self.running = False
                            break
            else:
                time.sleep(1)
        if self.lsock:
            self.lsock.close()
            self.lsock = None
        if self.wsock:
            self.wsock.close()
            self.wsock = None
        self.log(logging.INFO, "{} - Sockets Closed".format(self.log_name))

    def _on_open(self, ws):
        self.log(logging.INFO, "_on_open: starting thread loop")
        self.thread = threading.Thread(target=self._on_local_message)
        self.thread.start()

    def _on_message(self, ws, data):

        if data:
            if data == CONNECT_MSG:
                # If the local socket has not been established yet,
                # and we have received the connection string, start
                # local socket.
                self._connect_local()
                self.lconnect = 1;
                self.log(logging.DEBUG, "{} Local socket opened".format(self.log_name))
            else:
                # send to local socket
                self.log(logging.DEBUG, "_on_message: send {} -> local socket".format(len(data)))

                # py3 data of type string needs to be byte encoded
                if isinstance(data, str) and sys.version_info[0] > 2:
                    data = bytes(data, 'utf-8')
                self.lsock.send(data)

    def _on_error(self, ws, exception):
        self.log(logging.ERROR, "_on_error: exception {}".format(str(exception)))
        if self.lsock:
            self.lsock.close()
        if self.wsock:
            self.wsock.close()
        self.stop()

    def _on_close(self, ws):
        self.log(logging.INFO,"_on_close: websocket closed")
        if self.lsock:
            self.lsock.close()
        self.running = False

    def start(self):
        """
        Establish the websocket connection and start the main loop
        """

        if not self.running:
            self.running = True
            sslopt = {}
            if not self.secure:
                sslopt["cert_reqs"] = ssl.CERT_NONE
            self.wsock = websocket.WebSocketApp(
                    self.wsock_host,
                    on_message=self._on_message,
                    on_error=self._on_error,
                    on_close=self._on_close,
                    on_open=self._on_open)
            kwargs = {'sslopt': sslopt}
            if self.proxy:
                self.log(logging.DEBUG, "start:self.proxy={} ".format(self.proxy)),
                kwargs['http_proxy_host'] = self.proxy.host
                kwargs['http_proxy_port'] = self.proxy.port
            self.ws_thread = threading.Thread(target=self.wsock.run_forever, kwargs=kwargs)
            self.ws_thread.start()
        else:
            raise RuntimeError("{} - Already running!".format(self.log_name))

    def stop(self):
        """
        Stop piping data between the two connections and stop the loop thread
        """

        self.log(logging.INFO, "{} Stopping".format(self.log_name))
        self.running = False
        self.reconnect = False
        if self.thread:
            self.thread.join()
            self.thread = None
        if self.ws_thread:
            self.ws_thread.join()
            self.ws_thread = None


relays = []

def create_relay(url, host, port, secure=True, log_func=None, local_socket=None,
                 reconnect=False, proxy=None):
    global relays, non_proxy_socket

    non_proxy_socket = local_socket
    newrelay = Relay(url, host, port, secure=secure, log=log_func, reconnect=reconnect)
    if proxy:
        newrelay.proxy = proxy
    newrelay.start()
    relays.append(newrelay)

def stop_relays():
    global relays

    threads = []
    while relays:
        relay = relays.pop()
        thread = threading.Thread(target=relay.stop)
        thread.start()
        threads.append(thread)

    for thread in threads:
        thread.join()

