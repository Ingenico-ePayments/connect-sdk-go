// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// AdditionalOrderInputAirlineData represents class AdditionalOrderInputAirlineData
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AdditionalOrderInputAirlineData
type AdditionalOrderInputAirlineData struct {
	AirlineData *AirlineData `json:"airlineData,omitempty"`
}

// NewAdditionalOrderInputAirlineData constructs a new AdditionalOrderInputAirlineData
func NewAdditionalOrderInputAirlineData() *AdditionalOrderInputAirlineData {
	return &AdditionalOrderInputAirlineData{}
}
