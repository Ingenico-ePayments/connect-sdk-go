// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package files

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/internal/apiresource"
)

// Client represents a files client. Thread-safe.
type Client struct {
	apiResource *apiresource.APIResource
}

// GetFile represents the resource /{merchantId}/files/{fileId} - Retrieve File
// Documentation can be found at https://epayments-api.developer-ingenico.com/fileserviceapi/v1/en_US/go/files/getFile.html
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
func (c *Client) GetFile(fileID string, context communication.CallContext, bodyHandler communicator.BodyHandler) error {
	pathContext := map[string]string{
		"fileId": fileID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/files/v1/{merchantId}/files/{fileId}", pathContext)
	if err != nil {
		return err
	}

	clientHeaders := c.apiResource.ClientHeaders()

	getErr := c.apiResource.Communicator().GetWithHandler(uri, clientHeaders, nil, context, bodyHandler)
	if getErr != nil {
		responseError, isResponseError := getErr.(*sdkErrors.ResponseError)
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

		return getErr
	}

	return nil
}

// NewClient constructs a Files Client
//
// parent is the *apiresource.APIResource on top of which we want to build the new Files Client
func NewClient(parent *apiresource.APIResource, pathContext map[string]string) *Client {
	apiResource := apiresource.NewAPIResourceWithParent(parent, pathContext)

	return &Client{apiResource}
}
