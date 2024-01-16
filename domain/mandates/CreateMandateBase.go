// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package mandates

// CreateMandateBase represents class CreateMandateBase
type CreateMandateBase struct {
	Alias                  *string          `json:"alias,omitempty"`
	Customer               *MandateCustomer `json:"customer,omitempty"`
	CustomerReference      *string          `json:"customerReference,omitempty"`
	Language               *string          `json:"language,omitempty"`
	RecurrenceType         *string          `json:"recurrenceType,omitempty"`
	SignatureType          *string          `json:"signatureType,omitempty"`
	UniqueMandateReference *string          `json:"uniqueMandateReference,omitempty"`
}

// NewCreateMandateBase constructs a new CreateMandateBase
func NewCreateMandateBase() *CreateMandateBase {
	return &CreateMandateBase{}
}
