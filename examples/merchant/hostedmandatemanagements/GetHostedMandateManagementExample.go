// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import "fmt"

func getHostedMandateManagementExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	response, err := client.Merchant("merchantId").Hostedmandatemanagements().Get("hostedMandateManagementId", nil)

	fmt.Println(response, err)
}
