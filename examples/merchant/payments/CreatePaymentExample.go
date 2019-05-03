// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

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

	var externalCardholderAuthenticationData payment.ExternalCardholderAuthenticationData
	externalCardholderAuthenticationData.Cavv = newString("AgAAAAAABk4DWZ4C28yUQAAAAAA=")
	externalCardholderAuthenticationData.CavvAlgorithm = newString("1")
	externalCardholderAuthenticationData.Eci = newInt32(8)
	externalCardholderAuthenticationData.ThreeDSecureVersion = newString("v2")
	externalCardholderAuthenticationData.ThreeDServerTransactionID = newString("3DSTID1234")
	externalCardholderAuthenticationData.ValidationResult = newString("Y")
	externalCardholderAuthenticationData.Xid = newString("n3h2uOQPUgnmqhCkXNfxl8pOZJA=")

	var priorThreeDSecureData payment.ThreeDSecureData
	priorThreeDSecureData.AcsTransactionID = newString("empty")
	priorThreeDSecureData.Method = newString("challenged")
	priorThreeDSecureData.UtcTimestamp = newString("201901311530")

	var deviceRenderOptions payment.DeviceRenderOptions
	deviceRenderOptions.SdkInterface = newString("native")
	deviceRenderOptions.SdkUIType = newString("multi-select")

	var sdkData payment.SdkDataInput
	sdkData.DeviceInfo = newString("abc123")
	sdkData.DeviceRenderOptions = &deviceRenderOptions
	sdkData.SdkAppID = newString("xyz")
	sdkData.SdkEncryptedData = newString("abc123")
	sdkData.SdkEphemeralPublicKey = newString("123xyz")
	sdkData.SdkMaxTimeout = newString("30")
	sdkData.SdkReferenceNumber = newString("zaq123")
	sdkData.SdkTransactionID = newString("xsw321")

	var threeDSecure payment.ThreeDSecure
	threeDSecure.AuthenticationFlow = newString("browser")
	threeDSecure.ChallengeCanvasSize = newString("600x400")
	threeDSecure.ChallengeIndicator = newString("challenge-requested")
	threeDSecure.ExternalCardholderAuthenticationData = &externalCardholderAuthenticationData
	threeDSecure.PriorThreeDSecureData = &priorThreeDSecureData
	threeDSecure.SdkData = &sdkData
	threeDSecure.SkipAuthentication = newBool(false)

	var cardPaymentMethodSpecificInput payment.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = newInt32(1)
	cardPaymentMethodSpecificInput.ThreeDSecure = &threeDSecure

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
	companyInformation.VatNumber = newString("1234AB5678CD")

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
	personalInformation.Gender = newString("male")
	personalInformation.Name = &name

	var customer payment.Customer
	customer.BillingAddress = &billingAddress
	customer.CompanyInformation = &companyInformation
	customer.ContactDetails = &contactDetails
	customer.Locale = newString("en_US")
	customer.MerchantCustomerID = newString("1234")
	customer.PersonalInformation = &personalInformation

	var invoiceData payment.OrderInvoiceData
	invoiceData.InvoiceDate = newString("20140306191500")
	invoiceData.InvoiceNumber = newString("000000123")

	var references payment.OrderReferences
	references.Descriptor = newString("Fast and Furry-ous")
	references.InvoiceData = &invoiceData
	references.MerchantOrderID = newInt64(123456)
	references.MerchantReference = newString("AcmeOrder0001")

	var shippingName payment.PersonalName
	shippingName.FirstName = newString("Road")
	shippingName.Surname = newString("Runner")
	shippingName.Title = newString("Miss")

	var address payment.AddressPersonal
	address.AdditionalInfo = newString("Suite II")
	address.City = newString("Monument Valley")
	address.CountryCode = newString("US")
	address.HouseNumber = newString("1")
	address.Name = &shippingName
	address.State = newString("Utah")
	address.Street = newString("Desertroad")
	address.Zip = newString("84536")

	var shipping payment.Shipping
	shipping.Address = &address

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

	var shoppingCart payment.ShoppingCart
	shoppingCart.Items = &items

	var order payment.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer
	order.References = &references
	order.Shipping = &shipping
	order.ShoppingCart = &shoppingCart

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
