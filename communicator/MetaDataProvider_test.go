package communicator

import (
	"encoding/base64"
	"encoding/json"
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/metadata"
)

func TestGetServerMetaDataHeadersNoAdditionalHeaders(t *testing.T) {

	metaDataProvider, err := NewMetaDataProviderWithIntegrator("Ingenico")
	if err != nil {
		t.Fatalf("TestGetServerMetaDataHeadersNoAdditionalHeaders : %v", err)
	}

	requestHeaders := metaDataProvider.MetaDataHeaders()
	if requestHeaders == nil {
		t.Fatal("TestGetServerMetaDataHeadersNoAdditionalHeaders : nil headers")
	}
	if len(requestHeaders) != 1 {
		t.Fatal("TestGetServerMetaDataHeadersNoAdditionalHeaders : len != 1")
	}

	AssertServerMetaInfo("TestGetServerMetaDataHeadersNoAdditionalHeaders",
		t, metaDataProvider, "Ingenico", requestHeaders[0])
}

func TestServerMetaDataHeadersFull(t *testing.T) {
	shoppingCartExtension, err := metadata.NewShoppingCartExtension("Ingenico.creator", "Extension", "1.0")
	if err != nil {
		t.Fatalf("TestServerMetaDataHeadersFull : %v", err)
	}

	integrator := "Ingenico"
	metaDataProviderBuilder := NewMetaDataProviderBuilder(integrator)
	metaDataProviderBuilder.ShoppingCartExtension = shoppingCartExtension
	metaDataProvider, err := metaDataProviderBuilder.Build()
	if err != nil {
		t.Fatalf("TestServerMetaDataHeadersFull : %v", err)
	}

	requestHeaders := metaDataProvider.MetaDataHeaders()
	if len(requestHeaders) != 1 {
		t.Fatal("TestServerMetaDataHeadersFull : len != 1")
	}

	AssertServerMetaInfo("TestServerMetaDataHeadersFull",
		t, metaDataProvider, integrator, requestHeaders[0])
	AssertShoppingCard("TestServerMetaDataHeadersFull",
		t, metaDataProvider, shoppingCartExtension)
}

func TestServerMetaDataHeadersAdditionalHeaders(t *testing.T) {
	header1, err1 := communication.NewHeader("Header1", "&=$%")
	header2, err2 := communication.NewHeader("Header2", "blah blah")
	header3, err3 := communication.NewHeader("Header3", "foo")
	if err1 != nil {
		t.Fatalf("TestServerMetaDataHeadersAdditionalHeaders : %v", err1)
	}
	if err2 != nil {
		t.Fatalf("TestServerMetaDataHeadersAdditionalHeaders : %v", err2)
	}
	if err3 != nil {
		t.Fatalf("TestServerMetaDataHeadersAdditionalHeaders : %v", err3)
	}

	additionalHeaders := []communication.Header{*header1, *header2, *header3}
	metaDataProviderBuilder := NewMetaDataProviderBuilder("Ingenico")
	metaDataProviderBuilder.AdditionalRequestHeaders = additionalHeaders
	metaDataProvider, err := metaDataProviderBuilder.Build()
	if err != nil {
		t.Fatalf("TestServerMetaDataHeadersAdditionalHeaders : %v", err)
	}

	requestHeaders := metaDataProvider.MetaDataHeaders()
	if len(requestHeaders) != 4 {
		t.Fatal("TestServerMetaDataHeadersAdditionalHeaders : len != 4")
	}

	for i := 1; i < len(requestHeaders); i++ {
		if additionalHeaders[i-1] != requestHeaders[i] {
			t.Fatalf("TestServerMetaDataHeadersAdditionalHeaders : [%d]%v != [%d]%v",
				i-1, additionalHeaders[i-1],
				i, requestHeaders[i])
		}
	}
}

func TestConstructorWithProhibitedHeaders(t *testing.T) {
	header1, err1 := communication.NewHeader("Header1", "&=$%")
	header2, err2 := communication.NewHeader("Header2", "blah blah")
	if err1 != nil {
		t.Fatalf("TestConstructorWithProhibitedHeaders : %v", err1)
	}
	if err2 != nil {
		t.Fatalf("TestConstructorWithProhibitedHeaders : %v", err2)
	}

	for _, v := range ProhibitedHeaders {
		header3, err3 := communication.NewHeader(v, "prohibited header")
		if err3 != nil {
			t.Fatalf("TestConstructorWithProhibitedHeaders : %v", err3)
		}

		additionalHeaders := []communication.Header{*header1, *header2, *header3}
		builder := NewMetaDataProviderBuilder("builder")
		builder.AdditionalRequestHeaders = additionalHeaders
		metadataProvider, err := builder.Build()
		if err == nil {
			t.Fatal("TestConstructorWithProhibitedHeaders : err == nil")
		}

		if metadataProvider != nil {
			t.Fatal("TestConstructorWithProhibitedHeaders : metadataProvider != nil")
		}
	}
}

func AssertServerMetaInfo(prefix string, t *testing.T, metaDataProvider *MetaDataProvider,
	integrator string, requestHeader communication.Header) {

	if requestHeader.Name() != "X-GCS-ServerMetaInfo" {
		t.Fatalf("[%s]AssertServerMetaInfo : %s != %s", prefix, requestHeader.Name(), "X-GCS-ServerMetaInfo")
	}
	if len(requestHeader.Value()) < 1 {
		t.Fatalf("[%s]AssertServerMetaInfo : requestHeader value is empty", prefix)
	}

	serverMetaInfo := metaDataProvider.serverMetaInfo
	reqHeaderBytes, err := base64.StdEncoding.DecodeString(requestHeader.Value())
	if err != nil {
		t.Fatalf("[%s]AssertServerMetaInfo : %v", prefix, err)
	}

	err = json.Unmarshal(reqHeaderBytes, &serverMetaInfo)
	if err != nil {
		t.Fatalf("[%s]AssertServerMetaInfo : %v", prefix, err)
	}

	if serverMetaInfo.PlatformIdentifier != metaDataProvider.serverMetaInfo.PlatformIdentifier {
		t.Fatalf("[%s]AssertServerMetaInfo : platformIdentifier '%s' != '%s'", prefix,
			serverMetaInfo.PlatformIdentifier, metaDataProvider.serverMetaInfo.PlatformIdentifier)
	}
	if serverMetaInfo.SDKIdentifier != metaDataProvider.serverMetaInfo.SDKIdentifier {
		t.Fatalf("[%s]AssertServerMetaInfo : SDKIdentifier '%s' != '%s'", prefix,
			serverMetaInfo.SDKIdentifier, metaDataProvider.serverMetaInfo.SDKIdentifier)
	}
	if serverMetaInfo.Integrator != integrator {
		t.Fatalf("[%s]AssertServerMetaInfo : Integrator '%s' != '%s'", prefix,
			serverMetaInfo.Integrator, integrator)
	}
	if serverMetaInfo.SDKCreator != "Ingenico" {
		t.Fatalf("[%s]AssertServerMetaInfo : SDKCreator '%s' != '%s'", prefix,
			serverMetaInfo.SDKCreator, "Ingenico")
	}
}

func AssertShoppingCard(prefix string, t *testing.T, metaDataProvider *MetaDataProvider,
	shoppingCard *metadata.ShoppingCartExtension) {
	serverMetaInfo := metaDataProvider.serverMetaInfo

	if serverMetaInfo.ShoppingCartExtension.Creator != shoppingCard.Creator {
		t.Fatalf("[%s]AssertShoppingCard : Creator '%s' != '%s'", prefix,
			serverMetaInfo.ShoppingCartExtension.Creator, shoppingCard.Creator)
	}
	if serverMetaInfo.ShoppingCartExtension.Name != shoppingCard.Name {
		t.Fatalf("[%s]AssertShoppingCard : Name '%s' != '%s'", prefix,
			serverMetaInfo.ShoppingCartExtension.Name, shoppingCard.Name)
	}
	if serverMetaInfo.ShoppingCartExtension.Version != shoppingCard.Version {
		t.Fatalf("[%s]AssertShoppingCard : Version '%s' != '%s'", prefix,
			serverMetaInfo.ShoppingCartExtension.Version, shoppingCard.Version)
	}
}
