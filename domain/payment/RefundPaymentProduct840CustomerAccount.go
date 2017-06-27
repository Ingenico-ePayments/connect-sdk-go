// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// RefundPaymentProduct840CustomerAccount represents class RefundPaymentProduct840CustomerAccount
type RefundPaymentProduct840CustomerAccount struct {
	CustomerAccountStatus *string `json:"customerAccountStatus,omitempty"`
	CustomerAddressStatus *string `json:"customerAddressStatus,omitempty"`
	PayerID               *string `json:"payerId,omitempty"`
}

// NewRefundPaymentProduct840CustomerAccount constructs a new RefundPaymentProduct840CustomerAccount
func NewRefundPaymentProduct840CustomerAccount() *RefundPaymentProduct840CustomerAccount {
	return &RefundPaymentProduct840CustomerAccount{}
}
