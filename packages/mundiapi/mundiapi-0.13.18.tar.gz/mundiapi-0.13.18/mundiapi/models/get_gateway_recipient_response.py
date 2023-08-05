# -*- coding: utf-8 -*-

"""
    mundiapi.models.get_gateway_recipient_response

    This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
"""


class GetGatewayRecipientResponse(object):

    """Implementation of the 'GetGatewayRecipientResponse' model.

    Information about the recipient on the gateway

    Attributes:
        gateway (string): Gateway name
        status (string): Status of the recipient on the gateway
        pgid (string): Recipient id on the gateway
        created_at (string): Creation date
        updated_at (string): Last update date

    """

    # Create a mapping from Model property names to API property names
    _names = {
        "gateway":'gateway',
        "status":'status',
        "pgid":'pgid',
        "created_at":'created_at',
        "updated_at":'updated_at'
    }

    def __init__(self,
                 gateway=None,
                 status=None,
                 pgid=None,
                 created_at=None,
                 updated_at=None):
        """Constructor for the GetGatewayRecipientResponse class"""

        # Initialize members of the class
        self.gateway = gateway
        self.status = status
        self.pgid = pgid
        self.created_at = created_at
        self.updated_at = updated_at


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
        gateway = dictionary.get('gateway')
        status = dictionary.get('status')
        pgid = dictionary.get('pgid')
        created_at = dictionary.get('created_at')
        updated_at = dictionary.get('updated_at')

        # Return an object of this model
        return cls(gateway,
                   status,
                   pgid,
                   created_at,
                   updated_at)


