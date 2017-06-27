// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package tokens

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/token"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/internal/apiresource"
)

// Client represents a tokens client. Thread-safe.
type Client struct {
	apiResource *apiresource.APIResource
}

// Create represents the resource /{merchantId}/tokens
// Create token
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/tokens/create.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the GlobalCollect platform,
// the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the GlobalCollect platform returned any other error
func (c *Client) Create(body token.CreateRequest, context communication.CallContext) (token.CreateResponse, error) {
	var resultObject token.CreateResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/tokens", nil)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &errors.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return resultObject, err
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// Get represents the resource /{merchantId}/tokens/{tokenId}
// Get token
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/tokens/get.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the GlobalCollect platform,
// the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the GlobalCollect platform returned any other error
func (c *Client) Get(tokenID string, context communication.CallContext) (token.Response, error) {
	var resultObject token.Response

	pathContext := map[string]string{
		"tokenId": tokenID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/tokens/{tokenId}", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().Get(uri, clientHeaders, nil, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &errors.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return resultObject, err
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, getErr
	}

	return resultObject, nil
}

// Update represents the resource /{merchantId}/tokens/{tokenId}
// Update token
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/tokens/update.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the GlobalCollect platform,
// the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the GlobalCollect platform returned any other error
func (c *Client) Update(tokenID string, body token.UpdateRequest, context communication.CallContext) error {
	pathContext := map[string]string{
		"tokenId": tokenID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/tokens/{tokenId}", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	putErr := c.apiResource.Communicator().Put(uri, clientHeaders, nil, body, context, &resultObject)
	if putErr != nil {
		responseError, isResponseError := putErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &errors.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return err
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return createErr
			}

			return err
		}

		return putErr
	}

	return nil
}

// Delete represents the resource /{merchantId}/tokens/{tokenId}
// Delete token
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/tokens/delete.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the GlobalCollect platform,
// the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the GlobalCollect platform returned any other error
func (c *Client) Delete(tokenID string, query DeleteParams, context communication.CallContext) error {
	pathContext := map[string]string{
		"tokenId": tokenID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/tokens/{tokenId}", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	deleteErr := c.apiResource.Communicator().Delete(uri, clientHeaders, &query, context, &resultObject)
	if deleteErr != nil {
		responseError, isResponseError := deleteErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &errors.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return err
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return createErr
			}

			return err
		}

		return deleteErr
	}

	return nil
}

// Approvesepadirectdebit represents the resource /{merchantId}/tokens/{tokenId}/approvesepadirectdebit
// Approve SEPA DD mandate
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/tokens/approvesepadirectdebit.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the GlobalCollect platform,
// the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the GlobalCollect platform returned any other error
func (c *Client) Approvesepadirectdebit(tokenID string, body token.ApproveRequest, context communication.CallContext) error {
	pathContext := map[string]string{
		"tokenId": tokenID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/tokens/{tokenId}/approvesepadirectdebit", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &errors.ErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return err
			}

			err, createErr := sdkErrors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return createErr
			}

			return err
		}

		return postErr
	}

	return nil
}

// NewClient constructs a Tokens Client
//
// parent is the *apiresource.APIResource on top of which we want to build the new Tokens Client
func NewClient(parent *apiresource.APIResource, pathContext map[string]string) *Client {
	apiResource := apiresource.NewAPIResourceWithParent(parent, pathContext)

	return &Client{apiResource}
}
