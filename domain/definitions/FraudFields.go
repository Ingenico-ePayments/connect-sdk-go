// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// FraudFields represents class FraudFields
type FraudFields struct {
	// Deprecated: For risk assessments there is no replacement. For other calls, use Order.shipping.addressIndicator instead
	AddressesAreIdentical          *bool                       `json:"addressesAreIdentical,omitempty"`
	BlackListData                  *string                     `json:"blackListData,omitempty"`
	// Deprecated: This should be the same as Order.customer.billingAddress
	CardOwnerAddress               *Address                    `json:"cardOwnerAddress,omitempty"`
	CustomerIPAddress              *string                     `json:"customerIpAddress,omitempty"`
	// Deprecated: Use Order.customer.device.defaultFormFill instead
	DefaultFormFill                *string                     `json:"defaultFormFill,omitempty"`
	// Deprecated: No replacement
	DeviceFingerprintActivated     *bool                       `json:"deviceFingerprintActivated,omitempty"`
	// Deprecated: Use Order.customer.device.deviceFingerprintTransactionId instead
	DeviceFingerprintTransactionID *string                     `json:"deviceFingerprintTransactionId,omitempty"`
	GiftCardType                   *string                     `json:"giftCardType,omitempty"`
	GiftMessage                    *string                     `json:"giftMessage,omitempty"`
	// Deprecated: Use Order.customer.account.hasForgottenPassword instead
	HasForgottenPwd                *bool                       `json:"hasForgottenPwd,omitempty"`
	// Deprecated: Use Order.customer.account.hasPassword instead
	HasPassword                    *bool                       `json:"hasPassword,omitempty"`
	// Deprecated: Use Order.customer.isPreviousCustomer instead
	IsPreviousCustomer             *bool                       `json:"isPreviousCustomer,omitempty"`
	OrderTimezone                  *string                     `json:"orderTimezone,omitempty"`
	// Deprecated: Use Order.shipping.comments instead
	ShipComments                   *string                     `json:"shipComments,omitempty"`
	// Deprecated: Use Order.shipping.trackingNumber instead
	ShipmentTrackingNumber         *string                     `json:"shipmentTrackingNumber,omitempty"`
	// Deprecated: No replacement
	ShippingDetails                *FraudFieldsShippingDetails `json:"shippingDetails,omitempty"`
	UserData                       *[]string                   `json:"userData,omitempty"`
	// Deprecated: Use Merchant.websiteUrl instead
	Website                        *string                     `json:"website,omitempty"`
}

// NewFraudFields constructs a new FraudFields
func NewFraudFields() *FraudFields {
	return &FraudFields{}
}
