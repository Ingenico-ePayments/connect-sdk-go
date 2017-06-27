// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

// EWalletData represents class TokenEWalletData
type EWalletData struct {
	BillingAgreementID *string `json:"billingAgreementId,omitempty"`
}

// NewEWalletData constructs a new EWalletData
func NewEWalletData() *EWalletData {
	return &EWalletData{}
}
