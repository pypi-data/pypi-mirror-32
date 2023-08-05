# -*- coding: utf-8 -*-

"""
    mundiapi.models.create_recipient_request

    This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
"""
import mundiapi.models.create_bank_account_request

class CreateRecipientRequest(object):

    """Implementation of the 'CreateRecipientRequest' model.

    Request for creating a recipient

    Attributes:
        name (string): Recipient name
        email (string): Recipient email
        description (string): Recipient description
        document (string): Recipient document number
        mtype (string): Recipient type
        default_bank_account (CreateBankAccountRequest): Bank account
        metadata (dict<object, string>): Metadata

    """

    # Create a mapping from Model property names to API property names
    _names = {
        "name":'name',
        "email":'email',
        "description":'description',
        "document":'document',
        "mtype":'type',
        "default_bank_account":'default_bank_account',
        "metadata":'metadata'
    }

    def __init__(self,
                 name=None,
                 email=None,
                 description=None,
                 document=None,
                 mtype=None,
                 default_bank_account=None,
                 metadata=None):
        """Constructor for the CreateRecipientRequest class"""

        # Initialize members of the class
        self.name = name
        self.email = email
        self.description = description
        self.document = document
        self.mtype = mtype
        self.default_bank_account = default_bank_account
        self.metadata = metadata


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
        name = dictionary.get('name')
        email = dictionary.get('email')
        description = dictionary.get('description')
        document = dictionary.get('document')
        mtype = dictionary.get('type')
        default_bank_account = mundiapi.models.create_bank_account_request.CreateBankAccountRequest.from_dictionary(dictionary.get('default_bank_account')) if dictionary.get('default_bank_account') else None
        metadata = dictionary.get('metadata')

        # Return an object of this model
        return cls(name,
                   email,
                   description,
                   document,
                   mtype,
                   default_bank_account,
                   metadata)


