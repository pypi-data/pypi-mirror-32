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


class AffectedComponentDTO(object):
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
        'process_group_id': 'str',
        'id': 'str',
        'reference_type': 'str',
        'name': 'str',
        'state': 'str',
        'active_thread_count': 'int',
        'validation_errors': 'list[str]'
    }

    attribute_map = {
        'process_group_id': 'processGroupId',
        'id': 'id',
        'reference_type': 'referenceType',
        'name': 'name',
        'state': 'state',
        'active_thread_count': 'activeThreadCount',
        'validation_errors': 'validationErrors'
    }

    def __init__(self, process_group_id=None, id=None, reference_type=None, name=None, state=None, active_thread_count=None, validation_errors=None):
        """
        AffectedComponentDTO - a model defined in Swagger
        """

        self._process_group_id = None
        self._id = None
        self._reference_type = None
        self._name = None
        self._state = None
        self._active_thread_count = None
        self._validation_errors = None

        if process_group_id is not None:
          self.process_group_id = process_group_id
        if id is not None:
          self.id = id
        if reference_type is not None:
          self.reference_type = reference_type
        if name is not None:
          self.name = name
        if state is not None:
          self.state = state
        if active_thread_count is not None:
          self.active_thread_count = active_thread_count
        if validation_errors is not None:
          self.validation_errors = validation_errors

    @property
    def process_group_id(self):
        """
        Gets the process_group_id of this AffectedComponentDTO.
        The UUID of the Process Group that this component is in

        :return: The process_group_id of this AffectedComponentDTO.
        :rtype: str
        """
        return self._process_group_id

    @process_group_id.setter
    def process_group_id(self, process_group_id):
        """
        Sets the process_group_id of this AffectedComponentDTO.
        The UUID of the Process Group that this component is in

        :param process_group_id: The process_group_id of this AffectedComponentDTO.
        :type: str
        """

        self._process_group_id = process_group_id

    @property
    def id(self):
        """
        Gets the id of this AffectedComponentDTO.
        The UUID of this component

        :return: The id of this AffectedComponentDTO.
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id):
        """
        Sets the id of this AffectedComponentDTO.
        The UUID of this component

        :param id: The id of this AffectedComponentDTO.
        :type: str
        """

        self._id = id

    @property
    def reference_type(self):
        """
        Gets the reference_type of this AffectedComponentDTO.
        The type of this component

        :return: The reference_type of this AffectedComponentDTO.
        :rtype: str
        """
        return self._reference_type

    @reference_type.setter
    def reference_type(self, reference_type):
        """
        Sets the reference_type of this AffectedComponentDTO.
        The type of this component

        :param reference_type: The reference_type of this AffectedComponentDTO.
        :type: str
        """
        allowed_values = ["PROCESSOR", "CONTROLLER_SERVICE", "INPUT_PORT", "OUTPUT_PORT", "REMOTE_INPUT_PORT", "REMOTE_OUTPUT_PORT"]
        if reference_type not in allowed_values:
            raise ValueError(
                "Invalid value for `reference_type` ({0}), must be one of {1}"
                .format(reference_type, allowed_values)
            )

        self._reference_type = reference_type

    @property
    def name(self):
        """
        Gets the name of this AffectedComponentDTO.
        The name of this component.

        :return: The name of this AffectedComponentDTO.
        :rtype: str
        """
        return self._name

    @name.setter
    def name(self, name):
        """
        Sets the name of this AffectedComponentDTO.
        The name of this component.

        :param name: The name of this AffectedComponentDTO.
        :type: str
        """

        self._name = name

    @property
    def state(self):
        """
        Gets the state of this AffectedComponentDTO.
        The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state.

        :return: The state of this AffectedComponentDTO.
        :rtype: str
        """
        return self._state

    @state.setter
    def state(self, state):
        """
        Sets the state of this AffectedComponentDTO.
        The scheduled state of a processor or reporting task referencing a controller service. If this component is another controller service, this field represents the controller service state.

        :param state: The state of this AffectedComponentDTO.
        :type: str
        """

        self._state = state

    @property
    def active_thread_count(self):
        """
        Gets the active_thread_count of this AffectedComponentDTO.
        The number of active threads for the referencing component.

        :return: The active_thread_count of this AffectedComponentDTO.
        :rtype: int
        """
        return self._active_thread_count

    @active_thread_count.setter
    def active_thread_count(self, active_thread_count):
        """
        Sets the active_thread_count of this AffectedComponentDTO.
        The number of active threads for the referencing component.

        :param active_thread_count: The active_thread_count of this AffectedComponentDTO.
        :type: int
        """

        self._active_thread_count = active_thread_count

    @property
    def validation_errors(self):
        """
        Gets the validation_errors of this AffectedComponentDTO.
        The validation errors for the component.

        :return: The validation_errors of this AffectedComponentDTO.
        :rtype: list[str]
        """
        return self._validation_errors

    @validation_errors.setter
    def validation_errors(self, validation_errors):
        """
        Sets the validation_errors of this AffectedComponentDTO.
        The validation errors for the component.

        :param validation_errors: The validation_errors of this AffectedComponentDTO.
        :type: list[str]
        """

        self._validation_errors = validation_errors

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
        if not isinstance(other, AffectedComponentDTO):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
