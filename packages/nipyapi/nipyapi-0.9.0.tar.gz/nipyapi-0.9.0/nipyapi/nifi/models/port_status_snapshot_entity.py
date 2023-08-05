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


class PortStatusSnapshotEntity(object):
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
        'port_status_snapshot': 'PortStatusSnapshotDTO',
        'can_read': 'bool'
    }

    attribute_map = {
        'id': 'id',
        'port_status_snapshot': 'portStatusSnapshot',
        'can_read': 'canRead'
    }

    def __init__(self, id=None, port_status_snapshot=None, can_read=None):
        """
        PortStatusSnapshotEntity - a model defined in Swagger
        """

        self._id = None
        self._port_status_snapshot = None
        self._can_read = None

        if id is not None:
          self.id = id
        if port_status_snapshot is not None:
          self.port_status_snapshot = port_status_snapshot
        if can_read is not None:
          self.can_read = can_read

    @property
    def id(self):
        """
        Gets the id of this PortStatusSnapshotEntity.
        The id of the port.

        :return: The id of this PortStatusSnapshotEntity.
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id):
        """
        Sets the id of this PortStatusSnapshotEntity.
        The id of the port.

        :param id: The id of this PortStatusSnapshotEntity.
        :type: str
        """

        self._id = id

    @property
    def port_status_snapshot(self):
        """
        Gets the port_status_snapshot of this PortStatusSnapshotEntity.

        :return: The port_status_snapshot of this PortStatusSnapshotEntity.
        :rtype: PortStatusSnapshotDTO
        """
        return self._port_status_snapshot

    @port_status_snapshot.setter
    def port_status_snapshot(self, port_status_snapshot):
        """
        Sets the port_status_snapshot of this PortStatusSnapshotEntity.

        :param port_status_snapshot: The port_status_snapshot of this PortStatusSnapshotEntity.
        :type: PortStatusSnapshotDTO
        """

        self._port_status_snapshot = port_status_snapshot

    @property
    def can_read(self):
        """
        Gets the can_read of this PortStatusSnapshotEntity.
        Indicates whether the user can read a given resource.

        :return: The can_read of this PortStatusSnapshotEntity.
        :rtype: bool
        """
        return self._can_read

    @can_read.setter
    def can_read(self, can_read):
        """
        Sets the can_read of this PortStatusSnapshotEntity.
        Indicates whether the user can read a given resource.

        :param can_read: The can_read of this PortStatusSnapshotEntity.
        :type: bool
        """

        self._can_read = can_read

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
        if not isinstance(other, PortStatusSnapshotEntity):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
