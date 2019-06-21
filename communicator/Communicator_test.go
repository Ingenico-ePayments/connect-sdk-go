package communicator

import (
	"io"
	"net/url"
	"testing"
)

var baseURL = url.URL{
	Scheme: "https",
	Host:   "eu.sandbox.api-ingenico.com",
}

var marshaller = &communicatorTestMarshaller{}

type communicatorTestMarshaller struct {
}

// Marshal encodes the given value using json.Marshal
func (m *communicatorTestMarshaller) Marshal(v interface{}) (string, error) {
	panic("Not implemented")
}

func (m *communicatorTestMarshaller) UnmarshalFromReader(reader io.Reader, v interface{}) error {
	panic("Not implemented")
}

// Unmarshal decodes the given data into the given value using json.Unmarshal
func (m *communicatorTestMarshaller) Unmarshal(data string, v interface{}) error {
	panic("Not implemented")
}

var mockSession = &Session{
	apiEndpoint: &baseURL,
}

func TestToURIWithoutRequestParams(t *testing.T) {
	communicator, err := NewCommunicator(mockSession, marshaller)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}

	expectedURL, err := url.Parse("https://eu.sandbox.api-ingenico.com/v1/merchant/20000/convertamount")
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}

	url1, err := communicator.toAbsoluteURI("v1/merchant/20000/convertamount", RequestParams{})
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if url1 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", url1, *expectedURL)
	}

	url2, err := communicator.toAbsoluteURI("/v1/merchant/20000/convertamount", RequestParams{})
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if url2 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", url2, *expectedURL)
	}
}

func TestToURIWithRequestParams(t *testing.T) {
	params := append(RequestParams{},
		RequestParam{"amount", "123"},
		RequestParam{"source", "USD"},
		RequestParam{"target", "EUR"},
		RequestParam{"dummy", "é&%="})

	communicator, err := NewCommunicator(mockSession, marshaller)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}

	expectedURL, err := url.Parse("https://eu.sandbox.api-ingenico.com/v1/merchant/20000/convertamount?amount=123&source=USD&target=EUR&dummy=%C3%A9%26%25%3D")
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}

	url1, err := communicator.toAbsoluteURI("v1/merchant/20000/convertamount", params)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if url1 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", url1, *expectedURL)
	}

	url2, err := communicator.toAbsoluteURI("/v1/merchant/20000/convertamount", params)
	if err != nil {
		t.Fatalf("TestToURIWithoutRequestParams: %v", err)
	}
	if url2 != *expectedURL {
		t.Fatalf("TestToURIWithoutRequestParams: url mismatch '%v' '%v'", url2, *expectedURL)
	}
}
