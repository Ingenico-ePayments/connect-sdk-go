// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package mandates

// MandateMerchantAction represents class MandateMerchantAction
type MandateMerchantAction struct {
	ActionType   *string              `json:"actionType,omitempty"`
	RedirectData *MandateRedirectData `json:"redirectData,omitempty"`
}

// NewMandateMerchantAction constructs a new MandateMerchantAction
func NewMandateMerchantAction() *MandateMerchantAction {
	return &MandateMerchantAction{}
}
