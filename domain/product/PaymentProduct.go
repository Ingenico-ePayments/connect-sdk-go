// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProduct represents class PaymentProduct
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProduct
type PaymentProduct struct {
	AccountsOnFile            *[]AccountOnFile            `json:"accountsOnFile,omitempty"`
	AllowsRecurring           *bool                       `json:"allowsRecurring,omitempty"`
	AllowsTokenization        *bool                       `json:"allowsTokenization,omitempty"`
	AutoTokenized             *bool                       `json:"autoTokenized,omitempty"`
	CanBeIframed              *bool                       `json:"canBeIframed,omitempty"`
	DisplayHints              *PaymentProductDisplayHints `json:"displayHints,omitempty"`
	Fields                    *[]PaymentProductField      `json:"fields,omitempty"`
	ID                        *int32                      `json:"id,omitempty"`
	MaxAmount                 *int64                      `json:"maxAmount,omitempty"`
	MinAmount                 *int64                      `json:"minAmount,omitempty"`
	MobileIntegrationLevel    *string                     `json:"mobileIntegrationLevel,omitempty"`
	PaymentMethod             *string                     `json:"paymentMethod,omitempty"`
	PaymentProductGroup       *string                     `json:"paymentProductGroup,omitempty"`
	UsesRedirectionTo3rdParty *bool                       `json:"usesRedirectionTo3rdParty,omitempty"`
}

// NewPaymentProduct constructs a new PaymentProduct
func NewPaymentProduct() *PaymentProduct {
	return &PaymentProduct{}
}
