// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
)

func createPaymentExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newBool, newInt32, newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var card definitions.Card
	card.CardNumber = newString("4567350000427977")
	card.CardholderName = newString("Wile E. Coyote")
	card.Cvv = newString("123")
	card.ExpiryDate = newString("1220")

	var cardPaymentMethodSpecificInput payment.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = newInt32(1)
	cardPaymentMethodSpecificInput.SkipAuthentication = newBool(false)

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(2980)
	amountOfMoney.CurrencyCode = newString("EUR")

	var billingAddress definitions.Address
	billingAddress.AdditionalInfo = newString("b")
	billingAddress.City = newString("Monument Valley")
	billingAddress.CountryCode = newString("US")
	billingAddress.HouseNumber = newString("13")
	billingAddress.State = newString("Utah")
	billingAddress.Street = newString("Desertroad")
	billingAddress.Zip = newString("84536")

	var companyInformation definitions.CompanyInformation
	companyInformation.Name = newString("Acme Labs")

	var contactDetails payment.ContactDetails
	contactDetails.EmailAddress = newString("wile.e.coyote@acmelabs.com")
	contactDetails.EmailMessageType = newString("html")
	contactDetails.FaxNumber = newString("+1234567891")
	contactDetails.PhoneNumber = newString("+1234567890")

	var name payment.PersonalName
	name.FirstName = newString("Wile")
	name.Surname = newString("Coyote")
	name.SurnamePrefix = newString("E.")
	name.Title = newString("Mr.")

	var personalInformation payment.PersonalInformation
	personalInformation.DateOfBirth = newString("19490917")
	personalInformation.Gender = newString("M")
	personalInformation.Name = &name

	var shippingName payment.PersonalName
	shippingName.FirstName = newString("Road")
	shippingName.Surname = newString("Runner")
	shippingName.Title = newString("Miss")

	var shippingAddress payment.AddressPersonal
	shippingAddress.AdditionalInfo = newString("Suite II")
	shippingAddress.City = newString("Monument Valley")
	shippingAddress.CountryCode = newString("US")
	shippingAddress.HouseNumber = newString("1")
	shippingAddress.Name = &shippingName
	shippingAddress.State = newString("Utah")
	shippingAddress.Street = newString("Desertroad")
	shippingAddress.Zip = newString("84536")

	var customer payment.Customer
	customer.BillingAddress = &billingAddress
	customer.CompanyInformation = &companyInformation
	customer.ContactDetails = &contactDetails
	customer.Locale = newString("en_US")
	customer.MerchantCustomerID = newString("1234")
	customer.PersonalInformation = &personalInformation
	customer.ShippingAddress = &shippingAddress
	customer.VatNumber = newString("1234AB5678CD")

	var items []payment.LineItem

	var item1AmountOfMoney definitions.AmountOfMoney
	item1AmountOfMoney.Amount = newInt64(2500)
	item1AmountOfMoney.CurrencyCode = newString("EUR")

	var item1InvoiceData payment.LineItemInvoiceData
	item1InvoiceData.Description = newString("ACME Super Outfit")
	item1InvoiceData.NrOfItems = newString("1")
	item1InvoiceData.PricePerItem = newInt64(2500)

	var item1 payment.LineItem
	item1.AmountOfMoney = &item1AmountOfMoney
	item1.InvoiceData = &item1InvoiceData

	items = append(items, item1)

	var item2AmountOfMoney definitions.AmountOfMoney
	item2AmountOfMoney.Amount = newInt64(480)
	item2AmountOfMoney.CurrencyCode = newString("EUR")

	var item2InvoiceData payment.LineItemInvoiceData
	item2InvoiceData.Description = newString("Aspirin")
	item2InvoiceData.NrOfItems = newString("12")
	item2InvoiceData.PricePerItem = newInt64(40)

	var item2 payment.LineItem
	item2.AmountOfMoney = &item2AmountOfMoney
	item2.InvoiceData = &item2InvoiceData

	items = append(items, item2)

	var invoiceData payment.OrderInvoiceData
	invoiceData.InvoiceDate = newString("20140306191500")
	invoiceData.InvoiceNumber = newString("000000123")

	var references payment.OrderReferences
	references.Descriptor = newString("Fast and Furry-ous")
	references.InvoiceData = &invoiceData
	references.MerchantOrderID = newInt64(123456)
	references.MerchantReference = newString("AcmeOrder0001")

	var order payment.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer
	order.Items = &items
	order.References = &references

	var body payment.CreateRequest
	body.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	body.Order = &order

	response, err := client.Merchant("merchantId").Payments().Create(body, nil)
	switch realError := err.(type) {
	case *sdkErrors.DeclinedPaymentError:
		{
			handleDeclinedPayment(realError.PaymentResult())

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
