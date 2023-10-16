// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

// RecurringPaymentsData represents class RecurringPaymentsData
type RecurringPaymentsData struct {
	RecurringInterval *Frequency        `json:"recurringInterval,omitempty"`
	TrialInformation  *TrialInformation `json:"trialInformation,omitempty"`
}

// NewRecurringPaymentsData constructs a new RecurringPaymentsData
func NewRecurringPaymentsData() *RecurringPaymentsData {
	return &RecurringPaymentsData{}
}
