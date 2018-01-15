// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/product"
)

func getDeviceFingerprintExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	var body product.DeviceFingerprintRequest

	response, err := client.Merchant("merchantId").Products().DeviceFingerprint(1, body, nil)

	fmt.Println(response, err)
}
