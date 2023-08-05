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


class StorageUsageDTO(object):
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
        'identifier': 'str',
        'free_space': 'str',
        'total_space': 'str',
        'used_space': 'str',
        'free_space_bytes': 'int',
        'total_space_bytes': 'int',
        'used_space_bytes': 'int',
        'utilization': 'str'
    }

    attribute_map = {
        'identifier': 'identifier',
        'free_space': 'freeSpace',
        'total_space': 'totalSpace',
        'used_space': 'usedSpace',
        'free_space_bytes': 'freeSpaceBytes',
        'total_space_bytes': 'totalSpaceBytes',
        'used_space_bytes': 'usedSpaceBytes',
        'utilization': 'utilization'
    }

    def __init__(self, identifier=None, free_space=None, total_space=None, used_space=None, free_space_bytes=None, total_space_bytes=None, used_space_bytes=None, utilization=None):
        """
        StorageUsageDTO - a model defined in Swagger
        """

        self._identifier = None
        self._free_space = None
        self._total_space = None
        self._used_space = None
        self._free_space_bytes = None
        self._total_space_bytes = None
        self._used_space_bytes = None
        self._utilization = None

        if identifier is not None:
          self.identifier = identifier
        if free_space is not None:
          self.free_space = free_space
        if total_space is not None:
          self.total_space = total_space
        if used_space is not None:
          self.used_space = used_space
        if free_space_bytes is not None:
          self.free_space_bytes = free_space_bytes
        if total_space_bytes is not None:
          self.total_space_bytes = total_space_bytes
        if used_space_bytes is not None:
          self.used_space_bytes = used_space_bytes
        if utilization is not None:
          self.utilization = utilization

    @property
    def identifier(self):
        """
        Gets the identifier of this StorageUsageDTO.
        The identifier of this storage location. The identifier will correspond to the identifier keyed in the storage configuration.

        :return: The identifier of this StorageUsageDTO.
        :rtype: str
        """
        return self._identifier

    @identifier.setter
    def identifier(self, identifier):
        """
        Sets the identifier of this StorageUsageDTO.
        The identifier of this storage location. The identifier will correspond to the identifier keyed in the storage configuration.

        :param identifier: The identifier of this StorageUsageDTO.
        :type: str
        """

        self._identifier = identifier

    @property
    def free_space(self):
        """
        Gets the free_space of this StorageUsageDTO.
        Amount of free space.

        :return: The free_space of this StorageUsageDTO.
        :rtype: str
        """
        return self._free_space

    @free_space.setter
    def free_space(self, free_space):
        """
        Sets the free_space of this StorageUsageDTO.
        Amount of free space.

        :param free_space: The free_space of this StorageUsageDTO.
        :type: str
        """

        self._free_space = free_space

    @property
    def total_space(self):
        """
        Gets the total_space of this StorageUsageDTO.
        Amount of total space.

        :return: The total_space of this StorageUsageDTO.
        :rtype: str
        """
        return self._total_space

    @total_space.setter
    def total_space(self, total_space):
        """
        Sets the total_space of this StorageUsageDTO.
        Amount of total space.

        :param total_space: The total_space of this StorageUsageDTO.
        :type: str
        """

        self._total_space = total_space

    @property
    def used_space(self):
        """
        Gets the used_space of this StorageUsageDTO.
        Amount of used space.

        :return: The used_space of this StorageUsageDTO.
        :rtype: str
        """
        return self._used_space

    @used_space.setter
    def used_space(self, used_space):
        """
        Sets the used_space of this StorageUsageDTO.
        Amount of used space.

        :param used_space: The used_space of this StorageUsageDTO.
        :type: str
        """

        self._used_space = used_space

    @property
    def free_space_bytes(self):
        """
        Gets the free_space_bytes of this StorageUsageDTO.
        The number of bytes of free space.

        :return: The free_space_bytes of this StorageUsageDTO.
        :rtype: int
        """
        return self._free_space_bytes

    @free_space_bytes.setter
    def free_space_bytes(self, free_space_bytes):
        """
        Sets the free_space_bytes of this StorageUsageDTO.
        The number of bytes of free space.

        :param free_space_bytes: The free_space_bytes of this StorageUsageDTO.
        :type: int
        """

        self._free_space_bytes = free_space_bytes

    @property
    def total_space_bytes(self):
        """
        Gets the total_space_bytes of this StorageUsageDTO.
        The number of bytes of total space.

        :return: The total_space_bytes of this StorageUsageDTO.
        :rtype: int
        """
        return self._total_space_bytes

    @total_space_bytes.setter
    def total_space_bytes(self, total_space_bytes):
        """
        Sets the total_space_bytes of this StorageUsageDTO.
        The number of bytes of total space.

        :param total_space_bytes: The total_space_bytes of this StorageUsageDTO.
        :type: int
        """

        self._total_space_bytes = total_space_bytes

    @property
    def used_space_bytes(self):
        """
        Gets the used_space_bytes of this StorageUsageDTO.
        The number of bytes of used space.

        :return: The used_space_bytes of this StorageUsageDTO.
        :rtype: int
        """
        return self._used_space_bytes

    @used_space_bytes.setter
    def used_space_bytes(self, used_space_bytes):
        """
        Sets the used_space_bytes of this StorageUsageDTO.
        The number of bytes of used space.

        :param used_space_bytes: The used_space_bytes of this StorageUsageDTO.
        :type: int
        """

        self._used_space_bytes = used_space_bytes

    @property
    def utilization(self):
        """
        Gets the utilization of this StorageUsageDTO.
        Utilization of this storage location.

        :return: The utilization of this StorageUsageDTO.
        :rtype: str
        """
        return self._utilization

    @utilization.setter
    def utilization(self, utilization):
        """
        Sets the utilization of this StorageUsageDTO.
        Utilization of this storage location.

        :param utilization: The utilization of this StorageUsageDTO.
        :type: str
        """

        self._utilization = utilization

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
        if not isinstance(other, StorageUsageDTO):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
