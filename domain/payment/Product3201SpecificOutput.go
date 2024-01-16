// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Product3201SpecificOutput represents class PaymentProduct3201SpecificOutput
type Product3201SpecificOutput struct {
	Card *definitions.CardEssentials `json:"card,omitempty"`
}

// NewProduct3201SpecificOutput constructs a new Product3201SpecificOutput
func NewProduct3201SpecificOutput() *Product3201SpecificOutput {
	return &Product3201SpecificOutput{}
}
