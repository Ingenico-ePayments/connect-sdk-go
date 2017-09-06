package apiresource

import (
	"errors"
	"strings"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
)

// ErrNilParent occurs when the parent is nil
var ErrNilParent = errors.New("parent is nil")

// ErrNilCommunicator occurs when the communicator is nil
var ErrNilCommunicator = errors.New("communicator is nil")

// APIResource represents the base type of all Ingenico ePayments platform API resources.
type APIResource struct {
	parentResource *APIResource
	communicator   *communicator.Communicator
	clientMetaInfo string
	pathContext    map[string]string
	errorObject    error
}

// Communicator returns the Communicator used by the resource
func (ar *APIResource) Communicator() *communicator.Communicator {
	return ar.communicator
}

// ClientMetaInfo returns the ClientMetaInfo used by the resource
func (ar *APIResource) ClientMetaInfo() string {
	return ar.clientMetaInfo
}

// ClientHeaders returns the headers used by the resource
func (ar *APIResource) ClientHeaders() []communication.Header {
	if len(ar.clientMetaInfo) != 0 {
		header, _ := communication.NewHeader("X-GCS-ClientMetaInfo", ar.clientMetaInfo)

		return []communication.Header{*header}
	}

	return nil
}

// InstantiateURIWithContext instantiates the given URI with the path context
func (ar *APIResource) InstantiateURIWithContext(uri string, pathContext map[string]string) (string, error) {
	return ar.InstantiateURI(replaceAll(uri, pathContext))
}

func (ar *APIResource) getError() error {
	if ar.parentResource != nil {
		if parentError := ar.parentResource.getError(); parentError != nil {
			return parentError
		}
		if ar.errorObject != nil {
			err := ar.errorObject
			ar.errorObject = nil

			return err
		}
	}
	return nil
}

// InstantiateURI instantiates the given uri with the path context of the resource
func (ar *APIResource) InstantiateURI(uri string) (string, error) {
	if err := ar.getError(); err != nil {
		return "", err
	}
	uri = replaceAll(uri, ar.pathContext)

	if ar.parentResource != nil {
		uriParent, err := ar.parentResource.InstantiateURI(uri)

		if err != nil {
			return uriParent, err
		}

		uri = uriParent
	}

	return uri, nil
}

// NewAPIResourceWithParent creates an APIResource with the given parent and pathContext
func NewAPIResourceWithParent(parent *APIResource, pathContext map[string]string) *APIResource {
	ar := &APIResource{parent, parent.communicator, parent.clientMetaInfo, pathContext, nil}
	if parent == nil {
		ar.errorObject = ErrNilParent
	}
	return ar
}

// NewAPIResource creates an APIResource with the given communicator, clientMetaInfo and pathContext
func NewAPIResource(communicator *communicator.Communicator, clientMetaInfo string, pathContext map[string]string) (*APIResource, error) {
	ar := &APIResource{nil, communicator, clientMetaInfo, pathContext, nil}
	if communicator == nil {
		ar.errorObject = ErrNilCommunicator
		return ar, ar.errorObject
	}
	return ar, nil
}

func replaceAll(uri string, pathContext map[string]string) string {
	if pathContext != nil {
		for key, value := range pathContext {
			uri = strings.Replace(uri, "{" + key + "}", value, -1)
		}
	}

	return uri
}
