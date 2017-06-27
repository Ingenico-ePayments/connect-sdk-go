// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// NonSepaDirectDebitPaymentProduct705SpecificData represents class TokenNonSepaDirectDebitPaymentProduct705SpecificData
type NonSepaDirectDebitPaymentProduct705SpecificData struct {
	AuthorisationID *string                      `json:"authorisationId,omitempty"`
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
}

// NewNonSepaDirectDebitPaymentProduct705SpecificData constructs a new NonSepaDirectDebitPaymentProduct705SpecificData
func NewNonSepaDirectDebitPaymentProduct705SpecificData() *NonSepaDirectDebitPaymentProduct705SpecificData {
	return &NonSepaDirectDebitPaymentProduct705SpecificData{}
}
