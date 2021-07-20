// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// BankTransferPayoutMethodSpecificInput represents class BankTransferPayoutMethodSpecificInput
type BankTransferPayoutMethodSpecificInput struct {
	BankAccountBban *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	// Deprecated: Moved to PayoutDetails
	Customer        *Customer                    `json:"customer,omitempty"`
	PayoutDate      *string                      `json:"payoutDate,omitempty"`
	PayoutText      *string                      `json:"payoutText,omitempty"`
	SwiftCode       *string                      `json:"swiftCode,omitempty"`
}

// NewBankTransferPayoutMethodSpecificInput constructs a new BankTransferPayoutMethodSpecificInput
func NewBankTransferPayoutMethodSpecificInput() *BankTransferPayoutMethodSpecificInput {
	return &BankTransferPayoutMethodSpecificInput{}
}
