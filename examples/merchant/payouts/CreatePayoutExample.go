// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/s2sapi/v1/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payout"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
)

func createPayoutExample() {
	client, clientErr := getClient()
	if clientErr != nil {
		panic(clientErr)
	}
	defer client.Close()

	// Assigning literals to pointer variables directly is not supported.
	// The below code uses helper functions newInt64 and newString to overcome this issue.
	// http://stackoverflow.com/a/30716481 lists a few more alternatives.
	// The code for these helper functions can be found in file Helper.go

	var bankAccountIban definitions.BankAccountIban
	bankAccountIban.AccountHolderName = newString("Wile E. Coyote")
	bankAccountIban.Iban = newString("IT60X0542811101000000123456")

	var bankTransferPayoutMethodSpecificInput payout.BankTransferPayoutMethodSpecificInput
	bankTransferPayoutMethodSpecificInput.BankAccountIban = &bankAccountIban
	bankTransferPayoutMethodSpecificInput.PayoutDate = newString("20150102")
	bankTransferPayoutMethodSpecificInput.PayoutText = newString("Payout Acme")
	bankTransferPayoutMethodSpecificInput.SwiftCode = newString("swift")

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(2345)
	amountOfMoney.CurrencyCode = newString("EUR")

	var address definitions.Address
	address.City = newString("Burbank")
	address.CountryCode = newString("US")
	address.HouseNumber = newString("411")
	address.State = newString("California")
	address.Street = newString("N Hollywood Way")
	address.Zip = newString("91505")

	var companyInformation definitions.CompanyInformation
	companyInformation.Name = newString("Acme Labs")

	var contactDetails definitions.ContactDetailsBase
	contactDetails.EmailAddress = newString("wile.e.coyote@acmelabs.com")

	var name payment.PersonalName
	name.FirstName = newString("Wile")
	name.Surname = newString("Coyote")
	name.SurnamePrefix = newString("E.")
	name.Title = newString("Mr.")

	var customer payout.Customer
	customer.Address = &address
	customer.CompanyInformation = &companyInformation
	customer.ContactDetails = &contactDetails
	customer.Name = &name

	var references payout.References
	references.MerchantReference = newString("AcmeOrder001")

	var payoutDetails payout.Details
	payoutDetails.AmountOfMoney = &amountOfMoney
	payoutDetails.Customer = &customer
	payoutDetails.References = &references

	var body payout.CreateRequest
	body.BankTransferPayoutMethodSpecificInput = &bankTransferPayoutMethodSpecificInput
	body.PayoutDetails = &payoutDetails

	response, err := client.Merchant("merchantId").Payouts().Create(body, nil)
	switch realError := err.(type) {
	case *sdkErrors.DeclinedPayoutError:
		{
			handleDeclinedPayout(realError.PayoutResult())

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
