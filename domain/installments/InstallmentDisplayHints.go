// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package installments

// InstallmentDisplayHints represents class InstallmentDisplayHints
type InstallmentDisplayHints struct {
	DisplayOrder *int32  `json:"displayOrder,omitempty"`
	Label        *string `json:"label,omitempty"`
	Logo         *string `json:"logo,omitempty"`
}

// NewInstallmentDisplayHints constructs a new InstallmentDisplayHints
func NewInstallmentDisplayHints() *InstallmentDisplayHints {
	return &InstallmentDisplayHints{}
}
