import logging

from pyrabbit2.api import Client
from pyrabbit2.http import HTTPError, NetworkError

from bg_utils.parser import BeerGardenSchemaParser


class PyrabbitClient(object):
    """Class that implements a connection to RabbitMQ Management HTTP API"""

    def __init__(self, host='localhost', port=15672, user='guest', password='guest',
                 virtual_host='/'):
        self.logger = logging.getLogger(__name__)

        # Pyrabbit won't infer the default virtual host ('/'). So we need to enforce it
        self._virtual_host = virtual_host or '/'

        # The client for doing Admin things over the HTTP API
        self._client = Client("%s:%s" % (host, port), user, password)

    def is_alive(self):
        try:
            return self._client.is_alive()
        except NetworkError:
            return False

    def verify_virtual_host(self):
        try:
            return self._client.get_vhost(self._virtual_host)
        except Exception:
            self.logger.error("Error verifying virtual host %s, does it exist?", self._virtual_host)
            raise

    def get_queue_size(self, queue_name):
        """Get the number of messages in a queue.

        :param queue_name: The name of the queue
        :return: The number of messages in the queue
        """
        self.logger.debug("Getting queue Size for: %s", queue_name)
        try:
            return self._client.get_queue(self._virtual_host, queue_name).get('messages', 0)
        except HTTPError as ex:
            if ex.status == 404:
                self.logger.error("Queue '%s' could not be found", queue_name)
            else:
                self.logger.error("Could not connect to queue '%s'", queue_name)
            raise ex

    def clear_queue(self, queue_name):
        """Remove all messages in a queue.

        :param queue_name: The name of the queue
        :return:None
        """
        self.logger.info("Clearing Queue: %s", queue_name)
        queue_dictionary = self._client.get_queue(self._virtual_host, queue_name)
        number_of_messages = queue_dictionary.get('messages_ready', 0)

        while number_of_messages > 0:
            self.logger.debug("Getting the Next Message")
            messages = self._client.get_messages(self._virtual_host, queue_name, count=1,
                                                 requeue=False)
            if messages and len(messages) > 0:
                message = messages[0]
                try:
                    request = BeerGardenSchemaParser.parse_request(message['payload'],
                                                                   from_string=True)
                    self.logger.debug("Canceling Request: %s", request.id)
                    request.status = 'CANCELED'
                    request.save()
                except Exception as ex:
                    self.logger.error('Error removing message:')
                    self.logger.exception(ex)
            else:
                self.logger.debug("Race condition: The while loop thought there were "
                                  "more messages to ingest but no more messages could "
                                  "be received.")
                break

            number_of_messages -= 1

    def delete_queue(self, queue_name):
        """Actually remove a queue.

        :param queue_name: The name of the queue
        :return:
        """
        self._client.delete_queue(self._virtual_host, queue_name)

    def destroy_queue(self, queue_name, force_disconnect=False):
        """Remove all remnants of a queue.

        Ignores exceptions and ensures all aspects of the queue are deleted.

        :param queue_name: The queue name
        :param force_disconnect: Attempt to forcefully disconnect consumers of this queue
        :return:
        """
        if queue_name is None:
            return

        if force_disconnect:
            try:
                self.disconnect_consumers(queue_name)
            except HTTPError as ex:
                if ex.status != 404:
                    self.logger.exception(ex)
            except Exception as ex:
                self.logger.exception(ex)

        try:
            self.clear_queue(queue_name)
        except HTTPError as ex:
            if ex.status != 404:
                self.logger.exception(ex)
        except Exception as ex:
            self.logger.exception(ex)

        try:
            self.delete_queue(queue_name)
        except HTTPError as ex:
            if ex.status != 404:
                self.logger.exception(ex)
        except Exception as ex:
            self.logger.exception(ex)

    def disconnect_consumers(self, queue_name):
        # If there are no channels, then there is nothing to do
        channels = self._client.get_channels() or []
        for channel in channels:

            # If the channel is already gone, just return an empty response
            channel_details = self._client.get_channel(channel['name']) or {'consumer_details': []}

            for consumer_details in channel_details['consumer_details']:
                if queue_name == consumer_details['queue']['name']:
                    self._client.delete_connection(
                        consumer_details['channel_details']['connection_name']
                    )
