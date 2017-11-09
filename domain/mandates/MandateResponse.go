// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package mandates

// MandateResponse represents class MandateResponse
type MandateResponse struct {
	Customer               *MandateCustomer `json:"customer,omitempty"`
	CustomerReference      *string          `json:"customerReference,omitempty"`
	RecurrenceType         *string          `json:"recurrenceType,omitempty"`
	Status                 *string          `json:"status,omitempty"`
	UniqueMandateReference *string          `json:"uniqueMandateReference,omitempty"`
}

// NewMandateResponse constructs a new MandateResponse
func NewMandateResponse() *MandateResponse {
	return &MandateResponse{}
}
