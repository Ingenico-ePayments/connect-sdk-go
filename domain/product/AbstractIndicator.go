// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// AbstractIndicator represents class AbstractIndicator
type AbstractIndicator struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewAbstractIndicator constructs a new AbstractIndicator
func NewAbstractIndicator() *AbstractIndicator {
	return &AbstractIndicator{}
}
