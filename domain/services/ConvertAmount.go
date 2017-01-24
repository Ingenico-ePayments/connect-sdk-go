// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package services

// ConvertAmount represents class ConvertAmount
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_ConvertAmount
type ConvertAmount struct {
	ConvertedAmount *int64 `json:"convertedAmount,omitempty"`
}

// NewConvertAmount constructs a new ConvertAmount
func NewConvertAmount() *ConvertAmount {
	return &ConvertAmount{}
}
