// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// TrialInformation represents class TrialInformation
type TrialInformation struct {
	AmountOfMoneyAfterTrial *definitions.AmountOfMoney `json:"amountOfMoneyAfterTrial,omitempty"`
	EndDate                 *string                    `json:"endDate,omitempty"`
	IsRecurring             *bool                      `json:"isRecurring,omitempty"`
	TrialPeriod             *TrialPeriod               `json:"trialPeriod,omitempty"`
	TrialPeriodRecurring    *Frequency                 `json:"trialPeriodRecurring,omitempty"`
}

// NewTrialInformation constructs a new TrialInformation
func NewTrialInformation() *TrialInformation {
	return &TrialInformation{}
}
