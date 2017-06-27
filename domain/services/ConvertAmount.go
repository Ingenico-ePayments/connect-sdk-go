// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package services

// ConvertAmount represents class ConvertAmount
type ConvertAmount struct {
	ConvertedAmount *int64 `json:"convertedAmount,omitempty"`
}

// NewConvertAmount constructs a new ConvertAmount
func NewConvertAmount() *ConvertAmount {
	return &ConvertAmount{}
}
