// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package mandates

// CreateMandateWithReturnURL represents class CreateMandateWithReturnUrl
type CreateMandateWithReturnURL struct {
	Customer          *MandateCustomer `json:"customer,omitempty"`
	CustomerReference *string          `json:"customerReference,omitempty"`
	Language          *string          `json:"language,omitempty"`
	RecurrenceType    *string          `json:"recurrenceType,omitempty"`
	ReturnURL         *string          `json:"returnUrl,omitempty"`
	SignatureType     *string          `json:"signatureType,omitempty"`
}

// NewCreateMandateWithReturnURL constructs a new CreateMandateWithReturnURL
func NewCreateMandateWithReturnURL() *CreateMandateWithReturnURL {
	return &CreateMandateWithReturnURL{}
}
