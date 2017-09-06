// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/riskassessments"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/internal/apiresource"
)

// Client represents a riskassessments client. Thread-safe.
type Client struct {
	apiResource *apiresource.APIResource
}

// Bankaccounts represents the resource /{merchantId}/riskassessments/bankaccounts
// Risk-assess bankaccount
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/riskassessments/bankaccounts.html
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
func (c *Client) Bankaccounts(body riskassessments.RiskAssessmentBankAccount, context communication.CallContext) (riskassessments.RiskAssessmentResponse, error) {
	var resultObject riskassessments.RiskAssessmentResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/riskassessments/bankaccounts", nil)
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

// Cards represents the resource /{merchantId}/riskassessments/cards
// Risk-assess card
// Documentation can be found at https://epayments-api.developer-ingenico.com/s2sapi/v1/en_US/go/riskassessments/cards.html
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
func (c *Client) Cards(body riskassessments.RiskAssessmentCard, context communication.CallContext) (riskassessments.RiskAssessmentResponse, error) {
	var resultObject riskassessments.RiskAssessmentResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/{apiVersion}/{merchantId}/riskassessments/cards", nil)
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

// NewClient constructs a Riskassessments Client
//
// parent is the *apiresource.APIResource on top of which we want to build the new Riskassessments Client
func NewClient(parent *apiresource.APIResource, pathContext map[string]string) *Client {
	apiResource := apiresource.NewAPIResourceWithParent(parent, pathContext)

	return &Client{apiResource}
}
