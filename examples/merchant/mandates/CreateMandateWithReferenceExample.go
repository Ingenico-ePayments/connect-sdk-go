// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/mandates"
)

func createMandateWithReferenceExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper function can be found in file Helper.go

	var bankAccountIban definitions.BankAccountIban
	bankAccountIban.Iban = newString("DE46720200700359736690")

	var contactDetails mandates.MandateContactDetails
	contactDetails.EmailAddress = newString("wile.e.coyote@acmelabs.com")

	var mandateAddress mandates.MandateAddress
	mandateAddress.City = newString("Monumentenvallei")
	mandateAddress.CountryCode = newString("NL")
	mandateAddress.Street = newString("Woestijnweg")
	mandateAddress.Zip = newString("1337XD")

	var name mandates.MandatePersonalName
	name.FirstName = newString("Wile")
	name.Surname = newString("Coyote")

	var personalInformation mandates.MandatePersonalInformation
	personalInformation.Name = &name
	personalInformation.Title = newString("Miss")

	var customer mandates.MandateCustomer
	customer.BankAccountIban = &bankAccountIban
	customer.CompanyName = newString("Acme labs")
	customer.ContactDetails = &contactDetails
	customer.MandateAddress = &mandateAddress
	customer.PersonalInformation = &personalInformation

	var body mandates.CreateMandateRequest
	body.Customer = &customer
	body.CustomerReference = newString("idonthaveareference")
	body.Language = newString("nl")
	body.RecurrenceType = newString("UNIQUE")
	body.SignatureType = newString("UNSIGNED")

	response, err := client.Merchant("merchantId").Mandates().CreateWithMandateReference("42268d8067df43e18a50a2ebf4bdb729", body, nil)

	fmt.Println(response, err)
}
