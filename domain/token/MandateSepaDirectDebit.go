// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// MandateSepaDirectDebit represents class MandateSepaDirectDebit
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
