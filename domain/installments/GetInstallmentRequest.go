// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package installments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// GetInstallmentRequest represents class GetInstallmentRequest
type GetInstallmentRequest struct {
	AmountOfMoney    *definitions.AmountOfMoney `json:"amountOfMoney,omitempty"`
	Bin              *string                    `json:"bin,omitempty"`
	CountryCode      *string                    `json:"countryCode,omitempty"`
	PaymentProductID *int32                     `json:"paymentProductId,omitempty"`
}

// NewGetInstallmentRequest constructs a new GetInstallmentRequest
func NewGetInstallmentRequest() *GetInstallmentRequest {
	return &GetInstallmentRequest{}
}
