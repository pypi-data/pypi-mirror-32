# -*- coding: utf-8 -*-

"""
    mundiapi.models.get_usages_details_response

    This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
"""
import mundiapi.models.get_period_response
import mundiapi.models.list_usages_details_response

class GetUsagesDetailsResponse(object):

    """Implementation of the 'GetUsagesDetailsResponse' model.

    TODO: type model description here.

    Attributes:
        subscription_id (string): Subscription Identifier
        total_amount (int): Current Invoice Amount
        period (GetPeriodResponse): Period Details
        usages (ListUsagesDetailsResponse): Usages Details
        total_discount (int): Total discounted value
        total_increment (int): Total inremented value

    """

    # Create a mapping from Model property names to API property names
    _names = {
        "subscription_id":'subscription_id',
        "total_amount":'total_amount',
        "period":'Period',
        "usages":'Usages',
        "total_discount":'total_discount',
        "total_increment":'total_increment'
    }

    def __init__(self,
                 subscription_id=None,
                 total_amount=None,
                 period=None,
                 usages=None,
                 total_discount=None,
                 total_increment=None):
        """Constructor for the GetUsagesDetailsResponse class"""

        # Initialize members of the class
        self.subscription_id = subscription_id
        self.total_amount = total_amount
        self.period = period
        self.usages = usages
        self.total_discount = total_discount
        self.total_increment = total_increment


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
        subscription_id = dictionary.get('subscription_id')
        total_amount = dictionary.get('total_amount')
        period = mundiapi.models.get_period_response.GetPeriodResponse.from_dictionary(dictionary.get('Period')) if dictionary.get('Period') else None
        usages = mundiapi.models.list_usages_details_response.ListUsagesDetailsResponse.from_dictionary(dictionary.get('Usages')) if dictionary.get('Usages') else None
        total_discount = dictionary.get('total_discount')
        total_increment = dictionary.get('total_increment')

        # Return an object of this model
        return cls(subscription_id,
                   total_amount,
                   period,
                   usages,
                   total_discount,
                   total_increment)


