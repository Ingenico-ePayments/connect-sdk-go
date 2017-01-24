// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package products

import (
	"strconv"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/product"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/publickey"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/internal/apiresource"
)

// Client represents a products client. Thread-safe.
type Client struct {
	apiResource *apiresource.APIResource
}

// Find represents the resource /{merchantId}/products
// Get payment products
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Find(query FindParams, context communication.CallContext) (product.PaymentProducts, error) {
	var resultObject product.PaymentProducts

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/products", nil)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().Get(uri, clientHeaders, &query, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
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

		return resultObject, getErr
	}

	return resultObject, nil
}

// Get represents the resource /{merchantId}/products/{paymentProductId}
// Get payment product
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Get(paymentProductID int32, query GetParams, context communication.CallContext) (product.PaymentProductResponse, error) {
	var resultObject product.PaymentProductResponse

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/products/{paymentProductId}", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().Get(uri, clientHeaders, &query, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
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

		return resultObject, getErr
	}

	return resultObject, nil
}

// Directory represents the resource /{merchantId}/products/{paymentProductId}/directory
// Get payment product directory
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Directory(paymentProductID int32, query DirectoryParams, context communication.CallContext) (product.Directory, error) {
	var resultObject product.Directory

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/products/{paymentProductId}/directory", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().Get(uri, clientHeaders, &query, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
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

		return resultObject, getErr
	}

	return resultObject, nil
}

// Networks represents the resource /{merchantId}/products/{paymentProductId}/networks
// Get payment product networks
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) Networks(paymentProductID int32, query NetworksParams, context communication.CallContext) (product.PaymentProductNetworksResponse, error) {
	var resultObject product.PaymentProductNetworksResponse

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/products/{paymentProductId}/networks", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().Get(uri, clientHeaders, &query, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
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

		return resultObject, getErr
	}

	return resultObject, nil
}

// PublicKey represents the resource /{merchantId}/products/{paymentProductId}/publicKey
// Get payment product specific public key
// Documentation can be found at $devportal_call_uri($overload)
//
// Can return any of the following errors:
// ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
// AuthorizationError if the request was not allowed (HTTP status code 403)
// IdempotenceError if an idempotent request caused a conflict (HTTP status code 409)
// ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//            or there was a conflict (HTTP status code 404, 409 or 410)
// GlobalCollectError if something went wrong at the GlobalCollect platform,
//            the GlobalCollect platform was unable to process a message from a downstream partner/acquirer,
//            or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
// APIError if the GlobalCollect platform returned any other error
func (c *Client) PublicKey(paymentProductID int32, context communication.CallContext) (publickey.PublicKey, error) {
	var resultObject publickey.PublicKey

	pathContext := map[string]string{
		"paymentProductId": strconv.FormatInt(int64(paymentProductID), 10),
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/products/{paymentProductId}/publicKey", pathContext)
	if err != nil {
		return resultObject, err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().Get(uri, clientHeaders, nil, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*sdkErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			switch responseError.StatusCode() {
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

		return resultObject, getErr
	}

	return resultObject, nil
}

// NewClient constructs a Products Client
//
// parent is the *apiresource.APIResource on top of which we want to build the new Products Client
func NewClient(parent *apiresource.APIResource, pathContext map[string]string) *Client {
	apiResource := apiresource.NewAPIResourceWithParent(parent, pathContext)

	return &Client{apiResource}
}
