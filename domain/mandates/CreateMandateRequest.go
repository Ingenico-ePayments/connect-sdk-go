// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package mandates

// CreateMandateRequest represents class CreateMandateRequest
type CreateMandateRequest struct {
	Alias                  *string          `json:"alias,omitempty"`
	Customer               *MandateCustomer `json:"customer,omitempty"`
	CustomerReference      *string          `json:"customerReference,omitempty"`
	Language               *string          `json:"language,omitempty"`
	RecurrenceType         *string          `json:"recurrenceType,omitempty"`
	ReturnURL              *string          `json:"returnUrl,omitempty"`
	SignatureType          *string          `json:"signatureType,omitempty"`
	UniqueMandateReference *string          `json:"uniqueMandateReference,omitempty"`
}

// NewCreateMandateRequest constructs a new CreateMandateRequest
func NewCreateMandateRequest() *CreateMandateRequest {
	return &CreateMandateRequest{}
}
