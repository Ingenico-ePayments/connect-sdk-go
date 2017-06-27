// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/riskassessments"
)

func riskAssessmentBankAccountExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var bankAccountBban definitions.BankAccountBban
	bankAccountBban.AccountNumber = newString("0532013000")
	bankAccountBban.BankCode = newString("37040044")
	bankAccountBban.CountryCode = newString("DE")

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(100)
	amountOfMoney.CurrencyCode = newString("EUR")

	var billingAddress definitions.Address
	billingAddress.CountryCode = newString("US")

	var customer riskassessments.CustomerRiskAssessment
	customer.BillingAddress = &billingAddress
	customer.Locale = newString("en_US")

	var order riskassessments.OrderRiskAssessment
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var body riskassessments.RiskAssessmentBankAccount
	body.BankAccountBban = &bankAccountBban
	body.Order = &order

	response, err := client.Merchant("merchantId").Riskassessments().Bankaccounts(body, nil)

	fmt.Println(response, err)
}
