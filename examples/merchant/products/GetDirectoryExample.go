// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/products"
)

func getDirectoryExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper functions can be found in file Helper.go

	var query products.DirectoryParams
	query.CurrencyCode = newString("EUR")
	query.CountryCode = newString("NL")

	response, err := client.Merchant("merchantId").Products().Directory(809, query, nil)

	fmt.Println(response, err)
}
