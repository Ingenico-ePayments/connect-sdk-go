package defaultimpl

import (
	"encoding/json"
	"io"
)

// DefaultMarshaller represents the default implementation of the JSON Marshaller
type DefaultMarshaller struct {
}

// Marshal encodes the given value using json.Marshal
func (m *DefaultMarshaller) Marshal(v interface{}) (string, error) {
	dataBytes, err := json.Marshal(v)
	data := string(dataBytes)

	return data, err
}

// Unmarshal decodes the given data into the given value using json.Unmarshal
func (m *DefaultMarshaller) Unmarshal(data string, v interface{}) error {
	if len(data) < 1 {
		return nil
	}

	return json.Unmarshal([]byte(data), v)
}

// UnmarshalFromReader decodes the data from the given reader into the given value using json.NewDecoder
func (m *DefaultMarshaller) UnmarshalFromReader(reader io.Reader, v interface{}) error {
	decoder := json.NewDecoder(reader)

	err := decoder.Decode(v)
	if err != io.EOF {
		return err
	}

	return nil
}

// NewDefaultMarshaller creates a DefaultMarshaller
func NewDefaultMarshaller() (*DefaultMarshaller, error) {
	return &DefaultMarshaller{}, nil
}
