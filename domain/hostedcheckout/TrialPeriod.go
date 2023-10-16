// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package hostedcheckout

// TrialPeriod represents class TrialPeriod
type TrialPeriod struct {
	Duration *int32  `json:"duration,omitempty"`
	Interval *string `json:"interval,omitempty"`
}

// NewTrialPeriod constructs a new TrialPeriod
func NewTrialPeriod() *TrialPeriod {
	return &TrialPeriod{}
}
