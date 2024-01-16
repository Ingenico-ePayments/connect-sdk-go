// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import "fmt"

func getMandateExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	response, err := client.Merchant("merchantId").Mandates().Get("42268d8067df43e18a50a2ebf4bdb729", nil)

	fmt.Println(response, err)
}
