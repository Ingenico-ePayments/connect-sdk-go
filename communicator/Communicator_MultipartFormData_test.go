package communicator

import (
	"io"
	"net/url"
	"strings"
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/configuration"
	"github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"
)

func TestMultipartFormDataUploadPostMultipartFormDataObjectWithResponse(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = communicator.Post("/post", nil, nil, *multipart, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataObjectPointerWithResponse(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = communicator.Post("/post", nil, nil, multipart, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataRequestWithResponse(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = communicator.Post("/post", nil, nil, &MultipartFormDataObjectWrapper{multipart}, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataObjectWithBodyHandler(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	err = communicator.PostWithHandler("/post", nil, nil, *multipart, nil, BodyHandlerFunc(func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller, err := defaultimpl.NewDefaultMarshaller()
		if err != nil {
			t.Fatal(err)
		}
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	}))
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataObjectPointerWithBodyHandler(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	err = communicator.PostWithHandler("/post", nil, nil, multipart, nil, BodyHandlerFunc(func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller, err := defaultimpl.NewDefaultMarshaller()
		if err != nil {
			t.Fatal(err)
		}
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	}))
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPostMultipartFormDataRequestWithBodyHandler(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	err = communicator.PostWithHandler("/post", nil, nil, &MultipartFormDataObjectWrapper{multipart}, nil, BodyHandlerFunc(func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller, err := defaultimpl.NewDefaultMarshaller()
		if err != nil {
			t.Fatal(err)
		}
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	}))
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataObjectWithResponse(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = communicator.Put("/put", nil, nil, *multipart, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataObjectPointerWithResponse(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = communicator.Put("/put", nil, nil, multipart, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataRequestWithResponse(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	response := HTTPBinResponse{}
	err = communicator.Put("/put", nil, nil, &MultipartFormDataObjectWrapper{multipart}, nil, &response)
	if err != nil {
		t.Fatal(err)
	}

	if len(response.Files) != 1 {
		t.Fatal("Expected 1 file")
	}
	if fileContent, ok := response.Files["file"]; ok {
		if fileContent != expected {
			t.Fatal("Wrong content in file: " + fileContent)
		}
	} else {
		t.Fatal("Expected file with index 'file'")
	}

	if len(response.Form) != 1 {
		t.Fatal("Expected 1 value")
	}
	if valueContent, ok := response.Form["value"]; ok {
		if valueContent != "Hello World" {
			t.Fatal("Wrong content in value: " + valueContent)
		}
	} else {
		t.Fatal("Expected value with index 'value'")
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataObjectWithBodyHandler(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	err = communicator.PutWithHandler("/put", nil, nil, *multipart, nil, BodyHandlerFunc(func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller, err := defaultimpl.NewDefaultMarshaller()
		if err != nil {
			t.Fatal(err)
		}
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	}))
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataObjectPointerWithBodyHandler(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	err = communicator.PutWithHandler("/put", nil, nil, multipart, nil, BodyHandlerFunc(func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller, err := defaultimpl.NewDefaultMarshaller()
		if err != nil {
			t.Fatal(err)
		}
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	}))
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultipartFormDataUploadPutMultipartFormDataRequestWithBodyHandler(t *testing.T) {
	communicator, err := createCommunicator()
	if err != nil {
		t.Fatal(err)
	}
	defer communicator.Close()

	expected := "file-content"

	content := strings.NewReader(expected)

	multipart, err := communication.NewMultipartFormDataObject()
	if err != nil {
		t.Fatal(err)
	}

	file, _ := communication.NewUploadableFile("file.txt", content, "text/plain")
	multipart.AddFile("file", *file)
	multipart.AddValue("value", "Hello World")

	err = communicator.PutWithHandler("/put", nil, nil, &MultipartFormDataObjectWrapper{multipart}, nil, BodyHandlerFunc(func(headers []communication.Header, reader io.Reader) error {
		response := HTTPBinResponse{}
		marshaller, err := defaultimpl.NewDefaultMarshaller()
		if err != nil {
			t.Fatal(err)
		}
		err = marshaller.UnmarshalFromReader(reader, &response)
		if err != nil {
			t.Fatal(err)
		}

		if len(response.Files) != 1 {
			t.Fatal("Expected 1 file")
		}
		if fileContent, ok := response.Files["file"]; ok {
			if fileContent != expected {
				t.Fatal("Wrong content in file: " + fileContent)
			}
		} else {
			t.Fatal("Expected file with index 'file'")
		}

		if len(response.Form) != 1 {
			t.Fatal("Expected 1 value")
		}
		if valueContent, ok := response.Form["value"]; ok {
			if valueContent != "Hello World" {
				t.Fatal("Wrong content in value: " + valueContent)
			}
		} else {
			t.Fatal("Expected value with index 'value'")
		}

		return nil
	}))
	if err != nil {
		t.Fatal(err)
	}
}

func createCommunicator() (*Communicator, error) {
	configuration := configuration.DefaultConfiguration("dummy", "dummy", "ingenico")
	configuration.APIEndpoint = url.URL{Scheme: "http", Host: "httpbin.org"}

	builder, err := NewSessionBuilder()
	if err != nil {
		return nil, err
	}

	connection, err := defaultimpl.NewDefaultConnection(configuration.SocketTimeout,
		configuration.ConnectTimeout,
		configuration.KeepAliveTimeout,
		configuration.IdleTimeout,
		configuration.MaxConnections,
		configuration.Proxy)
	if err != nil {
		return nil, err
	}

	metaDataProviderBuilder := NewMetaDataProviderBuilder(configuration.Integrator)
	metaDataProviderBuilder.ShoppingCartExtension = configuration.ShoppingCartExtension

	metaDataProvider, err := metaDataProviderBuilder.Build()
	if err != nil {
		return nil, err
	}

	authenticator, err := defaultimpl.NewDefaultAuthenticator(configuration.AuthorizationType,
		configuration.APIKeyID,
		configuration.SecretAPIKey)
	if err != nil {
		return nil, err
	}

	builder.WithAPIEndpoint(&configuration.APIEndpoint).WithConnection(connection).WithMetaDataProvider(metaDataProvider).WithAuthenticator(authenticator)

	session, err := builder.Build()
	if err != nil {
		return nil, err
	}

	marshaller, err := defaultimpl.NewDefaultMarshaller()
	if err != nil {
		return nil, err
	}

	return NewCommunicator(session, marshaller)
}

type HTTPBinResponse struct {
	Form  map[string]string `json:"form"`
	Files map[string]string `json:"files"`
}

type MultipartFormDataObjectWrapper struct {
	Multipart *communication.MultipartFormDataObject
}

func (w *MultipartFormDataObjectWrapper) ToMultipartFormDataObject() *communication.MultipartFormDataObject {
	return w.Multipart
}
