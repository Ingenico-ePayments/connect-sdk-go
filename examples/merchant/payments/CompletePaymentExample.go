// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

func completePaymentExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper function can be found in file Helper.go

	var card definitions.CardWithoutCvv
	card.CardNumber = newString("67030000000000003")
	card.CardholderName = newString("Wile E. Coyote")
	card.ExpiryDate = newString("1220")

	var cardPaymentMethodSpecificInput payment.CompletePaymentCardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card

	var body payment.CompletePaymentRequest
	body.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput

	response, err := client.Merchant("merchantId").Payments().Complete("paymentId", body, nil)

	fmt.Println(response, err)
}
