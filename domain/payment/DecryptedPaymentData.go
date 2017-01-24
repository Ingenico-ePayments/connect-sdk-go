// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// DecryptedPaymentData represents class DecryptedPaymentData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_DecryptedPaymentData
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
