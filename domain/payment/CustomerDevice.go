// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package payment

// CustomerDevice represents class CustomerDevice
type CustomerDevice struct {
	AcceptHeader                   *string      `json:"acceptHeader,omitempty"`
	BrowserData                    *BrowserData `json:"browserData,omitempty"`
	DefaultFormFill                *string      `json:"defaultFormFill,omitempty"`
	DeviceFingerprintTransactionID *string      `json:"deviceFingerprintTransactionId,omitempty"`
	IPAddress                      *string      `json:"ipAddress,omitempty"`
	Locale                         *string      `json:"locale,omitempty"`
	TimezoneOffsetUtcMinutes       *string      `json:"timezoneOffsetUtcMinutes,omitempty"`
	UserAgent                      *string      `json:"userAgent,omitempty"`
}

// NewCustomerDevice constructs a new CustomerDevice
func NewCustomerDevice() *CustomerDevice {
	return &CustomerDevice{}
}
