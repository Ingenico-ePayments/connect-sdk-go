// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// EWalletData represents class TokenEWalletData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenEWalletData
type EWalletData struct {
	BillingAgreementID *string `json:"billingAgreementId,omitempty"`
}

// NewEWalletData constructs a new EWalletData
func NewEWalletData() *EWalletData {
	return &EWalletData{}
}
