// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import "fmt"

func getPaymentDisputesExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	response, err := client.Merchant("merchantId").Payments().Disputes("paymentId", nil)

	fmt.Println(response, err)
}
