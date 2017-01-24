// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// MandateNonSepaDirectDebit represents class MandateNonSepaDirectDebit
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_MandateNonSepaDirectDebit
type MandateNonSepaDirectDebit struct {
	PaymentProduct705SpecificData *NonSepaDirectDebitPaymentProduct705SpecificData `json:"paymentProduct705SpecificData,omitempty"`
}

// NewMandateNonSepaDirectDebit constructs a new MandateNonSepaDirectDebit
func NewMandateNonSepaDirectDebit() *MandateNonSepaDirectDebit {
	return &MandateNonSepaDirectDebit{}
}
