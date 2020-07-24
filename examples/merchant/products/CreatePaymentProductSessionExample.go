// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/product"
)

func createPaymentProductSessionExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper function can be found in file Helper.go

	var paymentProductSession302SpecificInput product.MobilePaymentProductSession302SpecificInput
	paymentProductSession302SpecificInput.DisplayName = newString("Ingenico")
	paymentProductSession302SpecificInput.DomainName = newString("pay1.secured-by-ingenico.com")
	paymentProductSession302SpecificInput.ValidationURL = newString("https://apple-pay-gateway-cert.apple.com/paymentservices/startSession")

	var body product.CreatePaymentProductSessionRequest
	body.PaymentProductSession302SpecificInput = &paymentProductSession302SpecificInput

	response, err := client.Merchant("merchantId").Products().Sessions(302, body, nil)

	fmt.Println(response, err)
}
