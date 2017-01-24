// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Product840SpecificOutput represents class PaymentProduct840SpecificOutput
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProduct840SpecificOutput
type Product840SpecificOutput struct {
	CustomerAccount *Product840CustomerAccount `json:"customerAccount,omitempty"`
	CustomerAddress *definitions.Address       `json:"customerAddress,omitempty"`
}

// NewProduct840SpecificOutput constructs a new Product840SpecificOutput
func NewProduct840SpecificOutput() *Product840SpecificOutput {
	return &Product840SpecificOutput{}
}
