// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package definitions

// AirlineData represents class AirlineData
type AirlineData struct {
	AgentNumericCode     *string             `json:"agentNumericCode,omitempty"`
	Code                 *string             `json:"code,omitempty"`
	FlightDate           *string             `json:"flightDate,omitempty"`
	FlightLegs           *[]AirlineFlightLeg `json:"flightLegs,omitempty"`
	InvoiceNumber        *string             `json:"invoiceNumber,omitempty"`
	IsETicket            *bool               `json:"isETicket,omitempty"`
	// Deprecated: Use Order.customer.accountType instead
	IsRegisteredCustomer *bool               `json:"isRegisteredCustomer,omitempty"`
	IsRestrictedTicket   *bool               `json:"isRestrictedTicket,omitempty"`
	IsThirdParty         *bool               `json:"isThirdParty,omitempty"`
	IssueDate            *string             `json:"issueDate,omitempty"`
	MerchantCustomerID   *string             `json:"merchantCustomerId,omitempty"`
	Name                 *string             `json:"name,omitempty"`
	NumberInParty        *int32              `json:"numberInParty,omitempty"`
	PassengerName        *string             `json:"passengerName,omitempty"`
	Passengers           *[]AirlinePassenger `json:"passengers,omitempty"`
	PlaceOfIssue         *string             `json:"placeOfIssue,omitempty"`
	PNR                  *string             `json:"pnr,omitempty"`
	PointOfSale          *string             `json:"pointOfSale,omitempty"`
	PosCityCode          *string             `json:"posCityCode,omitempty"`
	TicketDeliveryMethod *string             `json:"ticketDeliveryMethod,omitempty"`
	TicketNumber         *string             `json:"ticketNumber,omitempty"`
	TotalFare            *int32              `json:"totalFare,omitempty"`
	TotalFee             *int32              `json:"totalFee,omitempty"`
	TotalTaxes           *int32              `json:"totalTaxes,omitempty"`
	TravelAgencyName     *string             `json:"travelAgencyName,omitempty"`
}

// NewAirlineData constructs a new AirlineData
func NewAirlineData() *AirlineData {
	return &AirlineData{}
}
