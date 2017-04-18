// This file was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/services"
)

func convertBankAccountExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper function newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for this helper function can be found in file Helper.go

	var bankAccountBban definitions.BankAccountBban
	bankAccountBban.AccountNumber = newString("0532013000")
	bankAccountBban.BankCode = newString("37040044")
	bankAccountBban.CountryCode = newString("DE")

	var body services.BankDetailsRequest
	body.BankAccountBban = &bankAccountBban

	response, err := client.Merchant("merchantId").Services().Bankaccount(body, nil)

	fmt.Println(response, err)
}
