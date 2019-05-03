// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RiskAssessmentBankAccount represents class RiskAssessmentBankAccount
type RiskAssessmentBankAccount struct {
	BankAccountBban  *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban  *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	FraudFields      *definitions.FraudFields     `json:"fraudFields,omitempty"`
	Merchant         *MerchantRiskAssessment      `json:"merchant,omitempty"`
	Order            *OrderRiskAssessment         `json:"order,omitempty"`
	PaymentProductID *int32                       `json:"paymentProductId,omitempty"`
}

// NewRiskAssessmentBankAccount constructs a new RiskAssessmentBankAccount
func NewRiskAssessmentBankAccount() *RiskAssessmentBankAccount {
	return &RiskAssessmentBankAccount{}
}
