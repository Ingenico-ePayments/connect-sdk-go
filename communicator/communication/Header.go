package communication

import "errors"

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
