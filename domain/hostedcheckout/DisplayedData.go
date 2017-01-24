// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package hostedcheckout

import "github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"

// DisplayedData represents class DisplayedData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_DisplayedData
type DisplayedData struct {
	DisplayedDataType *string                     `json:"displayedDataType,omitempty"`
	RenderingData     *string                     `json:"renderingData,omitempty"`
	ShowData          *[]definitions.KeyValuePair `json:"showData,omitempty"`
}

// NewDisplayedData constructs a new DisplayedData
func NewDisplayedData() *DisplayedData {
	return &DisplayedData{}
}
