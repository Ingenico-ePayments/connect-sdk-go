// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// AmountBreakdown represents class AmountBreakdown
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AmountBreakdown
type AmountBreakdown struct {
	Amount *int64  `json:"amount,omitempty"`
	Type   *string `json:"type,omitempty"`
}

// NewAmountBreakdown constructs a new AmountBreakdown
func NewAmountBreakdown() *AmountBreakdown {
	return &AmountBreakdown{}
}
