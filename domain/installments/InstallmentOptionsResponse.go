// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package installments

// InstallmentOptionsResponse represents class InstallmentOptionsResponse
type InstallmentOptionsResponse struct {
	InstallmentOptions *[]InstallmentOptions `json:"installmentOptions,omitempty"`
}

// NewInstallmentOptionsResponse constructs a new InstallmentOptionsResponse
func NewInstallmentOptionsResponse() *InstallmentOptionsResponse {
	return &InstallmentOptionsResponse{}
}
