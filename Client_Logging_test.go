package connectsdk

import (
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/logging"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/services"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/tokens"
)

var testConnectionJSON = `{
	"result": "OK"
}`

var convertAmountJSON = `{
	"convertedAmount": 4547504
}`

var createPaymentUnicodeJSON = `{
    "creationOutput": {
        "additionalReference": "00000012341000059598",
        "externalReference": "000000123410000595980000100001"
    },
    "payment": {
        "id": "000000123410000595980000100001",
        "paymentOutput": {
            "amountOfMoney": {
                "amount": 2345,
                "currencyCode": "CAD"
            },
            "references": {
                "paymentReference": "0"
            },
            "paymentMethod": "redirect",
            "redirectPaymentMethodSpecificOutput":{
               "paymentProductId":840,
               "paymentProduct840SpecificOutput":{
                  "customerAccount":{
                     "firstName":"Theresa",
                     "surname":"Schröder"
                  },
                  "customerAddress":{
                     "city":"sittensen",
                     "countryCode":"DE",
                     "street":"Westerberg 25",
                     "zip":"27419"
                  }
               }
            }
        },
        "status": "PENDING_APPROVAL",
        "statusOutput": {
            "isCancellable": true,
            "statusCategory": "PENDING_MERCHANT",
            "statusCode": 600,
            "statusCodeChangeDateTime": "20160310094054",
            "isAuthorized": true
        }
    }
}
`

var createPaymentJSON = `{
	"creationOutput": {
		"additionalReference": "00000012341000059598",
		"externalReference": "000000123410000595980000100001"
	},
	"payment": {
		"id": "000000123410000595980000100001",
		"paymentOutput": {
			"amountOfMoney": {
				"amount": 2345,
				"currencyCode": "CAD"
			},
			"references": {
				"paymentReference": "0"
			},
			"paymentMethod": "card",
			"cardPaymentMethodSpecificOutput": {
				"paymentProductId": 1,
				"authorisationCode": "OK1131",
				"card": {
					"cardNumber": "************3456",
					"expiryDate": "1220"
				},
				"fraudResults": {
					"fraudServiceResult": "error",
					"avsResult": "X",
					"cvvResult": "M"
				}
			}
		},
		"status": "PENDING_APPROVAL",
		"statusOutput": {
			"isCancellable": true,
			"statusCategory": "PENDING_MERCHANT",
			"statusCode": 600,
			"statusCodeChangeDateTime": "20160310094054",
			"isAuthorized": true
		}
	}
}`

var createPaymentFailedInvalidCardNumberJSON = `{
	"errorId": "0953f236-9e54-4f23-9556-d66bc757dda8",
	"errors": [{
		"code": "21000020",
		"requestId": "24146",
		"message": "VALUE **************** OF FIELD CREDITCARDNUMBER DID NOT PASS THE LUHNCHECK",
		"httpStatusCode": 400
	}]
}`

var createPaymentFailedRejectedJSON = `{
	"errorId": "833dfd83-52ae-419c-b871-9df1278da93e",
	"errors": [{
		"code": "430330",
		"message": "Not authorised",
		"httpStatusCode": 402
	}],
	"paymentResult": {
		"creationOutput": {
			"additionalReference": "00000200001000059614",
			"externalReference": "000002000010000596140000100001"
		},
		"payment": {
			"id": "000002000010000596140000100001",
			"paymentOutput": {
				"amountOfMoney": {
					"amount": 2345,
					"currencyCode": "CAD"
				},
				"references": {
					"paymentReference": "0"
				},
				"paymentMethod": "card",
				"cardPaymentMethodSpecificOutput": {
					"paymentProductId": 1
				}
			},
			"status": "REJECTED",
			"statusOutput": {
				"errors": [{
					"code": "430330",
					"requestId": "24162",
					"message": "Not authorised",
					"httpStatusCode": 402
				}],
				"isCancellable": false,
				"statusCategory": "UNSUCCESSFUL",
				"statusCode": 100,
				"statusCodeChangeDateTime": "20160310121151",
				"isAuthorized": false
			}
		}
	}
}`

var unknownServerErrorJSON = `{
	"errorId": "fbff1179-7ba4-4894-9021-d8a0011d23a7",
	"errors": [{
		"code": "9999",
		"message": "UNKNOWN_SERVER_ERROR",
		"httpStatusCode": 500
	}]
}`

var notFoundErrorHTML = `Not Found`

func TestLoggingTestConnection(t *testing.T) {
	logPrefix := "TestLoggingTestConnection"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/services/testconnection",
		createRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	response, err := client.Merchant("1234").Services().Testconnection(nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if response.Result != nil && *response.Result != "OK" {
		t.Fatalf("%v: responseResult %v", logPrefix, response.Result)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "GET" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/services/testconnection" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if secondEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response.StatusCode() != http.StatusOK {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, secondEntry.response.StatusCode())
	}
	if secondEntry.response.ContentType() != "application/json" {
		t.Fatalf("%v: secondEntryResponseContentType %v", logPrefix, secondEntry.response.ContentType())
	}
	if secondEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, secondEntry.response.Headers())
	}
	if secondEntry.response.Body() == "" {
		t.Fatalf("%v: secondEntryResponseBody %v", logPrefix, secondEntry.response.Body())
	}
	foundDate, foundDummy := false, false
	for k := range secondEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if secondEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}

func TestLoggingConvertAmount(t *testing.T) {
	logPrefix := "TestLoggingConvertAmount"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/services/convert/amount",
		createRecordRequest(http.StatusOK, convertAmountJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	var query services.ConvertAmountParams
	amount := int64(1000)
	query.Amount = &amount
	source := string("EUR")
	query.Source = &source
	target := string("USD")
	query.Target = &target

	response, err := client.Merchant("1234").Services().ConvertAmount(query, nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if response.ConvertedAmount == nil {
		t.Fatalf("%v: responseResult %v", logPrefix, response.ConvertedAmount)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "GET" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/services/convert/amount" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	if firstEntry.request.Body() != "" {
		t.Fatalf("%v: firstEntryRequestBody %v", logPrefix, firstEntry.request.Body())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if secondEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response.StatusCode() != http.StatusOK {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, secondEntry.response.StatusCode())
	}
	if secondEntry.response.ContentType() != "application/json" {
		t.Fatalf("%v: secondEntryResponseContentType %v", logPrefix, secondEntry.response.ContentType())
	}
	if secondEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, secondEntry.response.Headers())
	}
	if secondEntry.response.Body() == "" {
		t.Fatalf("%v: secondEntryResponseBody %v", logPrefix, secondEntry.response.Body())
	}
	foundDate, foundDummy := false, false
	for k := range secondEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if secondEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}

func TestDeleteToken(t *testing.T) {
	logPrefix := "TestDeleteToken"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/tokens/5678",
		createRecordRequest(http.StatusNoContent, "", responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	var query tokens.DeleteParams

	err = client.Merchant("1234").Tokens().Delete("5678", query, nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "DELETE" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/tokens/5678" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	if firstEntry.request.Body() != "" {
		t.Fatalf("%v: firstEntryRequestBody %v", logPrefix, firstEntry.request.Body())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if secondEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response.StatusCode() != http.StatusNoContent {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, secondEntry.response.StatusCode())
	}
	if secondEntry.response.ContentType() != "application/json" {
		t.Fatalf("%v: secondEntryResponseContentType %v", logPrefix, secondEntry.response.ContentType())
	}
	if secondEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, secondEntry.response.Headers())
	}
	if secondEntry.response.Body() != "" {
		t.Fatalf("%v: secondEntryResponseBody %v", logPrefix, secondEntry.response.Body())
	}
	foundDate, foundDummy := false, false
	for k := range secondEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if secondEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}

func TestLoggingCreatePaymentUnicode(t *testing.T) {
	logPrefix := "TestLoggingCreatePayment"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
		"Location":     "http://localhost/v1/1234/payments/000000123410000595980000100001",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/payments",
		createRecordRequest(http.StatusCreated, createPaymentUnicodeJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	var card definitions.Card
	card.CardNumber = newString("1234567890123456")
	card.Cvv = newString("123")
	card.ExpiryDate = newString("1220")

	var cardPaymentMethodSpecificInput payment.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = newInt32(1)

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(2345)
	amountOfMoney.CurrencyCode = newString("CAD")

	var billingAddress definitions.Address
	billingAddress.CountryCode = newString("CA")

	var customer payment.Customer
	customer.BillingAddress = &billingAddress

	var order payment.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var reqBody payment.CreateRequest
	reqBody.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	reqBody.Order = &order

	response, err := client.Merchant("1234").Payments().Create(reqBody, nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	if response.Payment == nil {
		t.Fatalf("%v: responsePayment nil", logPrefix)
	}
	if response.Payment.ID == nil {
		t.Fatalf("%v: responsePaymentID nil", logPrefix)
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "POST" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/payments" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if !strings.Contains(secondEntry.response.Body(), "Schröder") {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response.StatusCode() != http.StatusCreated {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, secondEntry.response.StatusCode())
	}
	if secondEntry.response.ContentType() != "application/json" {
		t.Fatalf("%v: secondEntryResponseContentType %v", logPrefix, secondEntry.response.ContentType())
	}
	if secondEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, secondEntry.response.Headers())
	}
	if secondEntry.response.Body() == "" {
		t.Fatalf("%v: secondEntryResponseBody %v", logPrefix, secondEntry.response.Body())
	}

	foundDate, foundDummy := false, false
	for k := range secondEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if secondEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}
func TestLoggingCreatePayment(t *testing.T) {
	logPrefix := "TestLoggingCreatePayment"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
		"Location":     "http://localhost/v1/1234/payments/000000123410000595980000100001",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/payments",
		createRecordRequest(http.StatusCreated, createPaymentJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	var card definitions.Card
	card.CardNumber = newString("1234567890123456")
	card.Cvv = newString("123")
	card.ExpiryDate = newString("1220")

	var cardPaymentMethodSpecificInput payment.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = newInt32(1)

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(2345)
	amountOfMoney.CurrencyCode = newString("CAD")

	var billingAddress definitions.Address
	billingAddress.CountryCode = newString("CA")

	var customer payment.Customer
	customer.BillingAddress = &billingAddress

	var order payment.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var reqBody payment.CreateRequest
	reqBody.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	reqBody.Order = &order

	response, err := client.Merchant("1234").Payments().Create(reqBody, nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	if response.Payment == nil {
		t.Fatalf("%v: responsePayment nil", logPrefix)
	}
	if response.Payment.ID == nil {
		t.Fatalf("%v: responsePaymentID nil", logPrefix)
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "POST" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/payments" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if secondEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response.StatusCode() != http.StatusCreated {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, secondEntry.response.StatusCode())
	}
	if secondEntry.response.ContentType() != "application/json" {
		t.Fatalf("%v: secondEntryResponseContentType %v", logPrefix, secondEntry.response.ContentType())
	}
	if secondEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, secondEntry.response.Headers())
	}
	if secondEntry.response.Body() == "" {
		t.Fatalf("%v: secondEntryResponseBody %v", logPrefix, secondEntry.response.Body())
	}

	foundDate, foundDummy := false, false
	for k := range secondEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if secondEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}

func TestCreatePaymentInvalidCardNumber(t *testing.T) {
	logPrefix := "TestCreatePaymentInvalidCardNumber"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/payments",
		createRecordRequest(http.StatusBadRequest, createPaymentFailedInvalidCardNumberJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	var card definitions.Card
	card.CardNumber = newString("1234567890123456")
	card.Cvv = newString("123")
	card.ExpiryDate = newString("1220")

	var cardPaymentMethodSpecificInput payment.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = newInt32(1)

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(2345)
	amountOfMoney.CurrencyCode = newString("CAD")

	var billingAddress definitions.Address
	billingAddress.CountryCode = newString("CA")

	var customer payment.Customer
	customer.BillingAddress = &billingAddress

	var order payment.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var reqBody payment.CreateRequest
	reqBody.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	reqBody.Order = &order

	_, err = client.Merchant("1234").Payments().Create(reqBody, nil)
	switch ce := err.(type) {
	case *sdkErrors.ValidateError:
		{
			if ce.StatusCode() != http.StatusBadRequest {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.ResponseBody() != createPaymentFailedInvalidCardNumberJSON {
				t.Fatalf("%v: responseBody %v", logPrefix, ce.ResponseBody())
			}

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)

			break
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "POST" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/payments" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if secondEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response.StatusCode() != http.StatusBadRequest {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, secondEntry.response.StatusCode())
	}
	if secondEntry.response.ContentType() != "application/json" {
		t.Fatalf("%v: secondEntryResponseContentType %v", logPrefix, secondEntry.response.ContentType())
	}
	if secondEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, secondEntry.response.Headers())
	}
	foundDate, foundDummy := false, false
	for k := range secondEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if secondEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}

func TestCreatePaymentRejected(t *testing.T) {
	logPrefix := "TestCreatePaymentRejected"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/payments",
		createRecordRequest(http.StatusPaymentRequired, createPaymentFailedRejectedJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	var card definitions.Card
	card.CardNumber = newString("1234567890123456")
	card.Cvv = newString("123")
	card.ExpiryDate = newString("1220")

	var cardPaymentMethodSpecificInput payment.CardPaymentMethodSpecificInput
	cardPaymentMethodSpecificInput.Card = &card
	cardPaymentMethodSpecificInput.PaymentProductID = newInt32(1)

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(2345)
	amountOfMoney.CurrencyCode = newString("CAD")

	var billingAddress definitions.Address
	billingAddress.CountryCode = newString("CA")

	var customer payment.Customer
	customer.BillingAddress = &billingAddress

	var order payment.Order
	order.AmountOfMoney = &amountOfMoney
	order.Customer = &customer

	var reqBody payment.CreateRequest
	reqBody.CardPaymentMethodSpecificInput = &cardPaymentMethodSpecificInput
	reqBody.Order = &order

	_, err = client.Merchant("1234").Payments().Create(reqBody, nil)
	switch ce := err.(type) {
	case *sdkErrors.DeclinedPaymentError:
		{
			if ce.StatusCode() != http.StatusPaymentRequired {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.ResponseBody() != createPaymentFailedRejectedJSON {
				t.Fatalf("%v: responseBody %v", logPrefix, ce.ResponseBody())
			}

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)

			break
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "POST" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/payments" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if secondEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response.StatusCode() != http.StatusPaymentRequired {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, secondEntry.response.StatusCode())
	}
	if secondEntry.response.ContentType() != "application/json" {
		t.Fatalf("%v: secondEntryResponseContentType %v", logPrefix, secondEntry.response.ContentType())
	}
	if secondEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, secondEntry.response.Headers())
	}
	foundDate, foundDummy := false, false
	for k := range secondEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if secondEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}

func TestLoggingUnknownServerError(t *testing.T) {
	logPrefix := "TestLoggingUnknownServerError"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/services/testconnection",
		createRecordRequest(http.StatusInternalServerError, unknownServerErrorJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	_, err = client.Merchant("1234").Services().Testconnection(nil)
	switch ce := err.(type) {
	case *sdkErrors.GlobalCollectError:
		{
			if ce.StatusCode() != http.StatusInternalServerError {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.ResponseBody() != unknownServerErrorJSON {
				t.Fatalf("%v: responseBody %v", logPrefix, ce.ResponseBody())
			}

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)

			break
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "GET" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/services/testconnection" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if secondEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response.StatusCode() != http.StatusInternalServerError {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, secondEntry.response.StatusCode())
	}
	if secondEntry.response.ContentType() != "application/json" {
		t.Fatalf("%v: secondEntryResponseContentType %v", logPrefix, secondEntry.response.ContentType())
	}
	if secondEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, secondEntry.response.Headers())
	}
	if secondEntry.response.Body() == "" {
		t.Fatalf("%v: secondEntryResponseBody %v", logPrefix, secondEntry.response.Body())
	}
	foundDate, foundDummy := false, false
	for k := range secondEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if secondEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}

func TestNonJson(t *testing.T) {
	logPrefix := "TestNonJson"

	responseHeaders := map[string]string{
		"Dummy": "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTestEnvironment(
		"/v1/1234/services/testconnection",
		createRecordRequest(http.StatusNotFound, notFoundErrorHTML, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	_, err = client.Merchant("1234").Services().Testconnection(nil)
	switch err.(type) {
	case *sdkErrors.NotFoundError:
		{
			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)

			break
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "GET" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/services/testconnection" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if secondEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, secondEntry.response)
	}
	if secondEntry.response.StatusCode() != http.StatusNotFound {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, secondEntry.response.StatusCode())
	}
	if secondEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, secondEntry.response.Headers())
	}
	if secondEntry.response.Body() == "" {
		t.Fatalf("%v: secondEntryResponseBody %v", logPrefix, secondEntry.response.Body())
	}
	foundDate, foundDummy := false, false
	for k := range secondEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if secondEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}

func TestReadTimeout(t *testing.T) {
	logPrefix := "TestReadTimeout"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, err := createTimedTestEnvironment(
		"/v1/1234/services/testconnection",
		createDelayedRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders, 1*time.Second),
		1*time.Millisecond,
		10*time.Millisecond)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}
	client.EnableLogging(logger)

	_, err = client.Merchant("1234").Services().Testconnection(nil)
	switch ce := err.(type) {
	case *sdkErrors.CommunicationError:
		{
			internalError := ce.InternalError()

			if uErr, ok := internalError.(*url.Error); ok && uErr.Timeout() {
				break
			}

			t.Fatalf("%v: %v", logPrefix, internalError)

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)

			break
		}
	}

	if len(logger.entries) != 2 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	firstEntry := logger.entries[0]
	if firstEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, firstEntry.request)
	}
	if firstEntry.request.Method() != "GET" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, firstEntry.request.Method())
	}
	if firstEntry.request.URL().Path != "/v1/1234/services/testconnection" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, firstEntry.request.URL().Path)
	}
	if firstEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, firstEntry.request.Headers())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range firstEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if firstEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, firstEntry.err)
	}

	secondEntry := logger.entries[1]
	if secondEntry.err == nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, secondEntry.err)
	}
}

func TestLogRequestOnly(t *testing.T) {
	logPrefix := "TestLogRequestOnly"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, mux, err := createEmptyTestEnvironment()
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	mux.HandleFunc("/v1/1234/services/testconnection",
		createNonLoggedRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders, client))

	logger := &testLogger{}
	client.EnableLogging(logger)

	_, err = client.Merchant("1234").Services().Testconnection(nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 1 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	requestEntry := logger.entries[0]
	if requestEntry.request == nil {
		t.Fatalf("%v: firstEntryRequest %v", logPrefix, requestEntry.request)
	}
	if requestEntry.request.Method() != "GET" {
		t.Fatalf("%v: firstEntryRequestMethod %v", logPrefix, requestEntry.request.Method())
	}
	if requestEntry.request.URL().Path != "/v1/1234/services/testconnection" {
		t.Fatalf("%v: firstEntryRequestURL %v", logPrefix, requestEntry.request.URL().Path)
	}
	if requestEntry.request.Headers() == nil {
		t.Fatalf("%v: firstEntryRequestHeaders %v", logPrefix, requestEntry.request.Headers())
	}
	foundDate, foundMetainfo := false, false
	for k, v := range requestEntry.request.Headers() {
		switch k {
		case "Authorization":
			{
				if v[0] != "********" {
					t.Fatalf("%v: authorizationHeader %v", logPrefix, v)
				}

				break
			}
		case "Date":
			{
				foundDate = true

				break
			}
		case "X-GCS-ServerMetaInfo":
			{
				foundMetainfo = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundMetainfo {
		t.Fatalf("%v: meta info header not found", logPrefix)
	}
	if requestEntry.err != nil {
		t.Fatalf("%v: firstEntryErr %v", logPrefix, requestEntry.err)
	}
}

func TestLogResponseOnly(t *testing.T) {
	logPrefix := "TestLogResponseOnly"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, mux, err := createEmptyTestEnvironment()
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}

	mux.HandleFunc("/v1/1234/services/testconnection",
		createLoggedRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders, client, logger))

	_, err = client.Merchant("1234").Services().Testconnection(nil)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if len(logger.entries) != 1 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	responseEntry := logger.entries[0]
	if responseEntry.response == nil {
		t.Fatalf("%v: secondEntryResponse %v", logPrefix, responseEntry.response)
	}
	if responseEntry.response.StatusCode() != http.StatusOK {
		t.Fatalf("%v: secondEntryResponseStatusCode %v", logPrefix, responseEntry.response.StatusCode())
	}
	if responseEntry.response.ContentType() != "application/json" {
		t.Fatalf("%v: secondEntryResponseContentType %v", logPrefix, responseEntry.response.ContentType())
	}
	if responseEntry.response.Headers() == nil {
		t.Fatalf("%v: secondEntryResponseHeaders %v", logPrefix, responseEntry.response.Headers())
	}
	if responseEntry.response.Body() == "" {
		t.Fatalf("%v: secondEntryResponseBody %v", logPrefix, responseEntry.response.Body())
	}
	foundDate, foundDummy := false, false
	for k := range responseEntry.response.Headers() {
		switch k {
		case "Date":
			{
				foundDate = true

				break
			}
		case "Dummy":
			{
				foundDummy = true

				break
			}
		}
	}
	if !foundDate {
		t.Fatalf("%v: date header not found", logPrefix)
	}
	if !foundDummy {
		t.Fatalf("%v: dummy header not found", logPrefix)
	}
	if responseEntry.err != nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, responseEntry.err)
	}
}

func TestLogErrorOnly(t *testing.T) {
	logPrefix := "TestLogErrorOnly"

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
		"Dummy":        "dummy",
	}
	requestHeaders := map[string][]string{}

	listener, sl, client, mux, err := createEmptyTimedTestEnvironment(
		20*time.Millisecond,
		50*time.Millisecond)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	logger := &testLogger{}

	mux.HandleFunc("/v1/1234/services/testconnection",
		createLoggedDelayedRecordRequest(http.StatusOK, testConnectionJSON, responseHeaders, requestHeaders, 1*time.Second, client, logger))

	_, err = client.Merchant("1234").Services().Testconnection(nil)
	switch ce := err.(type) {
	case *sdkErrors.CommunicationError:
		{
			internalError := ce.InternalError()

			if uErr, ok := internalError.(*url.Error); ok && uErr.Timeout() {
				break
			}

			t.Fatalf("%v: %v", logPrefix, internalError)

			break
		}
	default:
		{
			t.Fatalf("%v: %v", logPrefix, err)

			break
		}
	}

	if len(logger.entries) != 1 {
		t.Fatalf("%v: loggerEntries %v", logPrefix, len(logger.entries))
	}

	errorEntry := logger.entries[0]
	if errorEntry.err == nil {
		t.Fatalf("%v: secondEntryErr %v", logPrefix, errorEntry.err)
	}
}

func createNonLoggedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, client *Client) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		client.DisableLogging()

		for k, v := range r.Header {
			if k == "X-Gcs-Idempotence-Key" {
				k = "X-GCS-Idempotence-Key"
			}

			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		rw.Write([]byte(body))
	}
}

func createLoggedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, client *Client, logger logging.CommunicatorLogger) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		client.EnableLogging(logger)

		for k, v := range r.Header {
			if k == "X-Gcs-Idempotence-Key" {
				k = "X-GCS-Idempotence-Key"
			}

			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		rw.Write([]byte(body))
	}
}

func createDelayedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, delay time.Duration) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)

		for k, v := range r.Header {
			if k == "X-Gcs-Idempotence-Key" {
				k = "X-GCS-Idempotence-Key"
			}

			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		rw.Write([]byte(body))
	}
}

func createLoggedDelayedRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string, delay time.Duration, client *Client, logger logging.CommunicatorLogger) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		client.EnableLogging(logger)
		time.Sleep(delay)

		for k, v := range r.Header {
			if k == "X-Gcs-Idempotence-Key" {
				k = "X-GCS-Idempotence-Key"
			}

			requestHeaders[k] = v
		}

		for k, v := range responseHeaders {
			rw.Header()[k] = []string{v}
		}

		rw.WriteHeader(statusCode)

		rw.Write([]byte(body))
	}
}

func createTimedTestEnvironment(path string, handleFunc http.HandlerFunc, socketTimeout, connectTimeout time.Duration) (net.Listener, *stoppableListener, *Client, error) {
	mux := http.NewServeMux()
	mux.Handle(path, handleFunc)

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1 << 15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, err
	}

	client, err := createClient(socketTimeout, connectTimeout, randomPort)
	if err != nil {
		return nil, nil, nil, err
	}

	return listener, sl, client, nil
}

func createEmptyTestEnvironment() (net.Listener, *stoppableListener, *Client, *http.ServeMux, error) {
	mux := http.NewServeMux()

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1 << 15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	client, err := createClient(50*time.Second, 50*time.Second, randomPort)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return listener, sl, client, mux, nil
}

func createEmptyTimedTestEnvironment(socketTimeout, connectTimeout time.Duration) (net.Listener, *stoppableListener, *Client, *http.ServeMux, error) {
	mux := http.NewServeMux()

	httpServer := &http.Server{
		Handler: mux,
	}

	randomPort := (1 << 12) + rand.Intn(1 << 15)
	listener, err := net.Listen("tcp", ":"+strconv.Itoa(randomPort))
	if err != nil {
		return nil, nil, nil, nil, err
	}

	sl, err := mockServer(httpServer, listener)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	client, err := createClient(socketTimeout, connectTimeout, randomPort)
	if err != nil {
		return nil, nil, nil, nil, err
	}

	return listener, sl, client, mux, nil
}

type testLogger struct {
	entries []testLoggerEntry
}

func (t *testLogger) Log(message string) {
	t.entries = append(t.entries, testLoggerEntry{message, nil, nil, nil})
}

func (t *testLogger) LogError(message string, err error) {
	t.entries = append(t.entries, testLoggerEntry{message, err, nil, nil})
}

func (t *testLogger) LogResponseLogMessage(response *logging.ResponseLogMessage) {
	t.entries = append(t.entries, testLoggerEntry{"", nil, response, nil})
}

func (t *testLogger) LogRequestLogMessage(request *logging.RequestLogMessage) {
	t.entries = append(t.entries, testLoggerEntry{"", nil, nil, request})
}

type testLoggerEntry struct {
	message  string
	err      error
	response *logging.ResponseLogMessage
	request  *logging.RequestLogMessage
}
