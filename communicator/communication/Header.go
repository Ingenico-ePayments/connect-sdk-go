package communication

import (
	"errors"
	"regexp"
	"strings"
)

var (
	// ErrNoName occurs when the given name is empty
	ErrNoName = errors.New("name is required")
)

// Header represents a single response header. Immutable.
type Header struct {
	name, value string
}

// NewHeader returns a Header with the given name and value
func NewHeader(name, value string) (*Header, error) {
	if len(name) == 0 {
		return nil, ErrNoName
	}
	return &Header{name, value}, nil
}

// String is the implementation of the Stringer interface
// Format: 'name:value'
func (h Header) String() string {
	return h.name + ":" + h.value
}

// Name returns the name of the header
func (h Header) Name() string {
	return h.name
}

// Value returns the value of the header
func (h Header) Value() string {
	return h.value
}

// Headers represents a slice of Header
type Headers []Header

// GetHeader searches for the headerName in the headers
func (h Headers) GetHeader(headerName string) *Header {
	for _, header := range h {
		if strings.EqualFold(header.Name(), headerName) {
			return &header
		}
	}
	return nil
}

// GetHeaderValue searches for the header name and returns the value as string,
// or the empty string if it doesn't exist. This doesn't break
// HTTP support since headers can't have empty values
func (h Headers) GetHeaderValue(headerName string) string {
	for _, header := range h {
		if strings.EqualFold(header.Name(), headerName) {
			return header.Value()
		}
	}
	return ""
}

var contentDispositionRegex = regexp.MustCompile("(?i)(?:^|;)\\s*filename\\s*=\\s*(.*?)\\s*(?:;|$)")

// GetDispositionFilename returns the content of the filename found in the Content-Disposition header,
// or the empty string if it couldn't have been found.
func (h Headers) GetDispositionFilename() string {
	headerValue := h.GetHeaderValue("Content-Disposition")

	if headerValue == "" {
		return ""
	}

	filenameResult := contentDispositionRegex.FindAllStringSubmatch(headerValue, -1)
	if filenameResult != nil && filenameResult[0] != nil {
		return trimQuotes(filenameResult[0][1])
	}

	return ""
}

// Removes the quotes from the file name
func trimQuotes(filename string) string {
	if len(filename) < 2 {
		return filename
	}

	if (strings.HasPrefix(filename, "\"") && strings.HasSuffix(filename, "\"")) ||
		(strings.HasPrefix(filename, "'") && strings.HasSuffix(filename, "'")) {
		return filename[1 : len(filename)-1]
	}

	return filename
}

// Len represents the length of the slice
func (h Headers) Len() int {
	return len(h)
}

// Swap swaps two elements
func (h Headers) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Less checks if two positions are in lexicographic order
func (h Headers) Less(i, j int) bool {
	return h[i].Name() < h[j].Name()
}
