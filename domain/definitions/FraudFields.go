// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// FraudFields represents class FraudFields
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_FraudFields
type FraudFields struct {
	CustomerIPAddress      *string   `json:"customerIpAddress,omitempty"`
	DefaultFormFill        *string   `json:"defaultFormFill,omitempty"`
	GiftCardType           *string   `json:"giftCardType,omitempty"`
	GiftMessage            *string   `json:"giftMessage,omitempty"`
	HasForgottenPwd        *bool     `json:"hasForgottenPwd,omitempty"`
	HasPassword            *bool     `json:"hasPassword,omitempty"`
	IsPreviousCustomer     *bool     `json:"isPreviousCustomer,omitempty"`
	OrderTimezone          *string   `json:"orderTimezone,omitempty"`
	ShipComments           *string   `json:"shipComments,omitempty"`
	ShipmentTrackingNumber *string   `json:"shipmentTrackingNumber,omitempty"`
	UserData               *[]string `json:"userData,omitempty"`
	Website                *string   `json:"website,omitempty"`
}

// NewFraudFields constructs a new FraudFields
func NewFraudFields() *FraudFields {
	return &FraudFields{}
}
