// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import "fmt"

func cancelRefundExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	err := client.Merchant("merchantId").Refunds().Cancel("refundId", nil)

	fmt.Println(err)
}
