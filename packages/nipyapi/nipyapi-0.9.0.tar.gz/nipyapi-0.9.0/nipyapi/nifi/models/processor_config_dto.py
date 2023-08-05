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


class ProcessorConfigDTO(object):
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
        'properties': 'dict(str, str)',
        'descriptors': 'dict(str, PropertyDescriptorDTO)',
        'scheduling_period': 'str',
        'scheduling_strategy': 'str',
        'execution_node': 'str',
        'penalty_duration': 'str',
        'yield_duration': 'str',
        'bulletin_level': 'str',
        'run_duration_millis': 'int',
        'concurrently_schedulable_task_count': 'int',
        'auto_terminated_relationships': 'list[str]',
        'comments': 'str',
        'custom_ui_url': 'str',
        'loss_tolerant': 'bool',
        'annotation_data': 'str',
        'default_concurrent_tasks': 'dict(str, str)',
        'default_scheduling_period': 'dict(str, str)'
    }

    attribute_map = {
        'properties': 'properties',
        'descriptors': 'descriptors',
        'scheduling_period': 'schedulingPeriod',
        'scheduling_strategy': 'schedulingStrategy',
        'execution_node': 'executionNode',
        'penalty_duration': 'penaltyDuration',
        'yield_duration': 'yieldDuration',
        'bulletin_level': 'bulletinLevel',
        'run_duration_millis': 'runDurationMillis',
        'concurrently_schedulable_task_count': 'concurrentlySchedulableTaskCount',
        'auto_terminated_relationships': 'autoTerminatedRelationships',
        'comments': 'comments',
        'custom_ui_url': 'customUiUrl',
        'loss_tolerant': 'lossTolerant',
        'annotation_data': 'annotationData',
        'default_concurrent_tasks': 'defaultConcurrentTasks',
        'default_scheduling_period': 'defaultSchedulingPeriod'
    }

    def __init__(self, properties=None, descriptors=None, scheduling_period=None, scheduling_strategy=None, execution_node=None, penalty_duration=None, yield_duration=None, bulletin_level=None, run_duration_millis=None, concurrently_schedulable_task_count=None, auto_terminated_relationships=None, comments=None, custom_ui_url=None, loss_tolerant=None, annotation_data=None, default_concurrent_tasks=None, default_scheduling_period=None):
        """
        ProcessorConfigDTO - a model defined in Swagger
        """

        self._properties = None
        self._descriptors = None
        self._scheduling_period = None
        self._scheduling_strategy = None
        self._execution_node = None
        self._penalty_duration = None
        self._yield_duration = None
        self._bulletin_level = None
        self._run_duration_millis = None
        self._concurrently_schedulable_task_count = None
        self._auto_terminated_relationships = None
        self._comments = None
        self._custom_ui_url = None
        self._loss_tolerant = None
        self._annotation_data = None
        self._default_concurrent_tasks = None
        self._default_scheduling_period = None

        if properties is not None:
          self.properties = properties
        if descriptors is not None:
          self.descriptors = descriptors
        if scheduling_period is not None:
          self.scheduling_period = scheduling_period
        if scheduling_strategy is not None:
          self.scheduling_strategy = scheduling_strategy
        if execution_node is not None:
          self.execution_node = execution_node
        if penalty_duration is not None:
          self.penalty_duration = penalty_duration
        if yield_duration is not None:
          self.yield_duration = yield_duration
        if bulletin_level is not None:
          self.bulletin_level = bulletin_level
        if run_duration_millis is not None:
          self.run_duration_millis = run_duration_millis
        if concurrently_schedulable_task_count is not None:
          self.concurrently_schedulable_task_count = concurrently_schedulable_task_count
        if auto_terminated_relationships is not None:
          self.auto_terminated_relationships = auto_terminated_relationships
        if comments is not None:
          self.comments = comments
        if custom_ui_url is not None:
          self.custom_ui_url = custom_ui_url
        if loss_tolerant is not None:
          self.loss_tolerant = loss_tolerant
        if annotation_data is not None:
          self.annotation_data = annotation_data
        if default_concurrent_tasks is not None:
          self.default_concurrent_tasks = default_concurrent_tasks
        if default_scheduling_period is not None:
          self.default_scheduling_period = default_scheduling_period

    @property
    def properties(self):
        """
        Gets the properties of this ProcessorConfigDTO.
        The properties for the processor. Properties whose value is not set will only contain the property name.

        :return: The properties of this ProcessorConfigDTO.
        :rtype: dict(str, str)
        """
        return self._properties

    @properties.setter
    def properties(self, properties):
        """
        Sets the properties of this ProcessorConfigDTO.
        The properties for the processor. Properties whose value is not set will only contain the property name.

        :param properties: The properties of this ProcessorConfigDTO.
        :type: dict(str, str)
        """

        self._properties = properties

    @property
    def descriptors(self):
        """
        Gets the descriptors of this ProcessorConfigDTO.
        Descriptors for the processor's properties.

        :return: The descriptors of this ProcessorConfigDTO.
        :rtype: dict(str, PropertyDescriptorDTO)
        """
        return self._descriptors

    @descriptors.setter
    def descriptors(self, descriptors):
        """
        Sets the descriptors of this ProcessorConfigDTO.
        Descriptors for the processor's properties.

        :param descriptors: The descriptors of this ProcessorConfigDTO.
        :type: dict(str, PropertyDescriptorDTO)
        """

        self._descriptors = descriptors

    @property
    def scheduling_period(self):
        """
        Gets the scheduling_period of this ProcessorConfigDTO.
        The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy.

        :return: The scheduling_period of this ProcessorConfigDTO.
        :rtype: str
        """
        return self._scheduling_period

    @scheduling_period.setter
    def scheduling_period(self, scheduling_period):
        """
        Sets the scheduling_period of this ProcessorConfigDTO.
        The frequency with which to schedule the processor. The format of the value will depend on th value of schedulingStrategy.

        :param scheduling_period: The scheduling_period of this ProcessorConfigDTO.
        :type: str
        """

        self._scheduling_period = scheduling_period

    @property
    def scheduling_strategy(self):
        """
        Gets the scheduling_strategy of this ProcessorConfigDTO.
        Indcates whether the prcessor should be scheduled to run in event or timer driven mode.

        :return: The scheduling_strategy of this ProcessorConfigDTO.
        :rtype: str
        """
        return self._scheduling_strategy

    @scheduling_strategy.setter
    def scheduling_strategy(self, scheduling_strategy):
        """
        Sets the scheduling_strategy of this ProcessorConfigDTO.
        Indcates whether the prcessor should be scheduled to run in event or timer driven mode.

        :param scheduling_strategy: The scheduling_strategy of this ProcessorConfigDTO.
        :type: str
        """

        self._scheduling_strategy = scheduling_strategy

    @property
    def execution_node(self):
        """
        Gets the execution_node of this ProcessorConfigDTO.
        Indicates the node where the process will execute.

        :return: The execution_node of this ProcessorConfigDTO.
        :rtype: str
        """
        return self._execution_node

    @execution_node.setter
    def execution_node(self, execution_node):
        """
        Sets the execution_node of this ProcessorConfigDTO.
        Indicates the node where the process will execute.

        :param execution_node: The execution_node of this ProcessorConfigDTO.
        :type: str
        """

        self._execution_node = execution_node

    @property
    def penalty_duration(self):
        """
        Gets the penalty_duration of this ProcessorConfigDTO.
        The amout of time that is used when the process penalizes a flowfile.

        :return: The penalty_duration of this ProcessorConfigDTO.
        :rtype: str
        """
        return self._penalty_duration

    @penalty_duration.setter
    def penalty_duration(self, penalty_duration):
        """
        Sets the penalty_duration of this ProcessorConfigDTO.
        The amout of time that is used when the process penalizes a flowfile.

        :param penalty_duration: The penalty_duration of this ProcessorConfigDTO.
        :type: str
        """

        self._penalty_duration = penalty_duration

    @property
    def yield_duration(self):
        """
        Gets the yield_duration of this ProcessorConfigDTO.
        The amount of time that must elapse before this processor is scheduled again after yielding.

        :return: The yield_duration of this ProcessorConfigDTO.
        :rtype: str
        """
        return self._yield_duration

    @yield_duration.setter
    def yield_duration(self, yield_duration):
        """
        Sets the yield_duration of this ProcessorConfigDTO.
        The amount of time that must elapse before this processor is scheduled again after yielding.

        :param yield_duration: The yield_duration of this ProcessorConfigDTO.
        :type: str
        """

        self._yield_duration = yield_duration

    @property
    def bulletin_level(self):
        """
        Gets the bulletin_level of this ProcessorConfigDTO.
        The level at which the processor will report bulletins.

        :return: The bulletin_level of this ProcessorConfigDTO.
        :rtype: str
        """
        return self._bulletin_level

    @bulletin_level.setter
    def bulletin_level(self, bulletin_level):
        """
        Sets the bulletin_level of this ProcessorConfigDTO.
        The level at which the processor will report bulletins.

        :param bulletin_level: The bulletin_level of this ProcessorConfigDTO.
        :type: str
        """

        self._bulletin_level = bulletin_level

    @property
    def run_duration_millis(self):
        """
        Gets the run_duration_millis of this ProcessorConfigDTO.
        The run duration for the processor in milliseconds.

        :return: The run_duration_millis of this ProcessorConfigDTO.
        :rtype: int
        """
        return self._run_duration_millis

    @run_duration_millis.setter
    def run_duration_millis(self, run_duration_millis):
        """
        Sets the run_duration_millis of this ProcessorConfigDTO.
        The run duration for the processor in milliseconds.

        :param run_duration_millis: The run_duration_millis of this ProcessorConfigDTO.
        :type: int
        """

        self._run_duration_millis = run_duration_millis

    @property
    def concurrently_schedulable_task_count(self):
        """
        Gets the concurrently_schedulable_task_count of this ProcessorConfigDTO.
        The number of tasks that should be concurrently schedule for the processor. If the processor doesn't allow parallol processing then any positive input will be ignored.

        :return: The concurrently_schedulable_task_count of this ProcessorConfigDTO.
        :rtype: int
        """
        return self._concurrently_schedulable_task_count

    @concurrently_schedulable_task_count.setter
    def concurrently_schedulable_task_count(self, concurrently_schedulable_task_count):
        """
        Sets the concurrently_schedulable_task_count of this ProcessorConfigDTO.
        The number of tasks that should be concurrently schedule for the processor. If the processor doesn't allow parallol processing then any positive input will be ignored.

        :param concurrently_schedulable_task_count: The concurrently_schedulable_task_count of this ProcessorConfigDTO.
        :type: int
        """

        self._concurrently_schedulable_task_count = concurrently_schedulable_task_count

    @property
    def auto_terminated_relationships(self):
        """
        Gets the auto_terminated_relationships of this ProcessorConfigDTO.
        The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere. This property differs from the 'isAutoTerminate' property of the RelationshipDTO in that the RelationshipDTO is meant to depict the current configuration, whereas this property can be set in a DTO when updating a Processor in order to change which Relationships should be auto-terminated.

        :return: The auto_terminated_relationships of this ProcessorConfigDTO.
        :rtype: list[str]
        """
        return self._auto_terminated_relationships

    @auto_terminated_relationships.setter
    def auto_terminated_relationships(self, auto_terminated_relationships):
        """
        Sets the auto_terminated_relationships of this ProcessorConfigDTO.
        The names of all relationships that cause a flow file to be terminated if the relationship is not connected elsewhere. This property differs from the 'isAutoTerminate' property of the RelationshipDTO in that the RelationshipDTO is meant to depict the current configuration, whereas this property can be set in a DTO when updating a Processor in order to change which Relationships should be auto-terminated.

        :param auto_terminated_relationships: The auto_terminated_relationships of this ProcessorConfigDTO.
        :type: list[str]
        """

        self._auto_terminated_relationships = auto_terminated_relationships

    @property
    def comments(self):
        """
        Gets the comments of this ProcessorConfigDTO.
        The comments for the processor.

        :return: The comments of this ProcessorConfigDTO.
        :rtype: str
        """
        return self._comments

    @comments.setter
    def comments(self, comments):
        """
        Sets the comments of this ProcessorConfigDTO.
        The comments for the processor.

        :param comments: The comments of this ProcessorConfigDTO.
        :type: str
        """

        self._comments = comments

    @property
    def custom_ui_url(self):
        """
        Gets the custom_ui_url of this ProcessorConfigDTO.
        The URL for the processor's custom configuration UI if applicable.

        :return: The custom_ui_url of this ProcessorConfigDTO.
        :rtype: str
        """
        return self._custom_ui_url

    @custom_ui_url.setter
    def custom_ui_url(self, custom_ui_url):
        """
        Sets the custom_ui_url of this ProcessorConfigDTO.
        The URL for the processor's custom configuration UI if applicable.

        :param custom_ui_url: The custom_ui_url of this ProcessorConfigDTO.
        :type: str
        """

        self._custom_ui_url = custom_ui_url

    @property
    def loss_tolerant(self):
        """
        Gets the loss_tolerant of this ProcessorConfigDTO.
        Whether the processor is loss tolerant.

        :return: The loss_tolerant of this ProcessorConfigDTO.
        :rtype: bool
        """
        return self._loss_tolerant

    @loss_tolerant.setter
    def loss_tolerant(self, loss_tolerant):
        """
        Sets the loss_tolerant of this ProcessorConfigDTO.
        Whether the processor is loss tolerant.

        :param loss_tolerant: The loss_tolerant of this ProcessorConfigDTO.
        :type: bool
        """

        self._loss_tolerant = loss_tolerant

    @property
    def annotation_data(self):
        """
        Gets the annotation_data of this ProcessorConfigDTO.
        The annotation data for the processor used to relay configuration between a custom UI and the procesosr.

        :return: The annotation_data of this ProcessorConfigDTO.
        :rtype: str
        """
        return self._annotation_data

    @annotation_data.setter
    def annotation_data(self, annotation_data):
        """
        Sets the annotation_data of this ProcessorConfigDTO.
        The annotation data for the processor used to relay configuration between a custom UI and the procesosr.

        :param annotation_data: The annotation_data of this ProcessorConfigDTO.
        :type: str
        """

        self._annotation_data = annotation_data

    @property
    def default_concurrent_tasks(self):
        """
        Gets the default_concurrent_tasks of this ProcessorConfigDTO.
        Maps default values for concurrent tasks for each applicable scheduling strategy.

        :return: The default_concurrent_tasks of this ProcessorConfigDTO.
        :rtype: dict(str, str)
        """
        return self._default_concurrent_tasks

    @default_concurrent_tasks.setter
    def default_concurrent_tasks(self, default_concurrent_tasks):
        """
        Sets the default_concurrent_tasks of this ProcessorConfigDTO.
        Maps default values for concurrent tasks for each applicable scheduling strategy.

        :param default_concurrent_tasks: The default_concurrent_tasks of this ProcessorConfigDTO.
        :type: dict(str, str)
        """

        self._default_concurrent_tasks = default_concurrent_tasks

    @property
    def default_scheduling_period(self):
        """
        Gets the default_scheduling_period of this ProcessorConfigDTO.
        Maps default values for scheduling period for each applicable scheduling strategy.

        :return: The default_scheduling_period of this ProcessorConfigDTO.
        :rtype: dict(str, str)
        """
        return self._default_scheduling_period

    @default_scheduling_period.setter
    def default_scheduling_period(self, default_scheduling_period):
        """
        Sets the default_scheduling_period of this ProcessorConfigDTO.
        Maps default values for scheduling period for each applicable scheduling strategy.

        :param default_scheduling_period: The default_scheduling_period of this ProcessorConfigDTO.
        :type: dict(str, str)
        """

        self._default_scheduling_period = default_scheduling_period

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
        if not isinstance(other, ProcessorConfigDTO):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
