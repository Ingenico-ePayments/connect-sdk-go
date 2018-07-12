// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/hostedcheckout"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

func createHostedCheckoutExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var hostedCheckoutSpecificInput hostedcheckout.SpecificInput
	hostedCheckoutSpecificInput.Locale = newString("en_GB")
	hostedCheckoutSpecificInput.Variant = newString("testVariant")

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(2345)
	amountOfMoney.CurrencyCode = newString("USD")

	var billingAddress definitions.Address
	billingAddress.CountryCode = newString("US")

	var customer payment.Customer
	customer.BillingAddress = &billingAddress
	customer.MerchantCustomerID = newString("1234")

	var order payment.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var body hostedcheckout.CreateRequest
	body.HostedCheckoutSpecificInput = &hostedCheckoutSpecificInput
	body.Order = &order

	response, err := client.Merchant("merchantId").Hostedcheckouts().Create(body, nil)

	fmt.Println(response, err)
}
