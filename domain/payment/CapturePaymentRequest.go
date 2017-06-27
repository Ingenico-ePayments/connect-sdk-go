// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// CapturePaymentRequest represents class CapturePaymentRequest
type CapturePaymentRequest struct {
	Amount  *int64 `json:"amount,omitempty"`
	IsFinal *bool  `json:"isFinal,omitempty"`
}

// NewCapturePaymentRequest constructs a new CapturePaymentRequest
func NewCapturePaymentRequest() *CapturePaymentRequest {
	return &CapturePaymentRequest{}
}
