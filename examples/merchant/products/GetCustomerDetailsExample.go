// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/product"
)

func getCustomerDetailsExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper function can be found in file Helper.go

	var values []definitions.KeyValuePair

	var value1 definitions.KeyValuePair
	value1.Key = newString("fiscalNumber")
	value1.Value = newString("01234567890")

	values = append(values, value1)

	var body product.GetCustomerDetailsRequest
	body.CountryCode = newString("SE")
	body.Values = &values

	response, err := client.Merchant("merchantId").Products().CustomerDetails(9000, body, nil)

	fmt.Println(response, err)
}
