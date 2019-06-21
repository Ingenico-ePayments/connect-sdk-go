// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"
	"io"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
)

func getFileExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	err := client.Merchant("merchantId").Files().GetFile("fileId", nil, communicator.BodyHandlerFunc(func(headers []communication.Header, bodyReader io.Reader) error {
		// use the headers and body reader, and return any error that occurred
		return nil
	}))

	fmt.Println(err)
}
