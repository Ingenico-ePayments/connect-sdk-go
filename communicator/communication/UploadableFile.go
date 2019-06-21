package communication

import (
	"errors"
	"io"
	"strings"
)

// UploadableFile represents a file that can be uploaded.
type UploadableFile struct {
	fileName string
	content io.Reader
	contentType string
	contentLength int64
}

// GetFileName returns the name of the file.
func (f *UploadableFile) GetFileName() string {
	return f.fileName
}

// GetContent returns a reader with the file's content.
func (f *UploadableFile) GetContent() io.Reader {
	return f.content
}

// GetContentType returns the file's content type.
func (f *UploadableFile) GetContentType() string {
	return f.contentType
}

// GetContentLength returns the file's content length, or -1 if not known.
func (f *UploadableFile) GetContentLength() int64 {
	return f.contentLength
}

// NewUploadableFile creates an uploadable file with the given file name, content and content type, and an unspecified content length.
func NewUploadableFile(fileName string, content io.Reader, contentType string) (*UploadableFile, error) {
	return NewUploadableFileWithLength(fileName, content, contentType, -1)
}

// NewUploadableFileWithLength creates an uploadable file with the given file name, content, content type and content length.
func NewUploadableFileWithLength(fileName string, content io.Reader, contentType string, contentLength int64) (*UploadableFile, error) {
	if strings.TrimSpace(fileName) == "" {
		err := errors.New("fileName is required")
		return nil, err
	}
	if content == nil {
		err := errors.New("content is required")
		return nil, err
	}
	if strings.TrimSpace(contentType) == "" {
		err := errors.New("contentType is required")
		return nil, err
	}
	if contentLength < -1 {
		contentLength = -1
	}

	return &UploadableFile{fileName, content, contentType, contentLength}, nil
}
