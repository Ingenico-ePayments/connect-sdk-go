// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductResponse represents class PaymentProductResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductResponse
type PaymentProductResponse struct {
	AccountsOnFile            *[]AccountOnFile            `json:"accountsOnFile,omitempty"`
	AllowsRecurring           *bool                       `json:"allowsRecurring,omitempty"`
	AllowsTokenization        *bool                       `json:"allowsTokenization,omitempty"`
	AutoTokenized             *bool                       `json:"autoTokenized,omitempty"`
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

// NewPaymentProductResponse constructs a new PaymentProductResponse
func NewPaymentProductResponse() *PaymentProductResponse {
	return &PaymentProductResponse{}
}
