package connectsdk

import (
	"crypto/rand"
	"fmt"
	"net/url"
	"os"
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
	"github.com/Ingenico-ePayments/connect-sdk-go/configuration"
	"github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/riskassessments"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/sessions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/token"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/productgroups"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/products"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/services"
	"github.com/Ingenico-ePayments/connect-sdk-go/merchant/tokens"
)

var envMerchantID = os.Getenv("connect.api.merchantId")
var envAPIKeyID = os.Getenv("connect.api.apiKeyId")
var envSecretAPIKey = os.Getenv("connect.api.secretApiKey")
var envProxyURL = os.Getenv("connect.api.proxyUrl")

func TestIntegratedConvertAmount(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	var query services.ConvertAmountParams
	query.Amount = newInt64(123)
	query.Source = newString("USD")
	query.Target = newString("EUR")

	response, err := client.Merchant(envMerchantID).Services().ConvertAmount(query, nil)
	if err != nil {
		t.Fatal(err)
	}

	if response.ConvertedAmount == nil {
		t.Fatal("nil converted amount")
	}
}

func TestIntegratedConnection(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	response, err := client.Merchant(envMerchantID).Services().Testconnection(nil)
	if err != nil {
		t.Fatal(err)
	}

	if response.Result == nil {
		t.Fatal("nil result")
	}
}

func TestIntegratedPayment(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	var query payment.CreateRequest
	var order payment.Order

	var amountOfMoney definitions.AmountOfMoney
	amountOfMoney.Amount = newInt64(100)
	amountOfMoney.CurrencyCode = newString("EUR")
	order.AmountOfMoney = &amountOfMoney

	var customer payment.Customer
	customer.Locale = newString("en")

	var address definitions.Address
	address.CountryCode = newString("NL")
	customer.BillingAddress = &address

	order.Customer = &customer
	query.Order = &order

	var paymentMethodSpecificInput payment.RedirectPaymentMethodSpecificInput
	paymentMethodSpecificInput.ReturnURL = newString("http://example.com/")
	paymentMethodSpecificInput.PaymentProductID = newInt32(809)

	var paymentProductSpecificInput payment.RedirectPaymentProduct809SpecificInput
	paymentProductSpecificInput.IssuerID = newString("INGBNL2A")
	paymentMethodSpecificInput.PaymentProduct809SpecificInput = &paymentProductSpecificInput

	query.RedirectPaymentMethodSpecificInput = &paymentMethodSpecificInput

	idempotenceKey, err := pseudoUUID()
	if err != nil {
		t.Fatal(err)
	}

	context, err := NewCallContext(idempotenceKey)
	if err != nil {
		t.Fatal(err)
	}

	result, err := doCreatePayment(client, query, context)
	if err != nil {
		t.Fatal(err)
	}
	paymentID := result.Payment.ID
	status := result.Payment.Status

	if idempotenceKey != context.IdempotenceKey {
		t.Fatalf("idempotence key mismatch")
	}
	if context.IdempotenceRequestTimestamp != nil {
		t.Fatalf("timestamp not nil")
	}

	secondResult, err := doCreatePayment(client, query, context)
	if idempotenceKey != context.IdempotenceKey {
		t.Fatalf("idempotence key mismatch")
	}
	if context.IdempotenceRequestTimestamp == nil {
		t.Fatalf("timestamp nil")
	}

	if *secondResult.Payment.ID != *paymentID {
		t.Fatalf("payment id mismatch")
	}
	if *secondResult.Payment.Status != *status {
		t.Fatalf("status mismatch")
	}
}

func TestIntegratedPaymentProducts(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	lParams := products.NewFindParams()
	lParams.CountryCode = newString("NL")
	lParams.CurrencyCode = newString("EUR")

	_, err = client.Merchant(envMerchantID).Products().Find(*lParams, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedPaymentProductDirectories(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	lParams := products.NewDirectoryParams()
	lParams.CountryCode = newString("NL")
	lParams.CurrencyCode = newString("EUR")

	_, err = client.Merchant(envMerchantID).Products().Directory(809, *lParams, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedPaymentProductsGroups(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	lParams := productgroups.NewGetParams()
	lParams.CountryCode = newString("NL")
	lParams.CurrencyCode = newString("EUR")

	_, err = client.Merchant(envMerchantID).Productgroups().Get("cards", *lParams, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedCreateSession(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	lParams := sessions.NewSessionRequest()

	_, err = client.Merchant(envMerchantID).Sessions().Create(*lParams, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedCreateToken(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	var query token.CreateRequest
	query.PaymentProductID = newInt32(1)

	tokenCard := token.NewCard()
	query.Card = tokenCard

	customerToken := token.NewCustomerToken()
	tokenCard.Customer = customerToken

	address := definitions.NewAddress()
	customerToken.BillingAddress = address
	address.CountryCode = newString("NL")

	mandate := token.NewCardData()
	tokenCard.Data = mandate

	cardWithoutCVV := definitions.NewCardWithoutCvv()
	mandate.CardWithoutCvv = cardWithoutCVV
	cardWithoutCVV.CardholderName = newString("Jan")
	cardWithoutCVV.IssueNumber = newString("12")
	cardWithoutCVV.CardNumber = newString("4567350000427977")
	cardWithoutCVV.ExpiryDate = newString("1225")

	response, err := client.Merchant(envMerchantID).Tokens().Create(query, nil)
	if err != nil {
		t.Fatal(err)
	}

	if response.Token == nil {
		t.Fatal("nil token")
	}

	var query2 tokens.DeleteParams

	err = client.Merchant(envMerchantID).Tokens().Delete(*response.Token, query2, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIntegratedRiskAssessments(t *testing.T) {
	skipTestIfNeeded(t)

	client, err := getClientIntegration()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	var query riskassessments.RiskAssessmentBankAccount

	bankAccountBBan := definitions.NewBankAccountBban()
	bankAccountBBan.CountryCode = newString("DE")
	bankAccountBBan.AccountNumber = newString("0532013000")
	bankAccountBBan.BankCode = newString("37040044")
	query.BankAccountBban = bankAccountBBan

	order := riskassessments.NewOrderRiskAssessment()

	amountOfMoney := definitions.NewAmountOfMoney()
	amountOfMoney.Amount = newInt64(100)
	amountOfMoney.CurrencyCode = newString("EUR")
	order.AmountOfMoney = amountOfMoney

	customer := riskassessments.NewCustomerRiskAssessment()
	customer.Locale = newString("en_GB")
	order.Customer = customer

	query.Order = order

	response, err := client.Merchant(envMerchantID).Riskassessments().Bankaccounts(query, nil)
	if err != nil {
		t.Fatal(err)
	}

	if response.Results == nil {
		t.Fatal("nil results")
	}
	if len(*response.Results) == 0 {
		t.Fatal("empty results")
	}
}

func TestIntegratedMultilineHeader(t *testing.T) {
	//skipTestIfNeeded(t)
	t.Skip("fails inside sandbox")

	configuration := configuration.DefaultConfiguration(envAPIKeyID, envSecretAPIKey, "Ingenico")
	configuration.APIEndpoint.Host = "eu.sandbox.api-ingenico.com"
	configuration.APIKeyID = envAPIKeyID
	configuration.SecretAPIKey = envSecretAPIKey

	connection, err := defaultimpl.NewDefaultConnection(configuration.SocketTimeout,
		configuration.ConnectTimeout,
		configuration.KeepAliveTimeout,
		configuration.IdleTimeout,
		configuration.MaxConnections,
		configuration.Proxy)
	if err != nil {
		t.Fatal(err)
	}

	multiLineHeader := newHeader("X-GCS-MultiLineHeader", "some\nvalue")

	metaDataProviderBuilder := communicator.NewMetaDataProviderBuilder(configuration.Integrator)
	metaDataProviderBuilder.ShoppingCartExtension = configuration.ShoppingCartExtension
	metaDataProviderBuilder.AdditionalRequestHeaders = append(metaDataProviderBuilder.AdditionalRequestHeaders, multiLineHeader)

	metaDataProvider, err := metaDataProviderBuilder.Build()
	if err != nil {
		t.Fatal(err)
	}

	authenticator, err := defaultimpl.NewDefaultAuthenticator(configuration.AuthorizationType,
		configuration.APIKeyID,
		configuration.SecretAPIKey)
	if err != nil {
		t.Fatal(err)
	}

	session, err := communicator.NewSession(&configuration.APIEndpoint, connection, authenticator, metaDataProvider)
	if err != nil {
		t.Fatal(err)
	}

	client, err := CreateClientFromSession(session)
	if err != nil {
		t.Fatal(err)
	}
	defer client.Close()

	_, err = client.Merchant(envMerchantID).Services().Testconnection(nil)
	if err != nil {
		t.Fatal(err)
	}
}

func getClientIntegration() (*Client, error) {
	configuration := configuration.DefaultConfiguration(envAPIKeyID, envSecretAPIKey, "Ingenico")
	configuration.APIEndpoint.Host = "eu.sandbox.api-ingenico.com"

	if len(envProxyURL) > 0 {
		proxy, err := url.Parse(envProxyURL)
		if err != nil {
			return nil, err
		}
		configuration.Proxy = proxy
	}

	client, err := CreateClientFromConfiguration(&configuration)
	if err != nil {
		return nil, err
	}

	return client.WithClientMetaInfo("{\"test\":\"test\"}")
}

func doCreatePayment(client *Client, query payment.CreateRequest, context *CallContext) (payment.CreateResult, error) {
	// For this test it doesn't matter if the response is successful or declined,
	// as long as idempotence is handled correctly

	response, err := client.Merchant(envMerchantID).Payments().Create(query, context)
	if err == nil {
		result := payment.CreateResult{
			CreationOutput: response.CreationOutput,
			MerchantAction: response.MerchantAction,
			Payment:        response.Payment,
		}
		return result, nil
	}
	switch realError := err.(type) {
	case *sdkErrors.DeclinedPaymentError:
		return *realError.PaymentResult(), nil
	default:
		return payment.CreateResult{}, nil
	}
}

func pseudoUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func skipTestIfNeeded(t *testing.T) {
	if len(envMerchantID) == 0 {
		t.Skip("empty env connect.api.merchantId")
	}
	if len(envAPIKeyID) == 0 {
		t.Skip("empty env connect.api.apiKeyId")
	}
	if len(envSecretAPIKey) == 0 {
		t.Skip("empty env connect.api.secretApiKey")
	}
}
