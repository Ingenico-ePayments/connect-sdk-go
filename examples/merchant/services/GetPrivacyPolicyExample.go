// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/services"
)

func getPrivacyPolicyExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt32 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var query services.PrivacypolicyParams
	query.Locale = newString("en_US")
	query.PaymentProductID = newInt32(771)

	response, err := client.Merchant("merchantId").Services().Privacypolicy(query, nil)

	fmt.Println(response, err)
}
