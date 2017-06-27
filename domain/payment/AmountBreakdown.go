// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// AmountBreakdown represents class AmountBreakdown
type AmountBreakdown struct {
	Amount *int64  `json:"amount,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// NewAmountBreakdown constructs a new AmountBreakdown
func NewAmountBreakdown() *AmountBreakdown {
	return &AmountBreakdown{}
}
