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


class VariableEntity(object):
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
        'variable': 'VariableDTO',
        'can_write': 'bool'
    }

    attribute_map = {
        'variable': 'variable',
        'can_write': 'canWrite'
    }

    def __init__(self, variable=None, can_write=None):
        """
        VariableEntity - a model defined in Swagger
        """

        self._variable = None
        self._can_write = None

        if variable is not None:
          self.variable = variable
        if can_write is not None:
          self.can_write = can_write

    @property
    def variable(self):
        """
        Gets the variable of this VariableEntity.
        The variable information

        :return: The variable of this VariableEntity.
        :rtype: VariableDTO
        """
        return self._variable

    @variable.setter
    def variable(self, variable):
        """
        Sets the variable of this VariableEntity.
        The variable information

        :param variable: The variable of this VariableEntity.
        :type: VariableDTO
        """

        self._variable = variable

    @property
    def can_write(self):
        """
        Gets the can_write of this VariableEntity.
        Indicates whether the user can write a given resource.

        :return: The can_write of this VariableEntity.
        :rtype: bool
        """
        return self._can_write

    @can_write.setter
    def can_write(self, can_write):
        """
        Sets the can_write of this VariableEntity.
        Indicates whether the user can write a given resource.

        :param can_write: The can_write of this VariableEntity.
        :type: bool
        """

        self._can_write = can_write

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
        if not isinstance(other, VariableEntity):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
