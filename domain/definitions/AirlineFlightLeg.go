// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// AirlineFlightLeg represents class AirlineFlightLeg
type AirlineFlightLeg struct {
	AirlineClass             *string `json:"airlineClass,omitempty"`
	ArrivalAirport           *string `json:"arrivalAirport,omitempty"`
	ArrivalTime              *string `json:"arrivalTime,omitempty"`
	CarrierCode              *string `json:"carrierCode,omitempty"`
	ConjunctionTicket        *string `json:"conjunctionTicket,omitempty"`
	CouponNumber             *string `json:"couponNumber,omitempty"`
	Date                     *string `json:"date,omitempty"`
	DepartureTime            *string `json:"departureTime,omitempty"`
	EndorsementOrRestriction *string `json:"endorsementOrRestriction,omitempty"`
	ExchangeTicket           *string `json:"exchangeTicket,omitempty"`
	Fare                     *string `json:"fare,omitempty"`
	FareBasis                *string `json:"fareBasis,omitempty"`
	Fee                      *int32  `json:"fee,omitempty"`
	FlightNumber             *string `json:"flightNumber,omitempty"`
	Number                   *int32  `json:"number,omitempty"`
	OriginAirport            *string `json:"originAirport,omitempty"`
	PassengerClass           *string `json:"passengerClass,omitempty"`
	// Deprecated: Use passengerClass instead
	ServiceClass             *string `json:"serviceClass,omitempty"`
	StopoverCode             *string `json:"stopoverCode,omitempty"`
	Taxes                    *int32  `json:"taxes,omitempty"`
}

// NewAirlineFlightLeg constructs a new AirlineFlightLeg
func NewAirlineFlightLeg() *AirlineFlightLeg {
	return &AirlineFlightLeg{}
}
