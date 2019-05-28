// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// Product863ThirdPartyData represents class PaymentProduct863ThirdPartyData
type Product863ThirdPartyData struct {
	AppID       *string `json:"appId,omitempty"`
	NonceStr    *string `json:"nonceStr,omitempty"`
	PackageSign *string `json:"packageSign,omitempty"`
	PaySign     *string `json:"paySign,omitempty"`
	PrepayID    *string `json:"prepayId,omitempty"`
	SignType    *string `json:"signType,omitempty"`
	TimeStamp   *string `json:"timeStamp,omitempty"`
}

// NewProduct863ThirdPartyData constructs a new Product863ThirdPartyData
func NewProduct863ThirdPartyData() *Product863ThirdPartyData {
	return &Product863ThirdPartyData{}
}
