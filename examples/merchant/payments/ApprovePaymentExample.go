// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
)

func approvePaymentExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var directDebitPaymentMethodSpecificInput payment.ApprovePaymentNonSepaDirectDebitPaymentMethodSpecificInput
	directDebitPaymentMethodSpecificInput.DateCollect = newString("20150201")
	directDebitPaymentMethodSpecificInput.Token = newString("bfa8a7e4-4530-455a-858d-204ba2afb77e")

	var references payment.OrderReferencesApprovePayment
	references.MerchantReference = newString("AcmeOrder0001")

	var order payment.OrderApprovePayment
	order.References = &references

	var body payment.ApproveRequest
	body.Amount = newInt64(2980)
	body.DirectDebitPaymentMethodSpecificInput = &directDebitPaymentMethodSpecificInput
	body.Order = &order

	response, err := client.Merchant("merchantId").Payments().Approve("paymentId", body, nil)

	fmt.Println(response, err)
}
