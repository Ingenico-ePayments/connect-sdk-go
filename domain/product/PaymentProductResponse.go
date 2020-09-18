// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package product

// PaymentProductResponse represents class PaymentProductResponse
type PaymentProductResponse struct {
	AccountsOnFile                *[]AccountOnFile               `json:"accountsOnFile,omitempty"`
	AcquirerCountry               *string                        `json:"acquirerCountry,omitempty"`
	AllowsInstallments            *bool                          `json:"allowsInstallments,omitempty"`
	AllowsRecurring               *bool                          `json:"allowsRecurring,omitempty"`
	AllowsTokenization            *bool                          `json:"allowsTokenization,omitempty"`
	AuthenticationIndicator       *AuthenticationIndicator       `json:"authenticationIndicator,omitempty"`
	AutoTokenized                 *bool                          `json:"autoTokenized,omitempty"`
	CanBeIframed                  *bool                          `json:"canBeIframed,omitempty"`
	DeviceFingerprintEnabled      *bool                          `json:"deviceFingerprintEnabled,omitempty"`
	DisplayHints                  *PaymentProductDisplayHints    `json:"displayHints,omitempty"`
	Fields                        *[]PaymentProductField         `json:"fields,omitempty"`
	FieldsWarning                 *string                        `json:"fieldsWarning,omitempty"`
	ID                            *int32                         `json:"id,omitempty"`
	IsJavaScriptRequired          *bool                          `json:"isJavaScriptRequired,omitempty"`
	MaxAmount                     *int64                         `json:"maxAmount,omitempty"`
	MinAmount                     *int64                         `json:"minAmount,omitempty"`
	MobileIntegrationLevel        *string                        `json:"mobileIntegrationLevel,omitempty"`
	PaymentMethod                 *string                        `json:"paymentMethod,omitempty"`
	PaymentProduct302SpecificData *PaymentProduct302SpecificData `json:"paymentProduct302SpecificData,omitempty"`
	PaymentProduct320SpecificData *PaymentProduct320SpecificData `json:"paymentProduct320SpecificData,omitempty"`
	PaymentProduct863SpecificData *PaymentProduct863SpecificData `json:"paymentProduct863SpecificData,omitempty"`
	PaymentProductGroup           *string                        `json:"paymentProductGroup,omitempty"`
	UsesRedirectionTo3rdParty     *bool                          `json:"usesRedirectionTo3rdParty,omitempty"`
}

// NewPaymentProductResponse constructs a new PaymentProductResponse
func NewPaymentProductResponse() *PaymentProductResponse {
	return &PaymentProductResponse{}
}
