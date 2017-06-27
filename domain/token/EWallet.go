// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package token

// EWallet represents class TokenEWallet
type EWallet struct {
	Alias    *string        `json:"alias,omitempty"`
	Customer *CustomerToken `json:"customer,omitempty"`
	Data     *EWalletData   `json:"data,omitempty"`
}

// NewEWallet constructs a new EWallet
func NewEWallet() *EWallet {
	return &EWallet{}
}
