// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import "fmt"

func cancelApprovalPayoutExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	err := client.Merchant("merchantId").Payouts().Cancelapproval("payoutId", nil)

	fmt.Println(err)
}
