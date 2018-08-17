// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import "fmt"

func getRefundsExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	response, err := client.Merchant("merchantId").Payments().Refunds("paymentId", nil)

	fmt.Println(response, err)
}
