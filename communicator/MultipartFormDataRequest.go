package communicator

import "github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"

// MultipartFormDataRequest represents a multipart/form-data request.
type MultipartFormDataRequest interface {

	// ToMultipartFormDataObject converts this multipart/form-data request into a MultipartFormDataObject.
	ToMultipartFormDataObject() *communication.MultipartFormDataObject
}
