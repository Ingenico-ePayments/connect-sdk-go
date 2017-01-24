// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductGroupResponse represents class PaymentProductGroupResponse
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductGroupResponse
type PaymentProductGroupResponse struct {
	AccountsOnFile *[]AccountOnFile            `json:"accountsOnFile,omitempty"`
	DisplayHints   *PaymentProductDisplayHints `json:"displayHints,omitempty"`
	Fields         *[]PaymentProductField      `json:"fields,omitempty"`
	ID             *string                     `json:"id,omitempty"`
}

// NewPaymentProductGroupResponse constructs a new PaymentProductGroupResponse
func NewPaymentProductGroupResponse() *PaymentProductGroupResponse {
	return &PaymentProductGroupResponse{}
}
