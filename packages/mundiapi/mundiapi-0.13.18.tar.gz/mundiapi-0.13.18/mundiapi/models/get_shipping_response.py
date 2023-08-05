# -*- coding: utf-8 -*-

"""
    mundiapi.models.get_shipping_response

    This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
"""
import mundiapi.models.get_address_response

class GetShippingResponse(object):

    """Implementation of the 'GetShippingResponse' model.

    Response object for getting the shipping data

    Attributes:
        amount (int): TODO: type description here.
        description (string): TODO: type description here.
        recipient_name (string): TODO: type description here.
        recipient_phone (string): TODO: type description here.
        address (GetAddressResponse): TODO: type description here.

    """

    # Create a mapping from Model property names to API property names
    _names = {
        "amount":'amount',
        "description":'description',
        "recipient_name":'recipient_name',
        "recipient_phone":'recipient_phone',
        "address":'address'
    }

    def __init__(self,
                 amount=None,
                 description=None,
                 recipient_name=None,
                 recipient_phone=None,
                 address=None):
        """Constructor for the GetShippingResponse class"""

        # Initialize members of the class
        self.amount = amount
        self.description = description
        self.recipient_name = recipient_name
        self.recipient_phone = recipient_phone
        self.address = address


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
        amount = dictionary.get('amount')
        description = dictionary.get('description')
        recipient_name = dictionary.get('recipient_name')
        recipient_phone = dictionary.get('recipient_phone')
        address = mundiapi.models.get_address_response.GetAddressResponse.from_dictionary(dictionary.get('address')) if dictionary.get('address') else None

        # Return an object of this model
        return cls(amount,
                   description,
                   recipient_name,
                   recipient_phone,
                   address)


