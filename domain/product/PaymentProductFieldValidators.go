// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// PaymentProductFieldValidators represents class PaymentProductFieldValidators
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
