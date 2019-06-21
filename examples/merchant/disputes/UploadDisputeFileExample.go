// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"
	"os"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/disputes"
)

func uploadDisputeFileExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	var body disputes.UploadFileRequest
	fileContent, _ := os.Open("file.pdf")
	defer fileContent.Close()
	body.File, _ = communication.NewUploadableFile("file.pdf", fileContent, "application/pdf")

	response, err := client.Merchant("merchantId").Disputes().UploadFile("disputeId", body, nil)

	fmt.Println(response, err)
}
