// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// Installments represents class Installments
type Installments struct {
	AmountOfMoneyPerInstallment *definitions.AmountOfMoney `json:"amountOfMoneyPerInstallment,omitempty"`
	AmountOfMoneyTotal          *definitions.AmountOfMoney `json:"amountOfMoneyTotal,omitempty"`
	FrequencyOfInstallments     *string                    `json:"frequencyOfInstallments,omitempty"`
	InstallmentPlanCode         *int32                     `json:"installmentPlanCode,omitempty"`
	InterestRate                *string                    `json:"interestRate,omitempty"`
	NumberOfInstallments        *int64                     `json:"numberOfInstallments,omitempty"`
}

// NewInstallments constructs a new Installments
func NewInstallments() *Installments {
	return &Installments{}
}
