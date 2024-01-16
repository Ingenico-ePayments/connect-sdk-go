// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package hostedmandatemanagement

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/mandates"

// HostedMandateInfo represents class HostedMandateInfo
type HostedMandateInfo struct {
	Alias                  *string                   `json:"alias,omitempty"`
	Customer               *mandates.MandateCustomer `json:"customer,omitempty"`
	CustomerReference      *string                   `json:"customerReference,omitempty"`
	RecurrenceType         *string                   `json:"recurrenceType,omitempty"`
	SignatureType          *string                   `json:"signatureType,omitempty"`
	UniqueMandateReference *string                   `json:"uniqueMandateReference,omitempty"`
}

// NewHostedMandateInfo constructs a new HostedMandateInfo
func NewHostedMandateInfo() *HostedMandateInfo {
	return &HostedMandateInfo{}
}
