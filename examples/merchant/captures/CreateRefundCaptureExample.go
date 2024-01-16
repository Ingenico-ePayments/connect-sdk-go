// This file was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import (
	"fmt"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/refund"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
)

func createRefundCaptureExample() {
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
	amountOfMoney.Amount = newInt64(500)
	amountOfMoney.CurrencyCode = newString("EUR")

	var refundReferences refund.References
	refundReferences.MerchantReference = newString("AcmeOrder0001")

	var body refund.Request
	body.AmountOfMoney = &amountOfMoney
	body.RefundReferences = &refundReferences

	response, err := client.Merchant("merchantId").Captures().Refund("captureId", body, nil)
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
