"""VirtuinTestViewClient is used to perform requests to view server
over AMQP (via RabbitMQ).
"""

import json
import math
from pyvirtuintestcomm.shared import openBroker


class VirtuinTestViewClient(object):
    """VirtuinTestViewClient is used to perform requests to view server
    over AMQP (via RabbitMQ).

    Attributes:
        None
    """
    def __init__(self, stationName, testName):
        self.stationName = stationName
        self.exName = '{:s}_view'.format(self.stationName)
        self.qName = testName
        self.conn = None
        self.ch = None
        self.isOpen = False
        self.retryDelay = 500

    def open(self, hostName='localhost', timeoutms=5000):
        """Open connection to RabbitMQ
        Args:
            hostName (str): Address of broker
        Returns:
            None
        """
        numAttempts = max(1, math.ceil(timeoutms / self.retryDelay))
        self.conn = openBroker(hostName, numAttempts, self.retryDelay)
        self.ch = self.conn.channel()
        self.ch.exchange_declare(exchange=self.exName, exchange_type='direct', durable=False)
        self.isOpen = True

    def close(self):
        """Close connection to RabbitMQ
        Args:
            None
        Returns:
            None
        """
        if self.isOpen:
            self.conn.close()
        self.conn = None
        self.ch = None
        self.isOpen = False

    def write(self, data, qName=None):
        """Write data to view server.
        Args:
            data (dict): Serializable data to send.
            qName (str): Name of queue to publish to.
        Returns:
            None
        """
        if not self.isOpen:
            raise Exception('Must open() before writing')
        dataBuffer = json.dumps(data)
        self.ch.basic_publish(exchange=self.exName,
                              routing_key=qName or self.qName,
                              body=dataBuffer)
