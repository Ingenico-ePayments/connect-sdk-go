// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/riskassessments"
)

func riskAssessmentCardsExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newBool, newInt32, newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var card definitions.Card
	card.CardNumber = newString("4567350000427977")
	card.Cvv = newString("123")
	card.ExpiryDate = newString("0820")

	var flightLegs []definitions.AirlineFlightLeg

	var flightLeg1 definitions.AirlineFlightLeg
	flightLeg1.AirlineClass = newString("1")
	flightLeg1.ArrivalAirport = newString("AMS")
	flightLeg1.CarrierCode = newString("KL")
	flightLeg1.Date = newString("20150102")
	flightLeg1.DepartureTime = newString("17:59")
	flightLeg1.Fare = newString("fare")
	flightLeg1.FareBasis = newString("INTERNET")
	flightLeg1.FlightNumber = newString("791")
	flightLeg1.Number = newInt32(1)
	flightLeg1.OriginAirport = newString("BCN")
	flightLeg1.StopoverCode = newString("non-permitted")

	flightLegs = append(flightLegs, flightLeg1)

	var flightLeg2 definitions.AirlineFlightLeg
	flightLeg2.AirlineClass = newString("1")
	flightLeg2.ArrivalAirport = newString("BCN")
	flightLeg2.CarrierCode = newString("KL")
	flightLeg2.Date = newString("20150102")
	flightLeg2.DepartureTime = newString("23:59")
	flightLeg2.Fare = newString("fare")
	flightLeg2.FareBasis = newString("INTERNET")
	flightLeg2.FlightNumber = newString("792")
	flightLeg2.Number = newInt32(2)
	flightLeg2.OriginAirport = newString("AMS")
	flightLeg2.StopoverCode = newString("non-permitted")

	flightLegs = append(flightLegs, flightLeg2)

	var airlineData definitions.AirlineData
	airlineData.AgentNumericCode = newString("123321")
	airlineData.Code = newString("123")
	airlineData.FlightDate = newString("20150102")
	airlineData.FlightLegs = &flightLegs
	airlineData.InvoiceNumber = newString("123456")
	airlineData.IsETicket = newBool(true)
	airlineData.IsRestrictedTicket = newBool(true)
	airlineData.IsThirdParty = newBool(true)
	airlineData.IssueDate = newString("20150101")
	airlineData.MerchantCustomerID = newString("14")
	airlineData.Name = newString("Air France KLM")
	airlineData.PassengerName = newString("WECOYOTE")
	airlineData.PlaceOfIssue = newString("Utah")
	airlineData.PNR = newString("4JTGKT")
	airlineData.PointOfSale = newString("IATA point of sale name")
	airlineData.PosCityCode = newString("AMS")
	airlineData.TicketDeliveryMethod = newString("e-ticket")
	airlineData.TicketNumber = newString("KLM20050000")

	var additionalInput definitions.AdditionalOrderInputAirlineData
	additionalInput.AirlineData = &airlineData

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(100)
	amountOfMoney.CurrencyCode = newString("EUR")

	var billingAddress definitions.Address
	billingAddress.CountryCode = newString("US")

	var customer riskassessments.CustomerRiskAssessment
	customer.AccountType = newString("existing")
	customer.BillingAddress = &billingAddress
	customer.Locale = newString("en_US")

	var order riskassessments.OrderRiskAssessment
	order.AdditionalInput = &additionalInput
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var body riskassessments.RiskAssessmentCard
	body.Card = &card
	body.Order = &order

	response, err := client.Merchant("merchantId").Riskassessments().Cards(body, nil)

	fmt.Println(response, err)
}
