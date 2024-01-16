// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import "fmt"

func deleteHostedCheckoutExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	err := client.Merchant("merchantId").Hostedcheckouts().Delete("hostedCheckoutId", nil)

	fmt.Println(err)
}
