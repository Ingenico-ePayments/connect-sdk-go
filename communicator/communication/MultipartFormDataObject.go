package communication

import (
	"crypto/rand"
	"errors"
	"fmt"
	"strings"
)

// MultipartFormDataObject is a representation of a multipart/form-data object.
type MultipartFormDataObject struct {
	boundary string
	contentType string
	values map[string]string
	files map[string]UploadableFile
}

// GetBoundary returns the boundary.
func (multipart *MultipartFormDataObject) GetBoundary() string {
	return multipart.boundary
}

// GetContentType returns the content type.
func (multipart *MultipartFormDataObject) GetContentType() string {
	return multipart.contentType
}

// GetValues returns the values.
func (multipart *MultipartFormDataObject) GetValues() map[string]string {
	return multipart.values
}

// GetFiles returns the files.
func (multipart *MultipartFormDataObject) GetFiles() map[string]UploadableFile {
	return multipart.files
}

// AddValue adds a value parameter.
func (multipart *MultipartFormDataObject) AddValue(parameterName string, value string) error {
	if strings.TrimSpace(parameterName) == "" {
		return errors.New("parameterName is required")
	}
	if multipart.containsParameter(parameterName) {
		return errors.New("duplicate parameterName: " + parameterName)
	}
	multipart.values[parameterName] = value

	return nil
}

// AddFile adds a file parameter.
func (multipart *MultipartFormDataObject) AddFile(parameterName string, file UploadableFile) error {
	if strings.TrimSpace(parameterName) == "" {
		return errors.New("parameterName is required")
	}
	if multipart.containsParameter(parameterName) {
		return errors.New("duplicate parameterName: " + parameterName)
	}
	multipart.files[parameterName] = file

	return nil
}

func (multipart *MultipartFormDataObject) containsParameter(parameterName string) bool {
	if _, contains := multipart.values[parameterName]; contains {
		return true
	}
	if _, contains := multipart.files[parameterName]; contains {
		return true
	}
	return false
}

// NewMultipartFormDataObject constructs a new MultipartFormDataObject
func NewMultipartFormDataObject() (*MultipartFormDataObject, error) {
	boundary, err := pseudoUUID()
	if err != nil {
		return nil, err
	}

	contentType := "multipart/form-data; boundary=" + boundary
	values := map[string]string{}
	files := map[string]UploadableFile{}

	return &MultipartFormDataObject{boundary, contentType, values, files}, nil
}

func pseudoUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}
