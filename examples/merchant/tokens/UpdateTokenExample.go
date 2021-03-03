// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/token"
)

func updateTokenExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt32 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var billingAddress definitions.Address
	billingAddress.AdditionalInfo = newString("b")
	billingAddress.City = newString("Monument Valley")
	billingAddress.CountryCode = newString("US")
	billingAddress.HouseNumber = newString("13")
	billingAddress.State = newString("Utah")
	billingAddress.Street = newString("Desertroad")
	billingAddress.Zip = newString("84536")

	var companyInformation definitions.CompanyInformation
	companyInformation.Name = newString("Acme Labs")

	var name token.PersonalNameToken
	name.FirstName = newString("Wile")
	name.Surname = newString("Coyote")
	name.SurnamePrefix = newString("E.")

	var personalInformation token.PersonalInformationToken
	personalInformation.Name = &name

	var customer token.CustomerToken
	customer.BillingAddress = &billingAddress
	customer.CompanyInformation = &companyInformation
	customer.MerchantCustomerID = newString("1234")
	customer.PersonalInformation = &personalInformation

	var cardWithoutCvv definitions.CardWithoutCvv
	cardWithoutCvv.CardNumber = newString("4567350000427977")
	cardWithoutCvv.CardholderName = newString("Wile E. Coyote")
	cardWithoutCvv.ExpiryDate = newString("1299")
	cardWithoutCvv.IssueNumber = newString("12")

	var data token.CardData
	data.CardWithoutCvv = &cardWithoutCvv

	var card token.Card
	card.Customer = &customer
	card.Data = &data

	var body token.UpdateRequest
	body.Card = &card
	body.PaymentProductID = newInt32(1)

	err := client.Merchant("merchantId").Tokens().Update("tokenId", body, nil)

	fmt.Println(err)
}
