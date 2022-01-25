// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// SettlementDetails represents class SettlementDetails
type SettlementDetails struct {
	AcquirerReferenceNumber  *string                    `json:"acquirerReferenceNumber,omitempty"`
	AmountOfMoney            *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	PaymentID                *string                    `json:"paymentId,omitempty"`
	RetrievalReferenceNumber *string                    `json:"retrievalReferenceNumber,omitempty"`
}

// NewSettlementDetails constructs a new SettlementDetails
func NewSettlementDetails() *SettlementDetails {
	return &SettlementDetails{}
}
