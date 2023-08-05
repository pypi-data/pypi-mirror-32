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


class AccessPolicyEntity(object):
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
        'revision': 'RevisionDTO',
        'id': 'str',
        'uri': 'str',
        'position': 'PositionDTO',
        'permissions': 'PermissionsDTO',
        'bulletins': 'list[BulletinEntity]',
        'generated': 'str',
        'component': 'AccessPolicyDTO'
    }

    attribute_map = {
        'revision': 'revision',
        'id': 'id',
        'uri': 'uri',
        'position': 'position',
        'permissions': 'permissions',
        'bulletins': 'bulletins',
        'generated': 'generated',
        'component': 'component'
    }

    def __init__(self, revision=None, id=None, uri=None, position=None, permissions=None, bulletins=None, generated=None, component=None):
        """
        AccessPolicyEntity - a model defined in Swagger
        """

        self._revision = None
        self._id = None
        self._uri = None
        self._position = None
        self._permissions = None
        self._bulletins = None
        self._generated = None
        self._component = None

        if revision is not None:
          self.revision = revision
        if id is not None:
          self.id = id
        if uri is not None:
          self.uri = uri
        if position is not None:
          self.position = position
        if permissions is not None:
          self.permissions = permissions
        if bulletins is not None:
          self.bulletins = bulletins
        if generated is not None:
          self.generated = generated
        if component is not None:
          self.component = component

    @property
    def revision(self):
        """
        Gets the revision of this AccessPolicyEntity.
        The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.

        :return: The revision of this AccessPolicyEntity.
        :rtype: RevisionDTO
        """
        return self._revision

    @revision.setter
    def revision(self, revision):
        """
        Sets the revision of this AccessPolicyEntity.
        The revision for this request/response. The revision is required for any mutable flow requests and is included in all responses.

        :param revision: The revision of this AccessPolicyEntity.
        :type: RevisionDTO
        """

        self._revision = revision

    @property
    def id(self):
        """
        Gets the id of this AccessPolicyEntity.
        The id of the component.

        :return: The id of this AccessPolicyEntity.
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id):
        """
        Sets the id of this AccessPolicyEntity.
        The id of the component.

        :param id: The id of this AccessPolicyEntity.
        :type: str
        """

        self._id = id

    @property
    def uri(self):
        """
        Gets the uri of this AccessPolicyEntity.
        The URI for futures requests to the component.

        :return: The uri of this AccessPolicyEntity.
        :rtype: str
        """
        return self._uri

    @uri.setter
    def uri(self, uri):
        """
        Sets the uri of this AccessPolicyEntity.
        The URI for futures requests to the component.

        :param uri: The uri of this AccessPolicyEntity.
        :type: str
        """

        self._uri = uri

    @property
    def position(self):
        """
        Gets the position of this AccessPolicyEntity.
        The position of this component in the UI if applicable.

        :return: The position of this AccessPolicyEntity.
        :rtype: PositionDTO
        """
        return self._position

    @position.setter
    def position(self, position):
        """
        Sets the position of this AccessPolicyEntity.
        The position of this component in the UI if applicable.

        :param position: The position of this AccessPolicyEntity.
        :type: PositionDTO
        """

        self._position = position

    @property
    def permissions(self):
        """
        Gets the permissions of this AccessPolicyEntity.
        The permissions for this component.

        :return: The permissions of this AccessPolicyEntity.
        :rtype: PermissionsDTO
        """
        return self._permissions

    @permissions.setter
    def permissions(self, permissions):
        """
        Sets the permissions of this AccessPolicyEntity.
        The permissions for this component.

        :param permissions: The permissions of this AccessPolicyEntity.
        :type: PermissionsDTO
        """

        self._permissions = permissions

    @property
    def bulletins(self):
        """
        Gets the bulletins of this AccessPolicyEntity.
        The bulletins for this component.

        :return: The bulletins of this AccessPolicyEntity.
        :rtype: list[BulletinEntity]
        """
        return self._bulletins

    @bulletins.setter
    def bulletins(self, bulletins):
        """
        Sets the bulletins of this AccessPolicyEntity.
        The bulletins for this component.

        :param bulletins: The bulletins of this AccessPolicyEntity.
        :type: list[BulletinEntity]
        """

        self._bulletins = bulletins

    @property
    def generated(self):
        """
        Gets the generated of this AccessPolicyEntity.
        When this content was generated.

        :return: The generated of this AccessPolicyEntity.
        :rtype: str
        """
        return self._generated

    @generated.setter
    def generated(self, generated):
        """
        Sets the generated of this AccessPolicyEntity.
        When this content was generated.

        :param generated: The generated of this AccessPolicyEntity.
        :type: str
        """

        self._generated = generated

    @property
    def component(self):
        """
        Gets the component of this AccessPolicyEntity.

        :return: The component of this AccessPolicyEntity.
        :rtype: AccessPolicyDTO
        """
        return self._component

    @component.setter
    def component(self, component):
        """
        Sets the component of this AccessPolicyEntity.

        :param component: The component of this AccessPolicyEntity.
        :type: AccessPolicyDTO
        """

        self._component = component

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
        if not isinstance(other, AccessPolicyEntity):
            return False

        return self.__dict__ == other.__dict__

    def __ne__(self, other):
        """
        Returns true if both objects are not equal
        """
        return not self == other
