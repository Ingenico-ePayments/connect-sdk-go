// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package token

// MandateNonSepaDirectDebit represents class MandateNonSepaDirectDebit
type MandateNonSepaDirectDebit struct {
	PaymentProduct705SpecificData *NonSepaDirectDebitPaymentProduct705SpecificData `json:"paymentProduct705SpecificData,omitempty"`
	PaymentProduct730SpecificData *NonSepaDirectDebitPaymentProduct730SpecificData `json:"paymentProduct730SpecificData,omitempty"`
}

// NewMandateNonSepaDirectDebit constructs a new MandateNonSepaDirectDebit
func NewMandateNonSepaDirectDebit() *MandateNonSepaDirectDebit {
	return &MandateNonSepaDirectDebit{}
}
