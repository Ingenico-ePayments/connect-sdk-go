// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/token"
)

func createTokenExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt32 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var billingAddress definitions.Address
	billingAddress.AdditionalInfo = newString("Suite II")
	billingAddress.City = newString("Monument Valley")
	billingAddress.CountryCode = newString("US")
	billingAddress.HouseNumber = newString("1")
	billingAddress.State = newString("Utah")
	billingAddress.Street = newString("Desertroad")
	billingAddress.Zip = newString("84536")

	var companyInformation definitions.CompanyInformation
	companyInformation.Name = newString("Acme Labs")

	var name token.PersonalNameToken
	name.FirstName = newString("Wile")
	name.Surname = newString("Coyote")
	name.SurnamePrefix = newString("E.")

	var personalInformation token.PersonalInformationToken
	personalInformation.Name = &name

	var customer token.CustomerToken
	customer.BillingAddress = &billingAddress
	customer.CompanyInformation = &companyInformation
	customer.MerchantCustomerID = newString("1234")
	customer.PersonalInformation = &personalInformation

	var bankAccountBban definitions.BankAccountBban
	bankAccountBban.AccountNumber = newString("000000123456")
	bankAccountBban.BankCode = newString("05428")
	bankAccountBban.BranchCode = newString("11101")
	bankAccountBban.CheckDigit = newString("X")
	bankAccountBban.CountryCode = newString("IT")

	var paymentProduct705SpecificData token.NonSepaDirectDebitPaymentProduct705SpecificData
	paymentProduct705SpecificData.AuthorisationID = newString("123456")
	paymentProduct705SpecificData.BankAccountBban = &bankAccountBban

	var mandate token.MandateNonSepaDirectDebit
	mandate.PaymentProduct705SpecificData = &paymentProduct705SpecificData

	var nonSepaDirectDebit token.NonSepaDirectDebit
	nonSepaDirectDebit.Customer = &customer
	nonSepaDirectDebit.Mandate = &mandate

	var body token.CreateRequest
	body.NonSepaDirectDebit = &nonSepaDirectDebit
	body.PaymentProductID = newInt32(705)

	response, err := client.Merchant("merchantId").Tokens().Create(body, nil)

	fmt.Println(response, err)
}
