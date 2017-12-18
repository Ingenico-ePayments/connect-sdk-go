// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// AirlinePassenger represents class AirlinePassenger
type AirlinePassenger struct {
	FirstName     *string `json:"firstName,omitempty"`
	Surname       *string `json:"surname,omitempty"`
	SurnamePrefix *string `json:"surnamePrefix,omitempty"`
	Title         *string `json:"title,omitempty"`
}

// NewAirlinePassenger constructs a new AirlinePassenger
func NewAirlinePassenger() *AirlinePassenger {
	return &AirlinePassenger{}
}
