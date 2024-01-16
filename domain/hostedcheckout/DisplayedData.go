// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package hostedcheckout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// DisplayedData represents class DisplayedData
type DisplayedData struct {
	DisplayedDataType *string                     `json:"displayedDataType,omitempty"`
	RenderingData     *string                     `json:"renderingData,omitempty"`
	ShowData          *[]definitions.KeyValuePair `json:"showData,omitempty"`
}

// NewDisplayedData constructs a new DisplayedData
func NewDisplayedData() *DisplayedData {
	return &DisplayedData{}
}
