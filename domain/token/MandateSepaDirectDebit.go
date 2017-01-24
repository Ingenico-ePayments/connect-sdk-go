// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// MandateSepaDirectDebit represents class MandateSepaDirectDebit
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_MandateSepaDirectDebit
type MandateSepaDirectDebit struct {
	BankAccountIban            *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	Creditor                   *Creditor                    `json:"creditor,omitempty"`
	CustomerContractIdentifier *string                      `json:"customerContractIdentifier,omitempty"`
	Debtor                     *Debtor                      `json:"debtor,omitempty"`
	IsRecurring                *bool                        `json:"isRecurring,omitempty"`
	MandateApproval            *MandateApproval             `json:"mandateApproval,omitempty"`
	MandateID                  *string                      `json:"mandateId,omitempty"`
	PreNotification            *string                      `json:"preNotification,omitempty"`
}

// NewMandateSepaDirectDebit constructs a new MandateSepaDirectDebit
func NewMandateSepaDirectDebit() *MandateSepaDirectDebit {
	return &MandateSepaDirectDebit{}
}
