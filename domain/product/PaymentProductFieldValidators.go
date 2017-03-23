// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package product

// PaymentProductFieldValidators represents class PaymentProductFieldValidators
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProductFieldValidators
type PaymentProductFieldValidators struct {
	BoletoBancarioRequiredness *BoletoBancarioRequirednessValidator `json:"boletoBancarioRequiredness,omitempty"`
	EmailAddress               *EmptyValidator                      `json:"emailAddress,omitempty"`
	ExpirationDate             *EmptyValidator                      `json:"expirationDate,omitempty"`
	FixedList                  *FixedListValidator                  `json:"fixedList,omitempty"`
	Length                     *LengthValidator                     `json:"length,omitempty"`
	Luhn                       *EmptyValidator                      `json:"luhn,omitempty"`
	Range                      *RangeValidator                      `json:"range,omitempty"`
	RegularExpression          *RegularExpressionValidator          `json:"regularExpression,omitempty"`
}

// NewPaymentProductFieldValidators constructs a new PaymentProductFieldValidators
func NewPaymentProductFieldValidators() *PaymentProductFieldValidators {
	return &PaymentProductFieldValidators{}
}
