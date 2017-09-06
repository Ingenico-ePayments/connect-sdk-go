// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payouts

import (
	"net/http"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payout"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/internal/apiresource"
)

// Client represents a payouts client. Thread-safe.
type Client struct {
	apiResource *apiresource.APIResource
}

// Create represents the resource /{merchantId}/payouts
// Create payout
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/payouts/create.html
//
// Can return any of the following errors:
// * DeclinedPayoutError if the Ingenico ePayments platform declined / rejected the payout. The payout result will be available from the exception.
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the Ingenico ePayments platform,
// the Ingenico ePayments platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Ingenico ePayments platform returned any other error
func (c *Client) Create(body payout.CreateRequest, context communication.CallContext) (payout.Response, error) {
	var resultObject payout.Response

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payouts", nil)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
			case http.StatusBadRequest:
				{
					errorObject = &payout.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
			default:
				{
					errorObject = &errors.ErrorResponse{}
					err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
					if err != nil {
						return resultObject, err
					}

					break
				}
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

// Get represents the resource /{merchantId}/payouts/{payoutId}
// Get payout
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/payouts/get.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the Ingenico ePayments platform,
// the Ingenico ePayments platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Ingenico ePayments platform returned any other error
func (c *Client) Get(payoutID string, context communication.CallContext) (payout.Response, error) {
	var resultObject payout.Response

	pathContext := map[string]string{
		"payoutId": payoutID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payouts/{payoutId}", pathContext)
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

// Approve represents the resource /{merchantId}/payouts/{payoutId}/approve
// Approve payout
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/payouts/approve.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the Ingenico ePayments platform,
// the Ingenico ePayments platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Ingenico ePayments platform returned any other error
func (c *Client) Approve(payoutID string, body payout.ApproveRequest, context communication.CallContext) (payout.Response, error) {
	var resultObject payout.Response

	pathContext := map[string]string{
		"payoutId": payoutID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payouts/{payoutId}/approve", pathContext)
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

// Cancel represents the resource /{merchantId}/payouts/{payoutId}/cancel
// Cancel payout
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/payouts/cancel.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the Ingenico ePayments platform,
// the Ingenico ePayments platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Ingenico ePayments platform returned any other error
func (c *Client) Cancel(payoutID string, context communication.CallContext) error {
	pathContext := map[string]string{
		"payoutId": payoutID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payouts/{payoutId}/cancel", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, nil, context, &resultObject)
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

// Cancelapproval represents the resource /{merchantId}/payouts/{payoutId}/cancelapproval
// Undo approve payout
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/payouts/cancelapproval.html
//
// Can return any of the following errors:
// * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// * AuthorizationError if the request was not allowed (HTTP status code 403)
// * IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
// or there was a conflict (HTTP status code 404, 409 or 410)
// * GlobalCollectError if something went wrong at the Ingenico ePayments platform,
// the Ingenico ePayments platform was unable to process a message from a downstream partner/acquirer,
// or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// * APIError if the Ingenico ePayments platform returned any other error
func (c *Client) Cancelapproval(payoutID string, context communication.CallContext) error {
	pathContext := map[string]string{
		"payoutId": payoutID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/payouts/{payoutId}/cancelapproval", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	var resultObject map[string]interface{}
	postErr := c.apiResource.Communicator().Post(uri, clientHeaders, nil, nil, context, &resultObject)
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

// NewClient constructs a Payouts Client
//
// parent is the *apiresource.APIResource on top of which we want to build the new Payouts Client
func NewClient(parent *apiresource.APIResource, pathContext map[string]string) *Client {
	apiResource := apiresource.NewAPIResourceWithParent(parent, pathContext)

	return &Client{apiResource}
}
