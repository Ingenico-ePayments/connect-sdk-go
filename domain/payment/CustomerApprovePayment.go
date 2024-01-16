// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

// CustomerApprovePayment represents class CustomerApprovePayment
type CustomerApprovePayment struct {
	AccountType *string `json:"accountType,omitempty"`
}

// NewCustomerApprovePayment constructs a new CustomerApprovePayment
func NewCustomerApprovePayment() *CustomerApprovePayment {
	return &CustomerApprovePayment{}
}
