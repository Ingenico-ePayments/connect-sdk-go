// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/installments"
)

func getInstallmentsInfoExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt32, newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(123)
	amountOfMoney.CurrencyCode = newString("EUR")

	var body installments.GetInstallmentRequest
	body.AmountOfMoney = &amountOfMoney
	body.Bin = newString("123455")
	body.CountryCode = newString("NL")
	body.PaymentProductID = newInt32(123)

	response, err := client.Merchant("merchantId").Installments().GetInstallmentsInfo(body, nil)

	fmt.Println(response, err)
}
