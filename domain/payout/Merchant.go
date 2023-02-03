// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payout

// Merchant represents class PayoutMerchant
type Merchant struct {
	ConfigurationID *string `json:"configurationId,omitempty"`
}

// NewMerchant constructs a new Merchant
func NewMerchant() *Merchant {
	return &Merchant{}
}
