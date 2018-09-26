// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/product"
)

func getDeviceFingerprintForGroupsExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	var body product.DeviceFingerprintRequest

	response, err := client.Merchant("merchantId").Productgroups().DeviceFingerprint("cards", body, nil)

	fmt.Println(response, err)
}
