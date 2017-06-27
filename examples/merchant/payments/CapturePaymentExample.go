// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

func capturePaymentExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newInt64 to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper function can be found in file Helper.go

	var body payment.CapturePaymentRequest
	body.Amount = newInt64(2980)

	response, err := client.Merchant("merchantId").Payments().Capture("paymentId", body, nil)

	fmt.Println(response, err)
}
