// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package token

// EWallet represents class TokenEWallet
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TokenEWallet
type EWallet struct {
	Alias    *string        `json:"alias,omitempty"`
	Customer *CustomerToken `json:"customer,omitempty"`
	Data     *EWalletData   `json:"data,omitempty"`
}

// NewEWallet constructs a new EWallet
func NewEWallet() *EWallet {
	return &EWallet{}
}
