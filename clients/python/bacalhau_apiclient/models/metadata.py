# coding: utf-8

"""
    Bacalhau API

    This page is the reference of the Bacalhau REST API. Project docs are available at https://docs.bacalhau.org/. Find more information about Bacalhau at https://github.com/filecoin-project/bacalhau.  # noqa: E501

    OpenAPI spec version: 0.3.18.post4
    Contact: team@bacalhau.org
    Generated by: https://github.com/swagger-api/swagger-codegen.git
"""


import pprint
import re  # noqa: F401

import six

from bacalhau_apiclient.configuration import Configuration


class Metadata(object):
    """NOTE: This class is auto generated by the swagger code generator program.

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
        'client_id': 'str',
        'created_at': 'str',
        'id': 'str'
    }

    attribute_map = {
        'client_id': 'ClientID',
        'created_at': 'CreatedAt',
        'id': 'ID'
    }

    def __init__(self, client_id=None, created_at=None, id=None, _configuration=None):  # noqa: E501
        """Metadata - a model defined in Swagger"""  # noqa: E501
        if _configuration is None:
            _configuration = Configuration()
        self._configuration = _configuration

        self._client_id = None
        self._created_at = None
        self._id = None
        self.discriminator = None

        if client_id is not None:
            self.client_id = client_id
        if created_at is not None:
            self.created_at = created_at
        if id is not None:
            self.id = id

    @property
    def client_id(self):
        """Gets the client_id of this Metadata.  # noqa: E501

        The ID of the client that created this job.  # noqa: E501

        :return: The client_id of this Metadata.  # noqa: E501
        :rtype: str
        """
        return self._client_id

    @client_id.setter
    def client_id(self, client_id):
        """Sets the client_id of this Metadata.

        The ID of the client that created this job.  # noqa: E501

        :param client_id: The client_id of this Metadata.  # noqa: E501
        :type: str
        """

        self._client_id = client_id

    @property
    def created_at(self):
        """Gets the created_at of this Metadata.  # noqa: E501

        Time the job was submitted to the bacalhau network.  # noqa: E501

        :return: The created_at of this Metadata.  # noqa: E501
        :rtype: str
        """
        return self._created_at

    @created_at.setter
    def created_at(self, created_at):
        """Sets the created_at of this Metadata.

        Time the job was submitted to the bacalhau network.  # noqa: E501

        :param created_at: The created_at of this Metadata.  # noqa: E501
        :type: str
        """

        self._created_at = created_at

    @property
    def id(self):
        """Gets the id of this Metadata.  # noqa: E501

        The unique global ID of this job in the bacalhau network.  # noqa: E501

        :return: The id of this Metadata.  # noqa: E501
        :rtype: str
        """
        return self._id

    @id.setter
    def id(self, id):
        """Sets the id of this Metadata.

        The unique global ID of this job in the bacalhau network.  # noqa: E501

        :param id: The id of this Metadata.  # noqa: E501
        :type: str
        """

        self._id = id

    def to_dict(self):
        """Returns the model properties as a dict"""
        result = {}

        for attr, _ in six.iteritems(self.swagger_types):
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
        if issubclass(Metadata, dict):
            for key, value in self.items():
                result[key] = value

        return result

    def to_str(self):
        """Returns the string representation of the model"""
        return pprint.pformat(self.to_dict())

    def __repr__(self):
        """For `print` and `pprint`"""
        return self.to_str()

    def __eq__(self, other):
        """Returns true if both objects are equal"""
        if not isinstance(other, Metadata):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, Metadata):
            return True

        return self.to_dict() != other.to_dict()
