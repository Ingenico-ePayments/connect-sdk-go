// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

import (
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/product"
)

// MerchantAction represents class MerchantAction
type MerchantAction struct {
	ActionType                            *string                                `json:"actionType,omitempty"`
	FormFields                            *[]product.PaymentProductField         `json:"formFields,omitempty"`
	MobileThreeDSecureChallengeParameters *MobileThreeDSecureChallengeParameters `json:"mobileThreeDSecureChallengeParameters,omitempty"`
	RedirectData                          *RedirectData                          `json:"redirectData,omitempty"`
	RenderingData                         *string                                `json:"renderingData,omitempty"`
	ShowData                              *[]definitions.KeyValuePair            `json:"showData,omitempty"`
	ThirdPartyData                        *ThirdPartyData                        `json:"thirdPartyData,omitempty"`
}

// NewMerchantAction constructs a new MerchantAction
func NewMerchantAction() *MerchantAction {
	return &MerchantAction{}
}
