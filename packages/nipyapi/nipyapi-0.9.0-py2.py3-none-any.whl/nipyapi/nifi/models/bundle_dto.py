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


class BundleDTO(object):
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
        'group': 'str',
        'artifact': 'str',
        'version': 'str'
    }

    attribute_map = {
        'group': 'group',
        'artifact': 'artifact',
        'version': 'version'
    }

    def __init__(self, group=None, artifact=None, version=None):
        """
        BundleDTO - a model defined in Swagger
        """

        self._group = None
        self._artifact = None
        self._version = None

        if group is not None:
          self.group = group
        if artifact is not None:
          self.artifact = artifact
        if version is not None:
          self.version = version

    @property
    def group(self):
        """
        Gets the group of this BundleDTO.
        The group of the bundle.

        :return: The group of this BundleDTO.
        :rtype: str
        """
        return self._group

    @group.setter
    def group(self, group):
        """
        Sets the group of this BundleDTO.
        The group of the bundle.

        :param group: The group of this BundleDTO.
        :type: str
        """

        self._group = group

    @property
    def artifact(self):
        """
        Gets the artifact of this BundleDTO.
        The artifact of the bundle.

        :return: The artifact of this BundleDTO.
        :rtype: str
        """
        return self._artifact

    @artifact.setter
    def artifact(self, artifact):
        """
        Sets the artifact of this BundleDTO.
        The artifact of the bundle.

        :param artifact: The artifact of this BundleDTO.
        :type: str
        """

        self._artifact = artifact

    @property
    def version(self):
        """
        Gets the version of this BundleDTO.
        The version of the bundle.

        :return: The version of this BundleDTO.
        :rtype: str
        """
        return self._version

    @version.setter
    def version(self, version):
        """
        Sets the version of this BundleDTO.
        The version of the bundle.

        :param version: The version of this BundleDTO.
        :type: str
        """

        self._version = version

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
        if not isinstance(other, BundleDTO):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
