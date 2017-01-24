// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package payment

// Product840CustomerAccount represents class PaymentProduct840CustomerAccount
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_PaymentProduct840CustomerAccount
type Product840CustomerAccount struct {
	AccountID             *string `json:"accountId,omitempty"`
	BillingAgreementID    *string `json:"billingAgreementId,omitempty"`
	CompanyName           *string `json:"companyName,omitempty"`
	CountryCode           *string `json:"countryCode,omitempty"`
	CustomerAccountStatus *string `json:"customerAccountStatus,omitempty"`
	CustomerAddressStatus *string `json:"customerAddressStatus,omitempty"`
	FirstName             *string `json:"firstName,omitempty"`
	PayerID               *string `json:"payerId,omitempty"`
	Surname               *string `json:"surname,omitempty"`
}

// NewProduct840CustomerAccount constructs a new Product840CustomerAccount
func NewProduct840CustomerAccount() *Product840CustomerAccount {
	return &Product840CustomerAccount{}
}
