package connectsdk

import (
	"errors"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator"
	"github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/definitions"
	"github.com/Ingenico-ePayments/connect-sdk-go/domain/payment"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
)

var idempotenceSuccessJSON = `{
	"creationOutput": {
		"additionalReference": "00000200002014254946",
		"externalReference": "000002000020142549460000100001"
	},
	"payment": {
		"id": "000002000020142549460000100001",
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
					"cardNumber": "************9176",
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
			"statusCode": 600,
			"statusCodeChangeDateTime": "20150331120036",
			"isAuthorized": true
		}
	}
}`

var idempotenceRejectJSON = `{
	"errorId": "2c164323-20d3-4e9e-8578-dc562cd7506c-0000003c",
	"errors": [{
		"code": "21000020",
		"requestId": "2001798",
		"message": "VALUE **************** OF FIELD CREDITCARDNUMBER DID NOT PASS THE LUHNCHECK"
	}],
	"paymentResult": {
		"creationOutput": {
			"additionalReference": "00000200002014254436",
			"externalReference": "000002000020142544360000100001"
		},
		"payment": {
			"id": "000002000020142544360000100001",
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
					"code": "21000020",
					"requestId": "2001798",
					"message": "VALUE **************** OF FIELD CREDITCARDNUMBER DID NOT PASS THE LUHNCHECK"
				}],
				"isCancellable": false,
				"statusCode": 100,
				"statusCodeChangeDateTime": "20150330173151",
				"isAuthorized": false
			}
		}
	}
}`

var idempotenceDepulicateFailureJSON = `{
	"errorId": "75b0f13a-04a5-41b3-83b8-b295ddb23439-000013c6",
	"errors": [{
		"code": "1409",
		"message": "DUPLICATE REQUEST IN PROGRESS",
		"httpStatusCode": 409
	}]
}`

func createClient(socketTimeout, connectTimeout time.Duration, port int) (*Client, error) {
	connection, err := defaultimpl.NewDefaultConnection(socketTimeout, connectTimeout, 30*time.Second, 50*time.Second, 500, nil)
	if err != nil {
		return nil, err
	}

	authenticator, err := defaultimpl.NewDefaultAuthenticator(defaultimpl.V1HMAC, "apiKey", "secret")
	if err != nil {
		return nil, err
	}

	metaDataProvider, err := communicator.NewMetaDataProviderWithIntegrator("Ingenico")
	if err != nil {
		return nil, err
	}

	endPoint := &url.URL{
		Scheme: "http",
		Host:   "localhost:" + strconv.Itoa(port),
	}

	session, err := communicator.NewSession(endPoint, connection, authenticator, metaDataProvider)
	if err != nil {
		return nil, err
	}

	marshaller, _ := defaultimpl.NewDefaultMarshaller()

	communicator, err := communicator.NewCommunicator(session, marshaller)
	if err != nil {
		return nil, err
	}

	return NewClient(communicator)
}

func createRequest() payment.CreateRequest {
	body := payment.CreateRequest{}
	order := payment.NewOrder()

	amountOfMoney := definitions.NewAmountOfMoney()
	amountOfMoney.Amount = newInt64(2345)
	amountOfMoney.CurrencyCode = newString("CAD")
	order.AmountOfMoney = amountOfMoney

	customer := payment.NewCustomer()

	billingAddress := definitions.NewAddress()
	billingAddress.CountryCode = newString("CA")
	customer.BillingAddress = billingAddress

	order.Customer = customer

	cardPaymentMethodSpecificInput := payment.NewCardPaymentMethodSpecificInput()
	cardPaymentMethodSpecificInput.PaymentProductID = newInt32(1)

	card := definitions.NewCard()
	card.Cvv = newString("123")
	card.CardNumber = newString("4567350000427977")
	card.ExpiryDate = newString("1220")
	cardPaymentMethodSpecificInput.Card = card

	body.CardPaymentMethodSpecificInput = cardPaymentMethodSpecificInput

	return body
}

func newBool(val bool) *bool {
	return &val
}
func newInt32(val int32) *int32 {
	return &val
}
func newInt64(val int64) *int64 {
	return &val
}
func newString(val string) *string {
	return &val
}

type stoppableListener struct {
	*net.TCPListener
	stop     chan int
	finished sync.WaitGroup
}

var errStopped = errors.New("listener stopped")

func (sl *stoppableListener) Accept() (net.Conn, error) {
	sl.finished.Add(1)
	defer sl.finished.Done()

	for {
		sl.SetDeadline(time.Now().Add(time.Second))

		newConn, err := sl.TCPListener.Accept()

		select {
		case <-sl.stop:
			return nil, errStopped
		default:
		}

		if err != nil {
			netErr, ok := err.(net.Error)

			if ok && netErr.Timeout() && netErr.Temporary() {
				continue
			}
		}

		return newConn, err
	}
}

func (sl *stoppableListener) Stop() {
	close(sl.stop)
	sl.finished.Wait()
}

func newStoppableListener(l net.Listener) (*stoppableListener, error) {
	tcpL, ok := l.(*net.TCPListener)

	if !ok {
		return nil, errors.New("Cannot wrap listener")
	}

	return &stoppableListener{tcpL, make(chan int), sync.WaitGroup{}}, nil
}

func mockServer(server *http.Server, listener net.Listener) (*stoppableListener, error) {
	ls, err := newStoppableListener(listener)
	if err != nil {
		return nil, err
	}

	go server.Serve(ls)

	return ls, nil
}

func createRecordRequest(statusCode int, body string, responseHeaders map[string]string, requestHeaders map[string][]string) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randString(n int) string {
	b := make([]byte, n, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func createTestEnvironment(path string, handleFunc http.HandlerFunc) (net.Listener, *stoppableListener, *Client, error) {
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

	client, err := createClient(50*time.Second, 50*time.Second, randomPort)
	if err != nil {
		return nil, nil, nil, err
	}

	return listener, sl, client, nil
}

func TestIdempotenceFirstRequest(t *testing.T) {
	logPrefix := "TestIdempotenceFirstRequest"

	idempotenceKey := randString(32)

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
	}
	requestHeaders := map[string][]string{}

	context, _ := NewCallContext(idempotenceKey)

	listener, sl, client, err := createTestEnvironment(
		"/v1/20000/payments",
		createRecordRequest(http.StatusOK, idempotenceSuccessJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	request := createRequest()
	response, err := client.Merchant("20000").Payments().Create(request, context)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if response.Payment == nil {
		t.Fatalf("%v: nil payment", logPrefix)
	}
	if response.Payment.ID == nil {
		t.Fatalf("%v: nil paymentID", logPrefix)
	}

	if idempotenceKey != requestHeaders["X-GCS-Idempotence-Key"][0] {
		t.Fatalf("%v: idempotenceKey mismatch %v %v", logPrefix,
			idempotenceKey, requestHeaders["X-GCS-Idempotence-Key"][0])
	}

	if context.IdempotenceRequestTimestamp != nil {
		t.Fatalf("%v: timestamp not nil", logPrefix)
	}
}

func TestIdempotenceSecondRequest(t *testing.T) {
	logPrefix := "TestIdempotenceSecondRequest"

	idempotenceKey := randString(32)
	idempotenceTimeStamp := time.Now().Sub(time.Unix(0, 0)).Nanoseconds() / int64(time.Millisecond)

	responseHeaders := map[string]string{
		"Content-Type":                        "application/json",
		"X-GCS-Idempotence-Request-Timestamp": strconv.FormatInt(idempotenceTimeStamp, 10),
	}
	requestHeaders := map[string][]string{}

	context, _ := NewCallContext(idempotenceKey)
	context.IdempotenceKey = idempotenceKey

	listener, sl, client, err := createTestEnvironment(
		"/v1/20000/payments",
		createRecordRequest(http.StatusOK, idempotenceSuccessJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	request := createRequest()
	response, err := client.Merchant("20000").Payments().Create(request, context)
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}

	if response.Payment == nil {
		t.Fatalf("%v: nil payment", logPrefix)
	}
	if response.Payment.ID == nil {
		t.Fatalf("%v: nil paymentID", logPrefix)
	}

	if idempotenceKey != requestHeaders["X-GCS-Idempotence-Key"][0] {
		t.Fatalf("%v: idempotenceKey mismatch %v %v", logPrefix,
			idempotenceKey, requestHeaders["X-GCS-Idempotence-Key"][0])
	}

	if context.IdempotenceRequestTimestamp != nil && *context.IdempotenceRequestTimestamp != idempotenceTimeStamp {
		t.Fatalf("%v: timestamp %v", logPrefix, context.IdempotenceRequestTimestamp)
	}
}

func TestIdempotenceFirstFailure(t *testing.T) {
	logPrefix := "TestIdempotenceFirstFailure"

	idempotenceKey := randString(32)
	idempotenceTimeStamp := time.Now().Sub(time.Unix(0, 0)).Nanoseconds() / int64(time.Millisecond)

	responseHeaders := map[string]string{
		"Content-Type":                        "application/json",
		"X-GCS-Idempotence-Request-Timestamp": strconv.FormatInt(idempotenceTimeStamp, 10),
	}
	requestHeaders := map[string][]string{}

	context, _ := NewCallContext(idempotenceKey)
	context.IdempotenceKey = idempotenceKey

	listener, sl, client, err := createTestEnvironment(
		"/v1/20000/payments",
		createRecordRequest(http.StatusPaymentRequired, idempotenceRejectJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	request := createRequest()
	_, err = client.Merchant("20000").Payments().Create(request, context)
	switch ce := err.(type) {
	case *sdkErrors.DeclinedPaymentError:
		{
			if ce.StatusCode() != http.StatusPaymentRequired {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.ResponseBody() != idempotenceRejectJSON {
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

	if idempotenceKey != requestHeaders["X-GCS-Idempotence-Key"][0] {
		t.Fatalf("%v: idempotenceKey mismatch %v %v", logPrefix,
			idempotenceKey, requestHeaders["X-GCS-Idempotence-Key"][0])
	}

	if context.IdempotenceRequestTimestamp != nil && *context.IdempotenceRequestTimestamp != idempotenceTimeStamp {
		t.Fatalf("%v: timestamp %v", logPrefix, context.IdempotenceRequestTimestamp)
	}
}

func TestIdempotenceSecondFailure(t *testing.T) {
	logPrefix := "TestIdempotenceSecondFailure"

	idempotenceKey := randString(32)

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
	}
	requestHeaders := map[string][]string{}

	context, _ := NewCallContext(idempotenceKey)
	context.IdempotenceKey = idempotenceKey

	listener, sl, client, err := createTestEnvironment(
		"/v1/20000/payments",
		createRecordRequest(http.StatusPaymentRequired, idempotenceRejectJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	request := createRequest()
	_, err = client.Merchant("20000").Payments().Create(request, context)
	switch ce := err.(type) {
	case *sdkErrors.DeclinedPaymentError:
		{
			if ce.StatusCode() != http.StatusPaymentRequired {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.ResponseBody() != idempotenceRejectJSON {
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

	if idempotenceKey != requestHeaders["X-GCS-Idempotence-Key"][0] {
		t.Fatalf("%v: idempotenceKey mismatch %v %v", logPrefix,
			idempotenceKey, requestHeaders["X-GCS-Idempotence-Key"][0])
	}

	if context.IdempotenceRequestTimestamp != nil {
		t.Fatalf("%v: timestamp %v", logPrefix, context.IdempotenceRequestTimestamp)
	}
}

func TestIdempotenceDuplicateRequest(t *testing.T) {
	logPrefix := "TestIdempotenceDuplicateRequest"

	idempotenceKey := randString(32)

	responseHeaders := map[string]string{
		"Content-Type": "application/json",
	}
	requestHeaders := map[string][]string{}

	context, _ := NewCallContext(idempotenceKey)
	context.IdempotenceKey = idempotenceKey

	listener, sl, client, err := createTestEnvironment(
		"/v1/20000/payments",
		createRecordRequest(http.StatusConflict, idempotenceDepulicateFailureJSON, responseHeaders, requestHeaders))
	if err != nil {
		t.Fatalf("%v: %v", logPrefix, err)
	}
	defer listener.Close()
	defer sl.Close()
	defer client.Close()

	request := createRequest()
	_, err = client.Merchant("20000").Payments().Create(request, context)
	switch ce := err.(type) {
	case *sdkErrors.IdempotenceError:
		{
			if ce.StatusCode() != http.StatusConflict {
				t.Fatalf("%v: statusCode %v", logPrefix, ce.StatusCode())
			}
			if ce.ResponseBody() != idempotenceDepulicateFailureJSON {
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

	if idempotenceKey != requestHeaders["X-GCS-Idempotence-Key"][0] {
		t.Fatalf("%v: idempotenceKey mismatch %v %v", logPrefix,
			idempotenceKey, requestHeaders["X-GCS-Idempotence-Key"][0])
	}
	if idempotenceKey != context.IdempotenceKey {
		t.Fatalf("%v: idempotenceKey mismatch %v %v", logPrefix,
			idempotenceKey, context.IdempotenceKey)
	}

	if context.IdempotenceRequestTimestamp != nil {
		t.Fatalf("%v: timestamp %v", logPrefix, context.IdempotenceRequestTimestamp)
	}
}
