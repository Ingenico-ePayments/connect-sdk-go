// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package riskassessments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// RiskAssessmentBankAccount represents class RiskAssessmentBankAccount
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_RiskAssessmentBankAccount
type RiskAssessmentBankAccount struct {
	BankAccountBban  *definitions.BankAccountBban `json:"bankAccountBban,omitempty"`
	BankAccountIban  *definitions.BankAccountIban `json:"bankAccountIban,omitempty"`
	FraudFields      *definitions.FraudFields     `json:"fraudFields,omitempty"`
	Order            *OrderRiskAssessment         `json:"order,omitempty"`
	PaymentProductID *int32                       `json:"paymentProductId,omitempty"`
}

// NewRiskAssessmentBankAccount constructs a new RiskAssessmentBankAccount
func NewRiskAssessmentBankAccount() *RiskAssessmentBankAccount {
	return &RiskAssessmentBankAccount{}
}
