// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Product806SpecificOutput represents class PaymentProduct806SpecificOutput
type Product806SpecificOutput struct {
	BillingAddress  *definitions.Address `json:"billingAddress,omitempty"`
	CustomerAccount *TrustlyBankAccount  `json:"customerAccount,omitempty"`
}

// NewProduct806SpecificOutput constructs a new Product806SpecificOutput
func NewProduct806SpecificOutput() *Product806SpecificOutput {
	return &Product806SpecificOutput{}
}
