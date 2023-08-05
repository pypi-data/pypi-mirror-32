# -*- coding: utf-8 -*-

"""
    mundiapi.models.get_discount_response

    This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
"""
from mundiapi.api_helper import APIHelper
import mundiapi.models.get_subscription_response

class GetDiscountResponse(object):

    """Implementation of the 'GetDiscountResponse' model.

    Response object for getting a discount

    Attributes:
        id (string): TODO: type description here.
        value (float): TODO: type description here.
        discount_type (string): TODO: type description here.
        status (string): TODO: type description here.
        created_at (datetime): TODO: type description here.
        subscription (GetSubscriptionResponse): TODO: type description here.
        cycles (int): TODO: type description here.
        deleted_at (datetime): TODO: type description here.
        description (string): TODO: type description here.

    """

    # Create a mapping from Model property names to API property names
    _names = {
        "id":'id',
        "value":'value',
        "discount_type":'discount_type',
        "status":'status',
        "created_at":'created_at',
        "subscription":'subscription',
        "cycles":'cycles',
        "deleted_at":'deleted_at',
        "description":'description'
    }

    def __init__(self,
                 id=None,
                 value=None,
                 discount_type=None,
                 status=None,
                 created_at=None,
                 subscription=None,
                 cycles=None,
                 deleted_at=None,
                 description=None):
        """Constructor for the GetDiscountResponse class"""

        # Initialize members of the class
        self.id = id
        self.value = value
        self.discount_type = discount_type
        self.status = status
        self.created_at = APIHelper.RFC3339DateTime(created_at) if created_at else None
        self.subscription = subscription
        self.cycles = cycles
        self.deleted_at = APIHelper.RFC3339DateTime(deleted_at) if deleted_at else None
        self.description = description


    @classmethod
    def from_dictionary(cls,
                        dictionary):
        """Creates an instance of this model from a dictionary

        Args:
            dictionary (dictionary): A dictionary representation of the object as
            obtained from the deserialization of the server's response. The keys
            MUST match property names in the API description.

        Returns:
            object: An instance of this structure class.

        """
        if dictionary is None:
            return None

        # Extract variables from the dictionary
        id = dictionary.get('id')
        value = dictionary.get('value')
        discount_type = dictionary.get('discount_type')
        status = dictionary.get('status')
        created_at = APIHelper.RFC3339DateTime.from_value(dictionary.get("created_at")).datetime if dictionary.get("created_at") else None
        subscription = mundiapi.models.get_subscription_response.GetSubscriptionResponse.from_dictionary(dictionary.get('subscription')) if dictionary.get('subscription') else None
        cycles = dictionary.get('cycles')
        deleted_at = APIHelper.RFC3339DateTime.from_value(dictionary.get("deleted_at")).datetime if dictionary.get("deleted_at") else None
        description = dictionary.get('description')

        # Return an object of this model
        return cls(id,
                   value,
                   discount_type,
                   status,
                   created_at,
                   subscription,
                   cycles,
                   deleted_at,
                   description)


