package defaultimpl

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/url"
	"regexp"
	"sort"
	"strings"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
)

// AuthorizationType represents the type of authentication that is used. Here only v1HMAC
type AuthorizationType string

const (
	// V1HMAC represent the authorization type for the HMAC algorithm, version 1
	V1HMAC = "v1HMAC"
)

//DefaultAuthenticator represents the default implementation of an authenticator
type DefaultAuthenticator struct {
	authType     AuthorizationType
	apiKeyID     string
	secretAPIKey string
}

// NewDefaultAuthenticator creates a DefaultAuthenticator with the given authType, apiKeyID and secretAPIKey
//- Based on the Authenticationtype value both the Ingenico ePayments platform and the merchant know which security implementation is used.
//  A version number is used for backward compatibility in the future.
//- The apiKeyID is an identifier for the secret API key. The apiKeyID can be retrieved from the Configuration Center
//  This identifier is visible in the HTTP request and is also used to identify the correct account.
//- secretAPIKey is a shared secret. The shared secret can be retrieved from the Configuration Center.
//  An apiKeyID and secretAPIKey always go hand-in-hand, the difference is that secretAPIKey is never visible in the HTTP request.
//  This secret is used as input for the HMAC algorithm.
func NewDefaultAuthenticator(authType AuthorizationType, apiKeyID string, secretAPIKey string) (*DefaultAuthenticator, error) {
	if strings.TrimSpace(apiKeyID) == "" {
		return nil, errors.New("apikeyid is required")
	}
	if strings.TrimSpace(secretAPIKey) == "" {
		return nil, errors.New("secretAPIKey is required")
	}
	defaultAuthenticator := DefaultAuthenticator{authType: authType, apiKeyID: apiKeyID, secretAPIKey: secretAPIKey}
	return &defaultAuthenticator, nil
}

// CreateSimpleAuthenticationSignature creates an authentication signature for the given httpMethod, resourceURI and requestHeaders
func (a DefaultAuthenticator) CreateSimpleAuthenticationSignature(httpMethod string, resourceURI url.URL, requestHeaders []communication.Header) (output string, err error) {
	if strings.TrimSpace(httpMethod) == "" {
		err = errors.New("httpMethod is required")
		return
	}
	dataToSign, err := toDataToSign(httpMethod, resourceURI, requestHeaders)
	if err != nil {
		return
	}
	signedData, err := a.signData(dataToSign)
	if err != nil {
		return
	}
	output = "GCS " + string(a.authType) + ":" + a.apiKeyID + ":" + signedData
	return
}

func toDataToSign(httpMethod string, resourceURI url.URL, requestHeaders []communication.Header) (output string, err error) {
	xgcsHTTPHeaders := communication.Headers{}
	var contentType, date string
	for _, header := range requestHeaders {
		if strings.ToLower(header.Name()) == "content-type" {
			contentType = header.Value()
		} else if strings.ToLower(header.Name()) == "date" {
			date = header.Value()
		} else if strings.HasPrefix(strings.ToLower(header.Name()), "x-gcs") {
			var newHeader *communication.Header
			var newValue string
			newValue, err = toCanonicalizedHeaderValue(header.Value())
			if err != nil {
				return
			}
			newHeader, err = communication.NewHeader(strings.ToLower(header.Name()), newValue)
			if err != nil {
				return
			}

			xgcsHTTPHeaders = append(xgcsHTTPHeaders, *newHeader)
		}
	}
	sort.Sort(xgcsHTTPHeaders)

	dataToSign := bytes.Buffer{}
	dataToSign.WriteString(httpMethod)
	dataToSign.WriteRune('\n')
	dataToSign.WriteString(contentType)
	dataToSign.WriteRune('\n')
	dataToSign.WriteString(date)
	dataToSign.WriteRune('\n')
	for _, header := range xgcsHTTPHeaders {
		dataToSign.WriteString(header.Name() + ":" + header.Value())
		dataToSign.WriteRune('\n')
	}
	var canonizalizedResource string
	canonizalizedResource, err = toCanonicalizedResource(resourceURI)
	dataToSign.WriteString(canonizalizedResource)
	dataToSign.WriteRune('\n')

	output = dataToSign.String()
	return
}

func toCanonicalizedResource(resourceURI url.URL) (string, error) {
	resource := bytes.Buffer{}

	resource.WriteString(resourceURI.EscapedPath())
	if resourceURI.RawQuery != "" {
		query, err := url.QueryUnescape(resourceURI.RawQuery)
		if err != nil {
			return "", err
		}

		resource.WriteString("?" + query)
	}

	return resource.String(), nil
}

func toCanonicalizedHeaderValue(s string) (string, error) {
	regex, err := regexp.Compile("\r?\n[\t\f ]*")
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(regex.ReplaceAllString(s, " ")), nil
}

func (a *DefaultAuthenticator) signData(s string) (string, error) {
	mac := hmac.New(sha256.New, []byte(a.secretAPIKey))
	mac.Write([]byte(s))
	hmacOutput := mac.Sum(nil)

	writableBuffer := bytes.Buffer{}
	encoder := base64.NewEncoder(base64.StdEncoding, &writableBuffer)
	encoder.Write(hmacOutput)
	err := encoder.Close()
	if err != nil {
		return "", err
	}
	output := writableBuffer.String()
	return output, nil
}
