// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

// PaymentProductGroupResponse represents class PaymentProductGroupResponse
type PaymentProductGroupResponse struct {
	AccountsOnFile           *[]AccountOnFile            `json:"accountsOnFile,omitempty"`
	AllowsInstallments       *bool                       `json:"allowsInstallments,omitempty"`
	DeviceFingerprintEnabled *bool                       `json:"deviceFingerprintEnabled,omitempty"`
	DisplayHints             *PaymentProductDisplayHints `json:"displayHints,omitempty"`
	Fields                   *[]PaymentProductField      `json:"fields,omitempty"`
	ID                       *string                     `json:"id,omitempty"`
}

// NewPaymentProductGroupResponse constructs a new PaymentProductGroupResponse
func NewPaymentProductGroupResponse() *PaymentProductGroupResponse {
	return &PaymentProductGroupResponse{}
}
