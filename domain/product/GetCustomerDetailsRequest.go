// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package product

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// GetCustomerDetailsRequest represents class GetCustomerDetailsRequest
type GetCustomerDetailsRequest struct {
	CountryCode *string                     `json:"countryCode,omitempty"`
	Values      *[]definitions.KeyValuePair `json:"values,omitempty"`
}

// NewGetCustomerDetailsRequest constructs a new GetCustomerDetailsRequest
func NewGetCustomerDetailsRequest() *GetCustomerDetailsRequest {
	return &GetCustomerDetailsRequest{}
}
