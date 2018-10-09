package defaultimpl

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
)

func TestToCanonicalizedHeaderValue(t *testing.T) {
	result, err := toCanonicalizedHeaderValue("aap\nnoot  ")
	if err != nil {
		t.Fatalf("TestToCanonicalizedHeaderValue : %v", err)
	}
	if result != "aap noot" {
		t.Fatalf("TestToCanonicalizedHeaderValue : wrong canonization 1 '%s' != '%s'",
			result, "aap noot")
	}

	result, err = toCanonicalizedHeaderValue("aap\r\n  noot")
	if err != nil {
		t.Fatalf("TestToCanonicalizedHeaderValue : %v", err)
	}
	if result != "aap noot" {
		t.Fatalf("TestToCanonicalizedHeaderValue : wrong canonization 2 '%s' != '%s'",
			result, "aap noot")
	}

	result, err = toCanonicalizedHeaderValue(" some value  \r\n \n with  some \r\n  spaces ")
	if err != nil {
		t.Fatalf("TestToCanonicalizedHeaderValue : %v", err)
	}
	if result != "some value    with  some  spaces" {
		t.Fatalf("TestToCanonicalizedHeaderValue : wrong canonization 3 '%s' != '%s'",
			result, "some value    with  some  spaces")
	}
}

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

func TestToDataToSign(t *testing.T) {
	httpHeaders := []communication.Header{}
	serverMetaInfoHeader, err := communication.NewHeader("X-GCS-ServerMetaInfo", "{\"platformIdentifier\":\"Windows 7/6.1 Java/1.7 (Oracle Corporation; Java HotSpot(TM) 64-Bit Server VM; 1.7.0_45)\",\"sdkIdentifier\":\"1.0\"}")
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}
	httpHeaders = append(httpHeaders, *serverMetaInfoHeader)

	contentTypeHeader, err := communication.NewHeader("Content-Type", "application/json")
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}
	httpHeaders = append(httpHeaders, *contentTypeHeader)

	clientMetaInfoHeader, err := communication.NewHeader("X-GCS-ClientMetaInfo", "{\"aap\",\"noot\"}")
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}
	httpHeaders = append(httpHeaders, *clientMetaInfoHeader)

	userAgentHeader, err := communication.NewHeader("User-Agent", "Apache-HttpClient/4.3.4 (java 1.5)")
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}
	httpHeaders = append(httpHeaders, *userAgentHeader)

	dateHeader, err := communication.NewHeader("Date", "Mon, 07 Jul 2014 12:12:40 GMT")
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}
	httpHeaders = append(httpHeaders, *dateHeader)

	urlArg, err := url.Parse("http://localhost:8080/v1/9991/services%20bla/convert/amount?aap=noot&mies=geen%20noot")
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}

	dataToSign, err := toDataToSign("POST", *urlArg, httpHeaders)
	if err != nil {
		t.Fatalf("TestToDataToSign : %v", err)
	}

	expectedStart := "POST\napplication/json\n"
	expectedEnd := "x-gcs-clientmetainfo:{\"aap\",\"noot\"}\nx-gcs-servermetainfo:{\"platformIdentifier\":\"Windows 7/6.1 Java/1.7 (Oracle Corporation; Java HotSpot(TM) 64-Bit Server VM; 1.7.0_45)\",\"sdkIdentifier\":\"1.0\"}\n/v1/9991/services%20bla/convert/amount?aap=noot&mies=geen noot\n"

	actualStart := substr(dataToSign, 0, 22)
	actualEnd := substr(dataToSign, 52, len(dataToSign)-52)

	if expectedStart != actualStart {
		t.Fatalf("TestToDataToSign : expectedStart '%s' != '%s'", expectedStart, actualStart)
	}
	if expectedEnd != actualEnd {
		t.Fatalf("TestToDataToSign : expectedEnd '%s' != '%s'", expectedEnd, actualEnd)
	}
}

func TestCreateAuthenticationSignature(t *testing.T) {
	auth, err := NewDefaultAuthenticator(V1HMAC, "apiKeyId", "secretApiKey")
	if err != nil {
		t.Fatalf("TestCreateAuthenticationSignature : %v", err)
	}

	dataToSign := "DELETE\napplication/json\nFri, 06 Jun 2014 13:39:43 GMT\nx-gcs-clientmetainfo:processed header value\nx-gcs-customerheader:processed header value\nx-gcs-servermetainfo:processed header value\n/v1/9991/tokens/123456789\n"

	authenticationSignature, err := auth.signData(dataToSign)
	if err != nil {
		t.Fatalf("TestCreateAuthenticationSignature : %v", err)
	}
	if authenticationSignature != "VfnXpPBQQoHZivTcAg0JvOWkhnzlPnaCPKpTQn/uMJM=" {
		t.Fatalf("TestCreateAuthenticationSignature : authenticationSignature '%s' != '%s'",
			"VfnXpPBQQoHZivTcAg0JvOWkhnzlPnaCPKpTQn/uMJM=", authenticationSignature)
	}
}

func TestCreateAuthenticationSignature2(t *testing.T) {
	auth, err := NewDefaultAuthenticator(V1HMAC, "EC36A74A98D21", "6Kj5HT0MQKC6D8eb7W3lTg71kVKVDSt1")
	if err != nil {
		t.Fatalf("TestCreateAuthenticationSignature2 : %v", err)
	}

	dataToSign := "GET\n\nFri, 06 Jun 2014 13:39:43 GMT\n/v1/9991/tokens/123456789\n"

	authenticationSignature, err := auth.signData(dataToSign)
	if err != nil {
		t.Fatalf("TestCreateAuthenticationSignature2 : %v", err)
	}
	if authenticationSignature != "9ond5EIN05dBXJGCLRK5om9pxHsyrh/12pZJ7bvmwNM=" {
		t.Fatalf("TestCreateAuthenticationSignature2 : authenticationSignature '%s' != '%s'",
			"9ond5EIN05dBXJGCLRK5om9pxHsyrh/12pZJ7bvmwNM=", authenticationSignature)
	}
}

func TestTotalMinimalExample(t *testing.T) {
	authenticator, err := NewDefaultAuthenticator(V1HMAC, "5e45c937b9db33ae", "I42Zf4pVnRdroHfuHnRiJjJ2B6+22h0yQt/R3nZR8Xg=")
	if err != nil {
		t.Fatalf("TestTotalMinimalExample : %v", err)
	}

	httpHeaders := []communication.Header{}

	userAgentHeader, err := communication.NewHeader("User-Agent", "Apache-HttpClient/4.3.4 (java 1.5)")
	if err != nil {
		t.Fatalf("TestTotalMinimalExample : %v", err)
	}
	httpHeaders = append(httpHeaders, *userAgentHeader)

	dateHeader, err := communication.NewHeader("Date", "Fri, 06 Jun 2014 13:39:43 GMT")
	if err != nil {
		t.Fatalf("TestTotalMinimalExample : %v", err)
	}
	httpHeaders = append(httpHeaders, *dateHeader)
	parsedURL, err := url.Parse("http://world.api-ingenico.com:8080/v1/9991/tokens/123456789")
	if err != nil {
		t.Fatalf("TestTotalMinimalExample : %v", err)
	}

	signature, err := authenticator.CreateSimpleAuthenticationSignature(http.MethodGet, *parsedURL, httpHeaders)
	if err != nil {
		t.Fatalf("TestTotalMinimalExample : %v", err)
	}
	if signature != "GCS v1HMAC:5e45c937b9db33ae:J5LjfSBvrQNhu7gG0gvifZt+IWNDReGCmHmBmth6ueI=" {
		t.Fatalf("TestTotalMinimalExample : signature '%s' != '%s'",
			"GCS v1HMAC:5e45c937b9db33ae:J5LjfSBvrQNhu7gG0gvifZt+IWNDReGCmHmBmth6ueI=", signature)
	}
}

func TestTotalFullExample(t *testing.T) {
	authenticator, err := NewDefaultAuthenticator(V1HMAC, "5e45c937b9db33ae", "I42Zf4pVnRdroHfuHnRiJjJ2B6+22h0yQt/R3nZR8Xg=")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}

	httpHeaders := []communication.Header{}

	userAgentHeader, err := communication.NewHeader("User-Agent", "Apache-HttpClient/4.3.4 (java 1.5)")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}
	httpHeaders = append(httpHeaders, *userAgentHeader)

	serverMetaInfoHeader, err := communication.NewHeader("X-GCS-ServerMetaInfo", "processed header value")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}
	httpHeaders = append(httpHeaders, *serverMetaInfoHeader)

	clientMetaInfoHeader, err := communication.NewHeader("X-GCS-ClientMetaInfo", "processed header value")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}
	httpHeaders = append(httpHeaders, *clientMetaInfoHeader)

	contentTypeHeader, err := communication.NewHeader("Content-Type", "application/json")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}
	httpHeaders = append(httpHeaders, *contentTypeHeader)

	customerHeader, err := communication.NewHeader("X-GCS-CustomerHeader", "processed header value")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}
	httpHeaders = append(httpHeaders, *customerHeader)

	dateHeader, err := communication.NewHeader("Date", "Fri, 06 Jun 2014 13:39:43 GMT")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}
	httpHeaders = append(httpHeaders, *dateHeader)

	parsedURL, err := url.Parse("http://world.api-ingenico.com:8080/v1/9991/tokens/123456789")
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}

	signature, err := authenticator.CreateSimpleAuthenticationSignature(http.MethodDelete, *parsedURL, httpHeaders)
	if err != nil {
		t.Fatalf("TestTotalFullExample : %v", err)
	}
	if signature != "GCS v1HMAC:5e45c937b9db33ae:jGWLz3ouN4klE+SkqO5gO+KkbQNM06Rric7E3dcfmqw=" {
		t.Fatalf("TestTotalFullExample : signature '%s' != '%s'",
			"GCS v1HMAC:5e45c937b9db33ae:jGWLz3ouN4klE+SkqO5gO+KkbQNM06Rric7E3dcfmqw=", signature)
	}
}
