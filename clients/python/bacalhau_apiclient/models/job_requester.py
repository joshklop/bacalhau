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


class JobRequester(object):
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
        'requester_node_id': 'str',
        'requester_public_key': 'list[int]'
    }

    attribute_map = {
        'requester_node_id': 'RequesterNodeID',
        'requester_public_key': 'RequesterPublicKey'
    }

    def __init__(self, requester_node_id=None, requester_public_key=None, _configuration=None):  # noqa: E501
        """JobRequester - a model defined in Swagger"""  # noqa: E501
        if _configuration is None:
            _configuration = Configuration()
        self._configuration = _configuration

        self._requester_node_id = None
        self._requester_public_key = None
        self.discriminator = None

        if requester_node_id is not None:
            self.requester_node_id = requester_node_id
        if requester_public_key is not None:
            self.requester_public_key = requester_public_key

    @property
    def requester_node_id(self):
        """Gets the requester_node_id of this JobRequester.  # noqa: E501

        The ID of the requester node that owns this job.  # noqa: E501

        :return: The requester_node_id of this JobRequester.  # noqa: E501
        :rtype: str
        """
        return self._requester_node_id

    @requester_node_id.setter
    def requester_node_id(self, requester_node_id):
        """Sets the requester_node_id of this JobRequester.

        The ID of the requester node that owns this job.  # noqa: E501

        :param requester_node_id: The requester_node_id of this JobRequester.  # noqa: E501
        :type: str
        """

        self._requester_node_id = requester_node_id

    @property
    def requester_public_key(self):
        """Gets the requester_public_key of this JobRequester.  # noqa: E501

        The public key of the Requester node that created this job This can be used to encrypt messages back to the creator  # noqa: E501

        :return: The requester_public_key of this JobRequester.  # noqa: E501
        :rtype: list[int]
        """
        return self._requester_public_key

    @requester_public_key.setter
    def requester_public_key(self, requester_public_key):
        """Sets the requester_public_key of this JobRequester.

        The public key of the Requester node that created this job This can be used to encrypt messages back to the creator  # noqa: E501

        :param requester_public_key: The requester_public_key of this JobRequester.  # noqa: E501
        :type: list[int]
        """

        self._requester_public_key = requester_public_key

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
        if issubclass(JobRequester, dict):
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
        if not isinstance(other, JobRequester):
            return False

        return self.to_dict() == other.to_dict()

    def __ne__(self, other):
        """Returns true if both objects are not equal"""
        if not isinstance(other, JobRequester):
            return True

        return self.to_dict() != other.to_dict()
