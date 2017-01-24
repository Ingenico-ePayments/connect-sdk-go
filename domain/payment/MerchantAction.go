// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// MerchantAction represents class MerchantAction
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_MerchantAction
type MerchantAction struct {
	ActionType    *string                     `json:"actionType,omitempty"`
	RedirectData  *RedirectData               `json:"redirectData,omitempty"`
	RenderingData *string                     `json:"renderingData,omitempty"`
	ShowData      *[]definitions.KeyValuePair `json:"showData,omitempty"`
}

// NewMerchantAction constructs a new MerchantAction
func NewMerchantAction() *MerchantAction {
	return &MerchantAction{}
}
