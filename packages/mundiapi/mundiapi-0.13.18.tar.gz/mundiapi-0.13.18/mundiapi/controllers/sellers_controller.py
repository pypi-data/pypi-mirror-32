# -*- coding: utf-8 -*-

"""
    mundiapi.controllers.sellers_controller

    This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io ).
"""

from .base_controller import BaseController
from ..api_helper import APIHelper
from ..configuration import Configuration
from ..http.auth.basic_auth import BasicAuth
from ..models.get_seller_response import GetSellerResponse
from ..models.list_seller_response import ListSellerResponse

class SellersController(BaseController):

    """A Controller to access Endpoints in the mundiapi API."""


    def create_seller(self,
                      request):
        """Does a POST request to /sellers/.

        TODO: type endpoint description here.

        Args:
            request (CreateSellerRequest): Seller Model

        Returns:
            GetSellerResponse: Response from the API. 

        Raises:
            APIException: When an error occurs while fetching the data from
                the remote API. This exception includes the HTTP Response
                code, an error message, and the HTTP body that was received in
                the request.

        """

        # Prepare query URL
        _query_builder = Configuration.base_uri
        _query_builder += '/sellers/'
        _query_url = APIHelper.clean_url(_query_builder)

        # Prepare headers
        _headers = {
            'accept': 'application/json',
            'content-type': 'application/json; charset=utf-8'
        }

        # Prepare and execute request
        _request = self.http_client.post(_query_url, headers=_headers, parameters=APIHelper.json_serialize(request))
        BasicAuth.apply(_request)
        _context = self.execute_request(_request)
        self.validate_response(_context)

        # Return appropriate type
        return APIHelper.json_deserialize(_context.response.raw_body, GetSellerResponse.from_dictionary)

    def delete_seller(self,
                      seller_id):
        """Does a DELETE request to /sellers/{sellerId}.

        TODO: type endpoint description here.

        Args:
            seller_id (string): Seller Id

        Returns:
            GetSellerResponse: Response from the API. 

        Raises:
            APIException: When an error occurs while fetching the data from
                the remote API. This exception includes the HTTP Response
                code, an error message, and the HTTP body that was received in
                the request.

        """

        # Prepare query URL
        _query_builder = Configuration.base_uri
        _query_builder += '/sellers/{sellerId}'
        _query_builder = APIHelper.append_url_with_template_parameters(_query_builder, { 
            'sellerId': seller_id
        })
        _query_url = APIHelper.clean_url(_query_builder)

        # Prepare headers
        _headers = {
            'accept': 'application/json'
        }

        # Prepare and execute request
        _request = self.http_client.delete(_query_url, headers=_headers)
        BasicAuth.apply(_request)
        _context = self.execute_request(_request)
        self.validate_response(_context)

        # Return appropriate type
        return APIHelper.json_deserialize(_context.response.raw_body, GetSellerResponse.from_dictionary)

    def get_seller_by_id(self,
                         id):
        """Does a GET request to /sellers/{id}.

        TODO: type endpoint description here.

        Args:
            id (string): Seller Id

        Returns:
            GetSellerResponse: Response from the API. 

        Raises:
            APIException: When an error occurs while fetching the data from
                the remote API. This exception includes the HTTP Response
                code, an error message, and the HTTP body that was received in
                the request.

        """

        # Prepare query URL
        _query_builder = Configuration.base_uri
        _query_builder += '/sellers/{id}'
        _query_builder = APIHelper.append_url_with_template_parameters(_query_builder, { 
            'id': id
        })
        _query_url = APIHelper.clean_url(_query_builder)

        # Prepare headers
        _headers = {
            'accept': 'application/json'
        }

        # Prepare and execute request
        _request = self.http_client.get(_query_url, headers=_headers)
        BasicAuth.apply(_request)
        _context = self.execute_request(_request)
        self.validate_response(_context)

        # Return appropriate type
        return APIHelper.json_deserialize(_context.response.raw_body, GetSellerResponse.from_dictionary)

    def get_sellers(self,
                    page=None,
                    size=None,
                    name=None,
                    document=None,
                    code=None,
                    status=None,
                    mtype=None,
                    created_since=None,
                    created_until=None):
        """Does a GET request to /sellers.

        TODO: type endpoint description here.

        Args:
            page (int, optional): Page number
            size (int, optional): Page size
            name (string, optional): TODO: type description here. Example: 
            document (string, optional): TODO: type description here. Example:
                            code (string, optional): TODO: type description here. Example: 
            status (string, optional): TODO: type description here. Example: 
            mtype (string, optional): TODO: type description here. Example: 
            created_since (datetime, optional): TODO: type description here.
                Example: 
            created_until (datetime, optional): TODO: type description here.
                Example: 

        Returns:
            ListSellerResponse: Response from the API. 

        Raises:
            APIException: When an error occurs while fetching the data from
                the remote API. This exception includes the HTTP Response
                code, an error message, and the HTTP body that was received in
                the request.

        """

        # Prepare query URL
        _query_builder = Configuration.base_uri
        _query_builder += '/sellers'
        _query_parameters = {
            'page': page,
            'size': size,
            'name': name,
            'document': document,
            'code': code,
            'status': status,
            'type': mtype,
            'created_Since': APIHelper.when_defined(APIHelper.RFC3339DateTime, created_since),
            'created_Until': APIHelper.when_defined(APIHelper.RFC3339DateTime, created_until)
        }
        _query_builder = APIHelper.append_url_with_query_parameters(_query_builder,
            _query_parameters, Configuration.array_serialization)
        _query_url = APIHelper.clean_url(_query_builder)

        # Prepare headers
        _headers = {
            'accept': 'application/json'
        }

        # Prepare and execute request
        _request = self.http_client.get(_query_url, headers=_headers)
        BasicAuth.apply(_request)
        _context = self.execute_request(_request)
        self.validate_response(_context)

        # Return appropriate type
        return APIHelper.json_deserialize(_context.response.raw_body, ListSellerResponse.from_dictionary)

    def update_seller(self,
                      id,
                      request):
        """Does a PUT request to /sellers/{id}.

        TODO: type endpoint description here.

        Args:
            id (string): TODO: type description here. Example: 
            request (UpdateSellerRequest): Update Seller model

        Returns:
            GetSellerResponse: Response from the API. 

        Raises:
            APIException: When an error occurs while fetching the data from
                the remote API. This exception includes the HTTP Response
                code, an error message, and the HTTP body that was received in
                the request.

        """

        # Prepare query URL
        _query_builder = Configuration.base_uri
        _query_builder += '/sellers/{id}'
        _query_builder = APIHelper.append_url_with_template_parameters(_query_builder, { 
            'id': id
        })
        _query_url = APIHelper.clean_url(_query_builder)

        # Prepare headers
        _headers = {
            'accept': 'application/json',
            'content-type': 'application/json; charset=utf-8'
        }

        # Prepare and execute request
        _request = self.http_client.put(_query_url, headers=_headers, parameters=APIHelper.json_serialize(request))
        BasicAuth.apply(_request)
        _context = self.execute_request(_request)
        self.validate_response(_context)

        # Return appropriate type
        return APIHelper.json_deserialize(_context.response.raw_body, GetSellerResponse.from_dictionary)

    def update_seller_metadata(self,
                               seller_id,
                               request):
        """Does a PATCH request to /sellers/{seller_id}/metadata.

        TODO: type endpoint description here.

        Args:
            seller_id (string): Seller Id
            request (UpdateMetadataRequest): Request for updating the charge
                metadata

        Returns:
            GetSellerResponse: Response from the API. 

        Raises:
            APIException: When an error occurs while fetching the data from
                the remote API. This exception includes the HTTP Response
                code, an error message, and the HTTP body that was received in
                the request.

        """

        # Prepare query URL
        _query_builder = Configuration.base_uri
        _query_builder += '/sellers/{seller_id}/metadata'
        _query_builder = APIHelper.append_url_with_template_parameters(_query_builder, { 
            'seller_id': seller_id
        })
        _query_url = APIHelper.clean_url(_query_builder)

        # Prepare headers
        _headers = {
            'accept': 'application/json',
            'content-type': 'application/json; charset=utf-8'
        }

        # Prepare and execute request
        _request = self.http_client.patch(_query_url, headers=_headers, parameters=APIHelper.json_serialize(request))
        BasicAuth.apply(_request)
        _context = self.execute_request(_request)
        self.validate_response(_context)

        # Return appropriate type
        return APIHelper.json_deserialize(_context.response.raw_body, GetSellerResponse.from_dictionary)
