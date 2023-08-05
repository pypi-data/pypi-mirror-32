# -*- coding: utf-8 -*-

"""
    mundiapi.models.create_checkout_card_installment_option_request

    This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
"""


class CreateCheckoutCardInstallmentOptionRequest(object):

    """Implementation of the 'CreateCheckoutCardInstallmentOptionRequest' model.

    Options for card installment

    Attributes:
        number (int): Installment quantity
        total (int): Total amount

    """

    # Create a mapping from Model property names to API property names
    _names = {
        "number":'number',
        "total":'total'
    }

    def __init__(self,
                 number=None,
                 total=None):
        """Constructor for the CreateCheckoutCardInstallmentOptionRequest class"""

        # Initialize members of the class
        self.number = number
        self.total = total


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
        number = dictionary.get('number')
        total = dictionary.get('total')

        # Return an object of this model
        return cls(number,
                   total)


