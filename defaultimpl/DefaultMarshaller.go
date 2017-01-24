package defaultimpl

import "encoding/json"

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

// NewDefaultMarshaller creates a DefaultMarshaller
func NewDefaultMarshaller() (*DefaultMarshaller, error) {
	return &DefaultMarshaller{}, nil
}
