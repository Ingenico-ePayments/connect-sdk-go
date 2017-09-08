package webhooks

import (
	"encoding/json"
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/webhooks"
)

var (
	testValidBody = `{
  "apiVersion": "v1",
  "id": "8ee793f6-4553-4749-85dc-f2ef095c5ab0",
  "created": "2017-02-02T11:24:14.040+0100",
  "merchantId": "20000",
  "type": "payment.paid",
  "payment": {
    "id": "00000200000143570012",
    "paymentOutput": {
      "amountOfMoney": {
        "amount": 1000,
        "currencyCode": "EUR"
      },
      "references": {
        "paymentReference": "200001681810"
      },
      "paymentMethod": "bankTransfer",
      "bankTransferPaymentMethodSpecificOutput": {
        "paymentProductId": 11
      }
    },
    "status": "PAID",
    "statusOutput": {
      "isCancellable": false,
      "statusCategory": "COMPLETED",
      "statusCode": 1000,
      "statusCodeChangeDateTime": "20170202112414",
      "isAuthorized": true
    }
  }
}`
	testInvalidBody = `{
  "apiVersion": "v1",
  "id": "8ee793f6-4553-4749-85dc-f2ef095c5ab0",
  "created": "2017-02-02T11:25:14.040+0100",
  "merchantId": "20000",
  "type": "payment.paid",
  "payment": {
    "id": "00000200000143570012",
    "paymentOutput": {
      "amountOfMoney": {
        "amount": 1000,
        "currencyCode": "EUR"
      },
      "references": {
        "paymentReference": "200001681810"
      },
      "paymentMethod": "bankTransfer",
      "bankTransferPaymentMethodSpecificOutput": {
        "paymentProductId": 11
      }
    },
    "status": "PAID",
    "statusOutput": {
      "isCancellable": false,
      "statusCategory": "COMPLETED",
      "statusCode": 1000,
      "statusCodeChangeDateTime": "20170202112514",
      "isAuthorized": true
    }
  }
}`

	testSignature = "2S7doBj/GnJnacIjSJzr5fxGM5xmfQyFAwxv1I53ZEk="
	testKeyID     = "dummy-key-id"
	testSecretKey = "hello+world"
)

var testSignatureHeader, testKeyHeader communication.Header

type fakeMarshaller struct{}

func (m *fakeMarshaller) Marshal(v interface{}) (string, error) {
	dataBytes, err := json.Marshal(v)
	data := string(dataBytes)

	return data, err
}

func (m *fakeMarshaller) Unmarshal(data string, v interface{}) error {
	if len(data) < 1 {
		return nil
	}

	err := json.Unmarshal([]byte(data), v)
	if err != nil {
		return err
	}

	fakeVersion := "v0"

	event, ok := v.(*webhooks.Event)
	if ok {
		event.APIVersion = &fakeVersion
	}

	return nil
}

func TestUnmarshalAPIVersionMismatch(t *testing.T) {
	keyStore := &InMemorySecretKeyStore{}
	keyStore.Clear()
	keyStore.StoreSecretKey(testKeyID, testSecretKey)

	helperBuilder, err := CreateHelperBuilder(keyStore)
	if err != nil {
		t.Fatal(err)
	}

	helper, err := helperBuilder.WithMarshaller(&fakeMarshaller{}).Build()
	if err != nil {
		t.Fatal(err)
	}

	requestHeaders := []communication.Header{
		testSignatureHeader,
		testKeyHeader,
	}

	_, err = helper.Unmarshal(testValidBody, requestHeaders)
	if err == nil {
		t.Fatal("nil error")
	}

	_, ok := err.(*APIVersionMismatchError)
	if !ok {
		t.Fatal("wrong error returned")
	}
}

func TestUnmarshalNoSecretKeyAvailable(t *testing.T) {
	keyStore := &InMemorySecretKeyStore{}
	keyStore.Clear()

	helper, err := CreateHelper(keyStore)
	if err != nil {
		t.Fatal(err)
	}

	requestHeaders := []communication.Header{
		testSignatureHeader,
		testKeyHeader,
	}

	_, err = helper.Unmarshal(testValidBody, requestHeaders)

	if err == nil {
		t.Fatal("nil error")
	}

	_, ok := err.(*SecretKeyNotAvailableError)
	if !ok {
		t.Fatal("wrong error returned")
	}
}

func TestUnmarshalMissingHeaders(t *testing.T) {
	keyStore := &InMemorySecretKeyStore{}
	keyStore.Clear()

	helper, err := CreateHelper(keyStore)
	if err != nil {
		t.Fatal(err)
	}

	requestHeaders := []communication.Header{}

	_, err = helper.Unmarshal(testValidBody, requestHeaders)

	if err == nil {
		t.Fatal("nil error")
	}

	_, ok := err.(*SignatureValidationError)
	if !ok {
		t.Fatal("wrong error returned")
	}
}

func TestUnmarshalDuplicateHeaders(t *testing.T) {
	keyStore := &InMemorySecretKeyStore{}
	keyStore.Clear()

	helper, err := CreateHelper(keyStore)
	if err != nil {
		t.Fatal(err)
	}

	requestHeaders := []communication.Header{
		testSignatureHeader,
		testKeyHeader,
		testSignatureHeader,
	}

	_, err = helper.Unmarshal(testValidBody, requestHeaders)

	if err == nil {
		t.Fatal("nil error")
	}

	_, ok := err.(*SignatureValidationError)
	if !ok {
		t.Fatal("wrong error returned")
	}
}

func TestUnmarshalSuccess(t *testing.T) {
	keyStore := &InMemorySecretKeyStore{}
	keyStore.Clear()
	keyStore.StoreSecretKey(testKeyID, testSecretKey)

	helper, err := CreateHelper(keyStore)
	if err != nil {
		t.Fatal(err)
	}

	requestHeaders := []communication.Header{
		testSignatureHeader,
		testKeyHeader,
	}

	event, err := helper.Unmarshal(testValidBody, requestHeaders)
	if err != nil {
		t.Fatal(err)
	}

	if event == nil {
		t.Fatal("nil event")
	}

	if event.APIVersion == nil || *event.APIVersion != "v1" {
		t.Fatalf("apiVersion mismatch %v", event.APIVersion)
	}
	if event.ID == nil || *event.ID != "8ee793f6-4553-4749-85dc-f2ef095c5ab0" {
		t.Fatalf("ID mismatch %v", event.ID)
	}
	if event.Created == nil || *event.Created != "2017-02-02T11:24:14.040+0100" {
		t.Fatalf("created mismatch %v", event.Created)
	}
	if event.MerchantID == nil || *event.MerchantID != "20000" {
		t.Fatalf("merchantID mismatch %v", event.MerchantID)
	}
	if event.Type == nil || *event.Type != "payment.paid" {
		t.Fatalf("type mismatch %v", event.Type)
	}

	if event.Refund != nil {
		t.Fatalf("non-nil refund %v", event.Refund)
	}
	if event.Payout != nil {
		t.Fatalf("non-nil payout %v", event.Payout)
	}
	if event.Token != nil {
		t.Fatalf("non-nil token %v", event.Token)
	}

	if event.Payment == nil {
		t.Fatal("nil paymentResponse")
	}
	if event.Payment.ID != nil && *event.Payment.ID != "00000200000143570012" {
		t.Fatalf("payment.ID mismatch %v", event.Payment.ID)
	}
	if event.Payment.PaymentOutput == nil {
		t.Fatal("nil payment.PaymentOutput")
	}
	if event.Payment.PaymentOutput.AmountOfMoney == nil {
		t.Fatal("nil payment.PaymentOutput.AmountOfMoney")
	}
	if event.Payment.PaymentOutput.AmountOfMoney.Amount == nil || *event.Payment.PaymentOutput.AmountOfMoney.Amount != 1000 {
		t.Fatalf("payment.PaymentOutput.AmountOfMoney.Amount mismatch %v", event.Payment.PaymentOutput.AmountOfMoney.Amount)
	}
	if event.Payment.PaymentOutput.AmountOfMoney.CurrencyCode == nil || *event.Payment.PaymentOutput.AmountOfMoney.CurrencyCode != "EUR" {
		t.Fatalf("payment.PaymentOutput.AmountOfMoney.CurrencyCode mismatch %v", event.Payment.PaymentOutput.AmountOfMoney.CurrencyCode)
	}
	if event.Payment.PaymentOutput.References == nil {
		t.Fatal("nil payment.PaymentOutput.References")
	}
	if event.Payment.PaymentOutput.References.PaymentReference == nil ||
		*event.Payment.PaymentOutput.References.PaymentReference != "200001681810" {
		t.Fatalf("payment.paymentOutput.References.PaymentReference mismatch %v", event.Payment.PaymentOutput.References.PaymentReference)
	}
	if event.Payment.PaymentOutput.PaymentMethod == nil || *event.Payment.PaymentOutput.PaymentMethod != "bankTransfer" {
		t.Fatalf("payment.paymentOutput.PaymentMethod mismatch %v", event.Payment.PaymentOutput.PaymentMethod)
	}

	if event.Payment.PaymentOutput.CardPaymentMethodSpecificOutput != nil {
		t.Fatalf("non-nil payment.PaymentOutput.CardPaymentMethodSpecificOutput %v", event.Payment.PaymentOutput.CardPaymentMethodSpecificOutput)
	}
	if event.Payment.PaymentOutput.DirectDebitPaymentMethodSpecificOutput != nil {
		t.Fatalf("non-nil payment.PaymentOutput.DirectDebitPaymentMethodSpecificOutput %v", event.Payment.PaymentOutput.DirectDebitPaymentMethodSpecificOutput)
	}
	if event.Payment.PaymentOutput.InvoicePaymentMethodSpecificOutput != nil {
		t.Fatalf("non-nil payment.PaymentOutput.InvoicePaymentMethodSpecificOutput %v", event.Payment.PaymentOutput.InvoicePaymentMethodSpecificOutput)
	}
	if event.Payment.PaymentOutput.RedirectPaymentMethodSpecificOutput != nil {
		t.Fatalf("non-nil payment.PaymentOutput.RedirectPaymentMethodSpecificOutput %v", event.Payment.PaymentOutput.RedirectPaymentMethodSpecificOutput)
	}
	if event.Payment.PaymentOutput.SepaDirectDebitPaymentMethodSpecificOutput != nil {
		t.Fatalf("non-nil payment.PaymentOutput.SepaDirectDebitPaymentMethodSpecificOutput %v", event.Payment.PaymentOutput.SepaDirectDebitPaymentMethodSpecificOutput)
	}
	if event.Payment.PaymentOutput.BankTransferPaymentMethodSpecificOutput == nil {
		t.Fatal("nil payment.PaymentOutput.BankTransferPaymentMethodSpecificOutput")
	}
	if event.Payment.PaymentOutput.BankTransferPaymentMethodSpecificOutput.PaymentProductID != nil &&
		*event.Payment.PaymentOutput.BankTransferPaymentMethodSpecificOutput.PaymentProductID != 11 {
		t.Fatalf("payment.PaymentOutput.BankTransferPaymentMethodSpecificOutput.PaymentProductID mismatch %v", event.Payment.PaymentOutput.BankTransferPaymentMethodSpecificOutput.PaymentProductID)
	}

	if event.Payment.Status == nil || *event.Payment.Status != "PAID" {
		t.Fatalf("payment.Status mismatch %v", event.Payment.Status)
	}
	if event.Payment.StatusOutput == nil {
		t.Fatal("nil payment.StatusOutput")
	}
	if event.Payment.StatusOutput.IsCancellable == nil || *event.Payment.StatusOutput.IsCancellable != false {
		t.Fatalf("payment.StatusOutput.IsCancellable mismatch %v", event.Payment.StatusOutput.IsCancellable)
	}
	if event.Payment.StatusOutput.StatusCategory == nil || *event.Payment.StatusOutput.StatusCategory != "COMPLETED" {
		t.Fatalf("payment.StatusOutput.StatusCategory mismatch %v", event.Payment.StatusOutput.StatusCategory)
	}
	if event.Payment.StatusOutput.StatusCode == nil || *event.Payment.StatusOutput.StatusCode != 1000 {
		t.Fatalf("payment.StatusOutput.StatusCode mismatch %v", event.Payment.StatusOutput.StatusCode)
	}
	if event.Payment.StatusOutput.StatusCodeChangeDateTime == nil || *event.Payment.StatusOutput.StatusCodeChangeDateTime != "20170202112414" {
		t.Fatalf("payment.StatusOutput.StatusCodeChangeDateTime mismatch %v", event.Payment.StatusOutput.StatusCodeChangeDateTime)
	}
	if event.Payment.StatusOutput.IsAuthorized == nil || *event.Payment.StatusOutput.IsAuthorized != true {
		t.Fatalf("payment.StatusOutput.IsAuthorized mismatch %v", event.Payment.StatusOutput.IsAuthorized)
	}
}

func TestUnmarshalInvalidBody(t *testing.T) {
	keyStore := &InMemorySecretKeyStore{}
	keyStore.Clear()
	keyStore.StoreSecretKey(testKeyID, testSecretKey)

	helper, err := CreateHelper(keyStore)
	if err != nil {
		t.Fatal(err)
	}

	requestHeaders := []communication.Header{
		testSignatureHeader,
		testKeyHeader,
	}

	_, err = helper.Unmarshal(testInvalidBody, requestHeaders)
	if err == nil {
		t.Fatal("nil error")
	}

	_, ok := err.(*SignatureValidationError)
	if !ok {
		t.Fatal("wrong error returned")
	}
}

func TestUnmarshalInvalidSecretKey(t *testing.T) {
	keyStore := &InMemorySecretKeyStore{}
	keyStore.Clear()
	keyStore.StoreSecretKey(testKeyID, "1"+testSecretKey)

	helper, err := CreateHelper(keyStore)
	if err != nil {
		t.Fatal(err)
	}

	requestHeaders := []communication.Header{
		testSignatureHeader,
		testKeyHeader,
	}

	_, err = helper.Unmarshal(testValidBody, requestHeaders)
	if err == nil {
		t.Fatal("nil error")
	}

	_, ok := err.(*SignatureValidationError)
	if !ok {
		t.Fatal("wrong error returned")
	}
}

func TestUnmarshalInvalidSignature(t *testing.T) {
	keyStore := &InMemorySecretKeyStore{}
	keyStore.Clear()
	keyStore.StoreSecretKey(testKeyID, testSecretKey)

	helper, err := CreateHelper(keyStore)
	if err != nil {
		t.Fatal(err)
	}

	testHeader, err := communication.NewHeader("X-GCS-Signature", "1"+testSignature)
	if err != nil {
		t.Fatal(err)
	}

	requestHeaders := []communication.Header{
		*testHeader,
		testKeyHeader,
	}

	_, err = helper.Unmarshal(testValidBody, requestHeaders)
	if err == nil {
		t.Fatal("nil error")
	}

	_, ok := err.(*SignatureValidationError)
	if !ok {
		t.Fatal("wrong error returned")
	}
}

func init() {
	signatureHeader, err := communication.NewHeader("X-GCS-Signature", testSignature)
	if err != nil {
		panic(err)
	}

	keyHeader, err := communication.NewHeader("X-GCS-KeyId", testKeyID)
	if err != nil {
		panic(err)
	}

	testSignatureHeader = *signatureHeader
	testKeyHeader = *keyHeader
}
