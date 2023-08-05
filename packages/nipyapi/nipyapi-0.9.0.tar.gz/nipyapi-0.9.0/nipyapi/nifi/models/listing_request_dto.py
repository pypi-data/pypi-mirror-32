# coding: utf-8

"""
    NiFi Rest Api

    The Rest Api provides programmatic access to command and control a NiFi instance in real time. Start and                                              stop processors, monitor queues, query provenance data, and more. Each endpoint below includes a description,                                             definitions of the expected input and output, potential response codes, and the authorizations required                                             to invoke each service.

    OpenAPI spec version: 1.6.0
    Contact: dev@nifi.apache.org
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


from pprint import pformat
from six import iteritems
import re


class ListingRequestDTO(object):
    """
    NOTE: This class is auto generated by the swagger code generator program.
    Do not edit the class manually.
    """


    """
    Attributes:
      swagger_types (dict): The key is attribute name
                            and the value is attribute type.
      attribute_map (dict): The key is attribute name
                            and the value is json key in definition.
    """
    swagger_types = {
        'id': 'str',
        'uri': 'str',
        'submission_time': 'str',
        'last_updated': 'str',
        'percent_completed': 'int',
        'finished': 'bool',
        'failure_reason': 'str',
        'max_results': 'int',
        'state': 'str',
        'queue_size': 'QueueSizeDTO',
        'flow_file_summaries': 'list[FlowFileSummaryDTO]',
        'source_running': 'bool',
        'destination_running': 'bool'
    }

    attribute_map = {
        'id': 'id',
        'uri': 'uri',
        'submission_time': 'submissionTime',
        'last_updated': 'lastUpdated',
        'percent_completed': 'percentCompleted',
        'finished': 'finished',
        'failure_reason': 'failureReason',
        'max_results': 'maxResults',
        'state': 'state',
        'queue_size': 'queueSize',
        'flow_file_summaries': 'flowFileSummaries',
        'source_running': 'sourceRunning',
        'destination_running': 'destinationRunning'
    }

    def __init__(self, id=None, uri=None, submission_time=None, last_updated=None, percent_completed=None, finished=None, failure_reason=None, max_results=None, state=None, queue_size=None, flow_file_summaries=None, source_running=None, destination_running=None):
        """
        ListingRequestDTO - a model defined in Swagger
        """

        self._id = None
        self._uri = None
        self._submission_time = None
        self._last_updated = None
        self._percent_completed = None
        self._finished = None
        self._failure_reason = None
        self._max_results = None
        self._state = None
        self._queue_size = None
        self._flow_file_summaries = None
        self._source_running = None
        self._destination_running = None

        if id is not None:
          self.id = id
        if uri is not None:
          self.uri = uri
        if submission_time is not None:
          self.submission_time = submission_time
        if last_updated is not None:
          self.last_updated = last_updated
        if percent_completed is not None:
          self.percent_completed = percent_completed
        if finished is not None:
          self.finished = finished
        if failure_reason is not None:
          self.failure_reason = failure_reason
        if max_results is not None:
          self.max_results = max_results
        if state is not None:
          self.state = state
        if queue_size is not None:
          self.queue_size = queue_size
        if flow_file_summaries is not None:
          self.flow_file_summaries = flow_file_summaries
        if source_running is not None:
          self.source_running = source_running
        if destination_running is not None:
          self.destination_running = destination_running

    @property
    def id(self):
        """
        Gets the id of this ListingRequestDTO.
        The id for this listing request.

        :return: The id of this ListingRequestDTO.
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id):
        """
        Sets the id of this ListingRequestDTO.
        The id for this listing request.

        :param id: The id of this ListingRequestDTO.
        :type: str
        """

        self._id = id

    @property
    def uri(self):
        """
        Gets the uri of this ListingRequestDTO.
        The URI for future requests to this listing request.

        :return: The uri of this ListingRequestDTO.
        :rtype: str
        """
        return self._uri

    @uri.setter
    def uri(self, uri):
        """
        Sets the uri of this ListingRequestDTO.
        The URI for future requests to this listing request.

        :param uri: The uri of this ListingRequestDTO.
        :type: str
        """

        self._uri = uri

    @property
    def submission_time(self):
        """
        Gets the submission_time of this ListingRequestDTO.
        The timestamp when the query was submitted.

        :return: The submission_time of this ListingRequestDTO.
        :rtype: str
        """
        return self._submission_time

    @submission_time.setter
    def submission_time(self, submission_time):
        """
        Sets the submission_time of this ListingRequestDTO.
        The timestamp when the query was submitted.

        :param submission_time: The submission_time of this ListingRequestDTO.
        :type: str
        """

        self._submission_time = submission_time

    @property
    def last_updated(self):
        """
        Gets the last_updated of this ListingRequestDTO.
        The last time this listing request was updated.

        :return: The last_updated of this ListingRequestDTO.
        :rtype: str
        """
        return self._last_updated

    @last_updated.setter
    def last_updated(self, last_updated):
        """
        Sets the last_updated of this ListingRequestDTO.
        The last time this listing request was updated.

        :param last_updated: The last_updated of this ListingRequestDTO.
        :type: str
        """

        self._last_updated = last_updated

    @property
    def percent_completed(self):
        """
        Gets the percent_completed of this ListingRequestDTO.
        The current percent complete.

        :return: The percent_completed of this ListingRequestDTO.
        :rtype: int
        """
        return self._percent_completed

    @percent_completed.setter
    def percent_completed(self, percent_completed):
        """
        Sets the percent_completed of this ListingRequestDTO.
        The current percent complete.

        :param percent_completed: The percent_completed of this ListingRequestDTO.
        :type: int
        """

        self._percent_completed = percent_completed

    @property
    def finished(self):
        """
        Gets the finished of this ListingRequestDTO.
        Whether the query has finished.

        :return: The finished of this ListingRequestDTO.
        :rtype: bool
        """
        return self._finished

    @finished.setter
    def finished(self, finished):
        """
        Sets the finished of this ListingRequestDTO.
        Whether the query has finished.

        :param finished: The finished of this ListingRequestDTO.
        :type: bool
        """

        self._finished = finished

    @property
    def failure_reason(self):
        """
        Gets the failure_reason of this ListingRequestDTO.
        The reason, if any, that this listing request failed.

        :return: The failure_reason of this ListingRequestDTO.
        :rtype: str
        """
        return self._failure_reason

    @failure_reason.setter
    def failure_reason(self, failure_reason):
        """
        Sets the failure_reason of this ListingRequestDTO.
        The reason, if any, that this listing request failed.

        :param failure_reason: The failure_reason of this ListingRequestDTO.
        :type: str
        """

        self._failure_reason = failure_reason

    @property
    def max_results(self):
        """
        Gets the max_results of this ListingRequestDTO.
        The maximum number of FlowFileSummary objects to return

        :return: The max_results of this ListingRequestDTO.
        :rtype: int
        """
        return self._max_results

    @max_results.setter
    def max_results(self, max_results):
        """
        Sets the max_results of this ListingRequestDTO.
        The maximum number of FlowFileSummary objects to return

        :param max_results: The max_results of this ListingRequestDTO.
        :type: int
        """

        self._max_results = max_results

    @property
    def state(self):
        """
        Gets the state of this ListingRequestDTO.
        The current state of the listing request.

        :return: The state of this ListingRequestDTO.
        :rtype: str
        """
        return self._state

    @state.setter
    def state(self, state):
        """
        Sets the state of this ListingRequestDTO.
        The current state of the listing request.

        :param state: The state of this ListingRequestDTO.
        :type: str
        """

        self._state = state

    @property
    def queue_size(self):
        """
        Gets the queue_size of this ListingRequestDTO.
        The size of the queue

        :return: The queue_size of this ListingRequestDTO.
        :rtype: QueueSizeDTO
        """
        return self._queue_size

    @queue_size.setter
    def queue_size(self, queue_size):
        """
        Sets the queue_size of this ListingRequestDTO.
        The size of the queue

        :param queue_size: The queue_size of this ListingRequestDTO.
        :type: QueueSizeDTO
        """

        self._queue_size = queue_size

    @property
    def flow_file_summaries(self):
        """
        Gets the flow_file_summaries of this ListingRequestDTO.
        The FlowFile summaries. The summaries will be populated once the request has completed.

        :return: The flow_file_summaries of this ListingRequestDTO.
        :rtype: list[FlowFileSummaryDTO]
        """
        return self._flow_file_summaries

    @flow_file_summaries.setter
    def flow_file_summaries(self, flow_file_summaries):
        """
        Sets the flow_file_summaries of this ListingRequestDTO.
        The FlowFile summaries. The summaries will be populated once the request has completed.

        :param flow_file_summaries: The flow_file_summaries of this ListingRequestDTO.
        :type: list[FlowFileSummaryDTO]
        """

        self._flow_file_summaries = flow_file_summaries

    @property
    def source_running(self):
        """
        Gets the source_running of this ListingRequestDTO.
        Whether the source of the connection is running

        :return: The source_running of this ListingRequestDTO.
        :rtype: bool
        """
        return self._source_running

    @source_running.setter
    def source_running(self, source_running):
        """
        Sets the source_running of this ListingRequestDTO.
        Whether the source of the connection is running

        :param source_running: The source_running of this ListingRequestDTO.
        :type: bool
        """

        self._source_running = source_running

    @property
    def destination_running(self):
        """
        Gets the destination_running of this ListingRequestDTO.
        Whether the destination of the connection is running

        :return: The destination_running of this ListingRequestDTO.
        :rtype: bool
        """
        return self._destination_running

    @destination_running.setter
    def destination_running(self, destination_running):
        """
        Sets the destination_running of this ListingRequestDTO.
        Whether the destination of the connection is running

        :param destination_running: The destination_running of this ListingRequestDTO.
        :type: bool
        """

        self._destination_running = destination_running

    def to_dict(self):
        """
        Returns the model properties as a dict
        """
        result = {}

        for attr, _ in iteritems(self.swagger_types):
            value = getattr(self, attr)
            if isinstance(value, list):
                result[attr] = list(map(
                    lambda x: x.to_dict() if hasattr(x, "to_dict") else x,
                    value
                ))
            elif hasattr(value, "to_dict"):
                result[attr] = value.to_dict()
            elif isinstance(value, dict):
                result[attr] = dict(map(
                    lambda item: (item[0], item[1].to_dict())
                    if hasattr(item[1], "to_dict") else item,
                    value.items()
                ))
            else:
                result[attr] = value

        return result

    def to_str(self):
        """
        Returns the string representation of the model
        """
        return pformat(self.to_dict())

    def __repr__(self):
        """
        For `print` and `pprint`
        """
        return self.to_str()

    def __eq__(self, other):
        """
        Returns true if both objects are equal
        """
        if not isinstance(other, ListingRequestDTO):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
