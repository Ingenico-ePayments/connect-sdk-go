// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/refunds"
)

func findRefundsExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt32, newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var query refunds.FindParams
	query.HostedCheckoutID = newString("15c09dac-bf44-486a-af6b-edfd8680a166")
	query.MerchantReference = newString("AcmeOrder0001")
	query.MerchantOrderID = newInt64(123456)
	query.Offset = newInt32(0)
	query.Limit = newInt32(10)

	response, err := client.Merchant("merchantId").Refunds().Find(query, nil)

	fmt.Println(response, err)
}
