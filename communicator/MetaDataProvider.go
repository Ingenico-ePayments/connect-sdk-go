package communicator

import (
	"encoding/base64"
	"errors"
	"runtime"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/metadata"
)

var (
	// ErrProhibitedHeader occurs when the header name is prohibited
	ErrProhibitedHeader = errors.New("header is prohibited")
)

type serverMetaInfo struct {
	PlatformIdentifier    string                          `json:"platformIdentifier,omitempty"`
	SDKIdentifier         string                          `json:"sdkIdentifier,omitempty"`
	SDKCreator            string                          `json:"sdkCreator,omitempty"`
	Integrator            string                          `json:"integrator,omitempty"`
	ShoppingCartExtension *metadata.ShoppingCartExtension `json:"shoppingCartExtension,omitempty"`
}

// MetaDataProvider provides meta info about the server. Thread-safe.
type MetaDataProvider struct {
	serverMetaInfo  serverMetaInfo
	metaDataHeaders []communication.Header
}

// ProhibitedHeaders are the headers that can't be included in the server requests
var ProhibitedHeaders = []string{serverMetaInfoHeader, "X-GCS-Idempotence-Key",
	"Date", "Content-Type", "Authorization"}

func getPlatformIdentifier() string {
	//I believe there is no standard function thus we need to do something like
	//https://github.com/matishsiao/goInfo
	return runtime.GOOS + " " + runtime.Version() + "(" + runtime.GOARCH + ")"
}

const sdkIdentifier = "GoServerSDK/v" + sdkVersion
const sdkVersion = "2.9.0"
const serverMetaInfoHeader = "X-GCS-ServerMetaInfo"

// NewMetaDataProviderWithBuilder creates a MetaDataProvider with the given MetaDataProviderBuilder
func NewMetaDataProviderWithBuilder(builder MetaDataProviderBuilder) (*MetaDataProvider, error) {
	return newMetaDataProvider(builder.integrator, builder.ShoppingCartExtension, builder.AdditionalRequestHeaders)
}

// NewMetaDataProviderWithIntegrator creates a MetaDataProvider with the given integrator
func NewMetaDataProviderWithIntegrator(integrator string) (*MetaDataProvider, error) {
	return newMetaDataProvider(integrator, nil, nil)
}

func newMetaDataProvider(integrator string, shoppingCartExtension *metadata.ShoppingCartExtension, additionalRequestHeaders []communication.Header) (*MetaDataProvider, error) {
	err := validateAdditionalRequestHeaders(additionalRequestHeaders)
	if err != nil {
		return nil, err
	}
	sMetaInfo := serverMetaInfo{
		getPlatformIdentifier(),
		sdkIdentifier,
		"Ingenico",
		integrator,
		shoppingCartExtension,
	}
	marshaller, err := defaultimpl.NewDefaultMarshaller()
	if err != nil {
		return nil, err
	}

	serverMetaInfoString, err := marshaller.Marshal(&sMetaInfo)
	if err != nil {
		return nil, err
	}

	serverMetaInfoBase64 := base64.StdEncoding.EncodeToString([]byte(serverMetaInfoString))
	header, err := communication.NewHeader(serverMetaInfoHeader, serverMetaInfoBase64)
	head := []communication.Header{*header}
	head = append(head, additionalRequestHeaders...)
	if err != nil {
		return nil, err
	}
	return &MetaDataProvider{sMetaInfo, head}, nil
}

func validateAdditionalRequestHeader(additionalRequestHeader communication.Header) error {
	for _, a := range ProhibitedHeaders {
		if a == additionalRequestHeader.Name() {
			return ErrProhibitedHeader
		}
	}
	return nil
}

func validateAdditionalRequestHeaders(additionalRequestHeaders []communication.Header) error {
	for _, additionalRequestHeader := range additionalRequestHeaders {
		err := validateAdditionalRequestHeader(additionalRequestHeader)
		if err != nil {
			return err
		}
	}
	return nil
}

// MetaDataHeaders returns the server related headers containing the metadata to be associated with the request (if any).
// This will always contain at least an automatically generated header X-GCS-ServerMetaInfo.
func (m *MetaDataProvider) MetaDataHeaders() []communication.Header {
	// Return a clone instead of the original slice - immutability insurance
	return append([]communication.Header{}, m.metaDataHeaders...)
}
