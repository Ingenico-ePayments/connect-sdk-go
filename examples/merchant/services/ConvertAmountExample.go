// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/services"
)

func convertAmountExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var query services.ConvertAmountParams
	query.Source = newString("EUR")
	query.Target = newString("USD")
	query.Amount = newInt64(100)

	response, err := client.Merchant("merchantId").Services().ConvertAmount(query, nil)

	fmt.Println(response, err)
}
