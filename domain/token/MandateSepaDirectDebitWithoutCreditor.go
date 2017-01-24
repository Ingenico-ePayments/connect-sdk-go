// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// MandateSepaDirectDebitWithoutCreditor represents class MandateSepaDirectDebitWithoutCreditor
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_MandateSepaDirectDebitWithoutCreditor
type MandateSepaDirectDebitWithoutCreditor struct {
	BankAccountIban            *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	CustomerContractIdentifier *string                      `json:"customerContractIdentifier,omitempty"`
	Debtor                     *Debtor                      `json:"debtor,omitempty"`
	IsRecurring                *bool                        `json:"isRecurring,omitempty"`
	MandateApproval            *MandateApproval             `json:"mandateApproval,omitempty"`
	PreNotification            *string                      `json:"preNotification,omitempty"`
}

// NewMandateSepaDirectDebitWithoutCreditor constructs a new MandateSepaDirectDebitWithoutCreditor
func NewMandateSepaDirectDebitWithoutCreditor() *MandateSepaDirectDebitWithoutCreditor {
	return &MandateSepaDirectDebitWithoutCreditor{}
}
