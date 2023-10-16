// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package installments

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"

// InstallmentOptions represents class InstallmentOptions
type InstallmentOptions struct {
	DisplayHints     *InstallmentDisplayHints `json:"displayHints,omitempty"`
	ID               *string                  `json:"id,omitempty"`
	InstallmentPlans *[]payment.Installments  `json:"installmentPlans,omitempty"`
}

// NewInstallmentOptions constructs a new InstallmentOptions
func NewInstallmentOptions() *InstallmentOptions {
	return &InstallmentOptions{}
}
