// This file was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/refund"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
)

func refundPaymentExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(1)
	amountOfMoney.CurrencyCode = newString("EUR")

	var bankAccountIban definitions.BankAccountIban
	bankAccountIban.Iban = newString("NL53INGB0000000036")

	var bankRefundMethodSpecificInput refund.BankRefundMethodSpecificInput
	bankRefundMethodSpecificInput.BankAccountIban = &bankAccountIban

	var name payment.PersonalName
	name.Surname = newString("Coyote")

	var address payment.AddressPersonal
	address.CountryCode = newString("US")
	address.Name = &name

	var contactDetails definitions.ContactDetailsBase
	contactDetails.EmailAddress = newString("wile.e.coyote@acmelabs.com")
	contactDetails.EmailMessageType = newString("html")

	var customer refund.Customer
	customer.Address = &address
	customer.ContactDetails = &contactDetails

	var refundReferences refund.References
	refundReferences.MerchantReference = newString("AcmeOrder0001")

	var body refund.Request
	body.AmountOfMoney = &amountOfMoney
	body.BankRefundMethodSpecificInput = &bankRefundMethodSpecificInput
	body.Customer = &customer
	body.RefundDate = newString("20140306")
	body.RefundReferences = &refundReferences

	response, err := client.Merchant("merchantId").Payments().Refund("paymentId", body, nil)
	switch realError := err.(type) {
	case *sdkErrors.DeclinedRefundError:
		{
			handleDeclinedRefund(realError.RefundResult())

			break
		}
	case sdkErrors.APIError:
		{
			handleAPIErrors(realError.Errors())

			break
		}
	}

	fmt.Println(response, err)
}
