// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/products"
)

func getPaymentProductExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newBool, newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var query products.GetParams
	query.CountryCode = newString("US")
	query.CurrencyCode = newString("USD")
	query.Locale = newString("en_US")
	query.Amount = newInt64(1000)
	query.IsRecurring = newBool(true)
	query.AddHide("fields")

	response, err := client.Merchant("merchantId").Products().Get(1, query, nil)

	fmt.Println(response, err)
}
