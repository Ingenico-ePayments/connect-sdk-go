// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import "fmt"

func unblockMandateExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	response, err := client.Merchant("merchantId").Mandates().Unblock("42268d8067df43e18a50a2ebf4bdb729", nil)

	fmt.Println(response, err)
}
