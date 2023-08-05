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


class CopySnippetRequestEntity(object):
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
        'snippet_id': 'str',
        'origin_x': 'float',
        'origin_y': 'float'
    }

    attribute_map = {
        'snippet_id': 'snippetId',
        'origin_x': 'originX',
        'origin_y': 'originY'
    }

    def __init__(self, snippet_id=None, origin_x=None, origin_y=None):
        """
        CopySnippetRequestEntity - a model defined in Swagger
        """

        self._snippet_id = None
        self._origin_x = None
        self._origin_y = None

        if snippet_id is not None:
          self.snippet_id = snippet_id
        if origin_x is not None:
          self.origin_x = origin_x
        if origin_y is not None:
          self.origin_y = origin_y

    @property
    def snippet_id(self):
        """
        Gets the snippet_id of this CopySnippetRequestEntity.
        The identifier of the snippet.

        :return: The snippet_id of this CopySnippetRequestEntity.
        :rtype: str
        """
        return self._snippet_id

    @snippet_id.setter
    def snippet_id(self, snippet_id):
        """
        Sets the snippet_id of this CopySnippetRequestEntity.
        The identifier of the snippet.

        :param snippet_id: The snippet_id of this CopySnippetRequestEntity.
        :type: str
        """

        self._snippet_id = snippet_id

    @property
    def origin_x(self):
        """
        Gets the origin_x of this CopySnippetRequestEntity.
        The x coordinate of the origin of the bounding box where the new components will be placed.

        :return: The origin_x of this CopySnippetRequestEntity.
        :rtype: float
        """
        return self._origin_x

    @origin_x.setter
    def origin_x(self, origin_x):
        """
        Sets the origin_x of this CopySnippetRequestEntity.
        The x coordinate of the origin of the bounding box where the new components will be placed.

        :param origin_x: The origin_x of this CopySnippetRequestEntity.
        :type: float
        """

        self._origin_x = origin_x

    @property
    def origin_y(self):
        """
        Gets the origin_y of this CopySnippetRequestEntity.
        The y coordinate of the origin of the bounding box where the new components will be placed.

        :return: The origin_y of this CopySnippetRequestEntity.
        :rtype: float
        """
        return self._origin_y

    @origin_y.setter
    def origin_y(self, origin_y):
        """
        Sets the origin_y of this CopySnippetRequestEntity.
        The y coordinate of the origin of the bounding box where the new components will be placed.

        :param origin_y: The origin_y of this CopySnippetRequestEntity.
        :type: float
        """

        self._origin_y = origin_y

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
        if not isinstance(other, CopySnippetRequestEntity):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
