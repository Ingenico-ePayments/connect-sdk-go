// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// GetCustomerDetailsResponse represents class GetCustomerDetailsResponse
type GetCustomerDetailsResponse struct {
	City         *string `json:"city,omitempty"`
	Country      *string `json:"country,omitempty"`
	EmailAddress *string `json:"emailAddress,omitempty"`
	FirstName    *string `json:"firstName,omitempty"`
	FiscalNumber *string `json:"fiscalNumber,omitempty"`
	LanguageCode *string `json:"languageCode,omitempty"`
	PhoneNumber  *string `json:"phoneNumber,omitempty"`
	Street       *string `json:"street,omitempty"`
	Surname      *string `json:"surname,omitempty"`
	Zip          *string `json:"zip,omitempty"`
}

// NewGetCustomerDetailsResponse constructs a new GetCustomerDetailsResponse
func NewGetCustomerDetailsResponse() *GetCustomerDetailsResponse {
	return &GetCustomerDetailsResponse{}
}
