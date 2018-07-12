// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/hostedmandatemanagement"
)

func createHostedMandateManagementExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper function can be found in file Helper.go

	var createMandateInfo hostedmandatemanagement.HostedMandateInfo
	createMandateInfo.CustomerReference = newString("idonthaveareference")
	createMandateInfo.RecurrenceType = newString("RECURRING")
	createMandateInfo.SignatureType = newString("UNSIGNED")

	var hostedMandateManagementSpecificInput hostedmandatemanagement.SpecificInput
	hostedMandateManagementSpecificInput.Locale = newString("fr_FR")
	hostedMandateManagementSpecificInput.ReturnURL = newString("http://www.example.com")
	hostedMandateManagementSpecificInput.Variant = newString("101")

	var body hostedmandatemanagement.CreateRequest
	body.CreateMandateInfo = &createMandateInfo
	body.HostedMandateManagementSpecificInput = &hostedMandateManagementSpecificInput

	response, err := client.Merchant("merchantId").Hostedmandatemanagements().Create(body, nil)

	fmt.Println(response, err)
}
