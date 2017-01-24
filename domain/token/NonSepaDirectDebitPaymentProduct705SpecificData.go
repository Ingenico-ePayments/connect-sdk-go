// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// NonSepaDirectDebitPaymentProduct705SpecificData represents class TokenNonSepaDirectDebitPaymentProduct705SpecificData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenNonSepaDirectDebitPaymentProduct705SpecificData
type NonSepaDirectDebitPaymentProduct705SpecificData struct {
	AuthorisationID *string                      `json:"authorisationId,omitempty"`
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
}

// NewNonSepaDirectDebitPaymentProduct705SpecificData constructs a new NonSepaDirectDebitPaymentProduct705SpecificData
func NewNonSepaDirectDebitPaymentProduct705SpecificData() *NonSepaDirectDebitPaymentProduct705SpecificData {
	return &NonSepaDirectDebitPaymentProduct705SpecificData{}
}
