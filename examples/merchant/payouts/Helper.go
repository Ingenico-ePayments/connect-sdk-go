// This class was auto-generated from the API references found at
// https://epayments-api.developer-ingenico.com/

package examples

import (
	"github.com/Ingenico-ePayments/connect-sdk-go"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payout"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/refund"
)

func getClient() (*connectsdk.Client, error) {
	apiKeyID := "someKey"
	secretAPIKey := "someSecret"

	return connectsdk.CreateClient(apiKeyID, secretAPIKey, "Ingenico")
}

func newBool(value bool) *bool {
	return &value
}

func newInt32(value int32) *int32 {
	return &value
}

func newInt64(value int64) *int64 {
	return &value
}

func newString(value string) *string {
	return &value
}

func handleDeclinedPayment(paymentResult *payment.CreateResult) {
	// handle the result here
}

func handleDeclinedPayout(payoutResult *payout.Result) {
	// handle the result here
}

func handleDeclinedRefund(refundResult *refund.Result) {
	// handle the result here
}

func handleAPIErrors(errors []errors.APIError) {
	// handle the errors here
}
