// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// Shipping represents class Shipping
type Shipping struct {
	Address          *AddressPersonal `json:"address,omitempty"`
	AddressIndicator *string          `json:"addressIndicator,omitempty"`
	Comments         *string          `json:"comments,omitempty"`
	EmailAddress     *string          `json:"emailAddress,omitempty"`
	FirstUsageDate   *string          `json:"firstUsageDate,omitempty"`
	IsFirstUsage     *bool            `json:"isFirstUsage,omitempty"`
	TrackingNumber   *string          `json:"trackingNumber,omitempty"`
	Type             *string          `json:"type,omitempty"`
}

// NewShipping constructs a new Shipping
func NewShipping() *Shipping {
	return &Shipping{}
}
