// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// Merchant represents class Merchant
type Merchant struct {
	ConfigurationID   *string `json:"configurationId,omitempty"`
	ContactWebsiteURL *string `json:"contactWebsiteUrl,omitempty"`
	Seller            *Seller `json:"seller,omitempty"`
	WebsiteURL        *string `json:"websiteUrl,omitempty"`
}

// NewMerchant constructs a new Merchant
func NewMerchant() *Merchant {
	return &Merchant{}
}
