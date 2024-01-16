// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/token"
)

func approveSepaDirectDebitTokenExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newBool and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var body token.ApproveRequest
	body.MandateSignatureDate = newString("20150201")
	body.MandateSignaturePlace = newString("Monument Valley")
	body.MandateSigned = newBool(true)

	err := client.Merchant("merchantId").Tokens().Approvesepadirectdebit("tokenId", body, nil)

	fmt.Println(err)
}
