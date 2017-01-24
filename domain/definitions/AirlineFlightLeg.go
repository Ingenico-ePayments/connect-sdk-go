// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package definitions

// AirlineFlightLeg represents class AirlineFlightLeg
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_AirlineFlightLeg
type AirlineFlightLeg struct {
	AirlineClass   *string `json:"airlineClass,omitempty"`
	ArrivalAirport *string `json:"arrivalAirport,omitempty"`
	CarrierCode    *string `json:"carrierCode,omitempty"`
	Date           *string `json:"date,omitempty"`
	DepartureTime  *string `json:"departureTime,omitempty"`
	Fare           *string `json:"fare,omitempty"`
	FareBasis      *string `json:"fareBasis,omitempty"`
	FlightNumber   *string `json:"flightNumber,omitempty"`
	Number         *int32  `json:"number,omitempty"`
	OriginAirport  *string `json:"originAirport,omitempty"`
	StopoverCode   *string `json:"stopoverCode,omitempty"`
}

// NewAirlineFlightLeg constructs a new AirlineFlightLeg
func NewAirlineFlightLeg() *AirlineFlightLeg {
	return &AirlineFlightLeg{}
}
