// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// FraudFields represents class FraudFields
type FraudFields struct {
	AddressesAreIdentical          *bool                       `json:"addressesAreIdentical,omitempty"`
	BlackListData                  *string                     `json:"blackListData,omitempty"`
	CardOwnerAddress               *Address                    `json:"cardOwnerAddress,omitempty"`
	CustomerIPAddress              *string                     `json:"customerIpAddress,omitempty"`
	DefaultFormFill                *string                     `json:"defaultFormFill,omitempty"`
	DeviceFingerprintActivated     *bool                       `json:"deviceFingerprintActivated,omitempty"`
	DeviceFingerprintTransactionID *string                     `json:"deviceFingerprintTransactionId,omitempty"`
	GiftCardType                   *string                     `json:"giftCardType,omitempty"`
	GiftMessage                    *string                     `json:"giftMessage,omitempty"`
	HasForgottenPwd                *bool                       `json:"hasForgottenPwd,omitempty"`
	HasPassword                    *bool                       `json:"hasPassword,omitempty"`
	IsPreviousCustomer             *bool                       `json:"isPreviousCustomer,omitempty"`
	OrderTimezone                  *string                     `json:"orderTimezone,omitempty"`
	ShipComments                   *string                     `json:"shipComments,omitempty"`
	ShipmentTrackingNumber         *string                     `json:"shipmentTrackingNumber,omitempty"`
	ShippingDetails                *FraudFieldsShippingDetails `json:"shippingDetails,omitempty"`
	UserData                       *[]string                   `json:"userData,omitempty"`
	Website                        *string                     `json:"website,omitempty"`
}

// NewFraudFields constructs a new FraudFields
func NewFraudFields() *FraudFields {
	return &FraudFields{}
}
