// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Product840SpecificOutput represents class PaymentProduct840SpecificOutput
type Product840SpecificOutput struct {
	CustomerAccount *Product840CustomerAccount `json:"customerAccount,omitempty"`
	CustomerAddress *definitions.Address       `json:"customerAddress,omitempty"`
}

// NewProduct840SpecificOutput constructs a new Product840SpecificOutput
func NewProduct840SpecificOutput() *Product840SpecificOutput {
	return &Product840SpecificOutput{}
}
