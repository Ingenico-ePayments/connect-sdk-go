// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/dispute"
)

func createPaymentDisputeExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(1234)
	amountOfMoney.CurrencyCode = newString("USD")

	var body dispute.CreateRequest
	body.AmountOfMoney = &amountOfMoney
	body.ContactPerson = newString("Wile Coyote")
	body.EmailAddress = newString("wile.e.coyote@acmelabs.com")
	body.ReplyTo = newString("r.runner@acmelabs.com")
	body.RequestMessage = newString("This is the message from the merchant to GlobalCollect. It is a a freeform text field.")

	response, err := client.Merchant("merchantId").Payments().Dispute("paymentId", body, nil)

	fmt.Println(response, err)
}
