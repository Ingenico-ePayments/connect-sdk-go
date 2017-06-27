// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// DecryptedPaymentData represents class DecryptedPaymentData
type DecryptedPaymentData struct {
	CardholderName *string `json:"cardholderName,omitempty"`
	Cryptogram     *string `json:"cryptogram,omitempty"`
	Dpan           *string `json:"dpan,omitempty"`
	Eci            *int32  `json:"eci,omitempty"`
	ExpiryDate     *string `json:"expiryDate,omitempty"`
}

// NewDecryptedPaymentData constructs a new DecryptedPaymentData
func NewDecryptedPaymentData() *DecryptedPaymentData {
	return &DecryptedPaymentData{}
}
