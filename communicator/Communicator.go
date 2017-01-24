package communicator

import (
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/logging"
)

var (
	// ErrNilSession occurs when the given session is nil
	ErrNilSession = errors.New("session is nil")
	// ErrNilMarshaller occurs when the given marshaller is nil
	ErrNilMarshaller = errors.New("marshaller is nil")
)

// A Communicator is used to communicate with the GlobalCollect platform web services.
// It contains all the logic to transform a request object to a HTTP request and a HTTP response to a response object.
// It is also thread safe.
type Communicator struct {
	session    *Session
	marshaller Marshaller
}

// Session returns the session of this Communicator
func (c *Communicator) Session() *Session {
	return c.session
}

// Marshaller returns the marshaller of this Communicator
func (c *Communicator) Marshaller() Marshaller {
	return c.marshaller
}

// Close closes the connection of the Communicator
func (c *Communicator) Close() error {
	return c.session.Connection().Close()
}

// EnableLogging turns on logging using the given communicator logger.
func (c *Communicator) EnableLogging(logger logging.CommunicatorLogger) {
	c.session.Connection().EnableLogging(logger)
}

// DisableLogging turns off logging.
func (c *Communicator) DisableLogging() {
	c.session.Connection().DisableLogging()
}

// Get corresponds to the HTTP Get method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set to nil
// context is an optional Call context which can be used
// expectedObject is a reference to the expected response object
//
// Possibly returns an error
func (c *Communicator) Get(relativePath string, requestHeaders []communication.Header, requestParameters ParamRequest, context communication.CallContext, expectedObject interface{}) error {
	connection := c.session.Connection()
	var requestParameterList []RequestParam
	var err error
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}
	if requestHeaders == nil {
		requestHeaders = []communication.Header{}
	}
	requestHeaders, err = c.addGenericHeaders(http.MethodGet, uri, requestHeaders, context)
	if err != nil {
		return err
	}
	response, err := connection.Get(uri, requestHeaders)
	if err != nil {
		return err
	}
	return c.processResponse(response, relativePath, context, expectedObject)
}

// Delete corresponds to the HTTP Delete method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set to nil
// context is an optional Call context which can be used
// expectedObject is a reference to the expected response object
//
// Possibly returns an error
func (c *Communicator) Delete(relativePath string, requestHeaders []communication.Header, requestParameters ParamRequest, context communication.CallContext, expectedObject interface{}) error {
	connection := c.session.Connection()
	var requestParameterList []RequestParam
	var err error
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}
	if requestHeaders == nil {
		requestHeaders = []communication.Header{}
	}
	requestHeaders, err = c.addGenericHeaders(http.MethodDelete, uri, requestHeaders, context)
	if err != nil {
		return err
	}
	response, err := connection.Delete(uri, requestHeaders)
	if err != nil {
		return err
	}
	return c.processResponse(response, relativePath, context, expectedObject)
}

// Post corresponds to the HTTP Post method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set it to nil
// requestBody is the body of the request. If not used set to nil
// context is an optional Call context which can be used. If not used set it to nil
// expectedObject is a reference to the expected response object
//
// Possibly returns an error
func (c *Communicator) Post(relativePath string, requestHeaders []communication.Header, requestParameters ParamRequest, requestBody interface{}, context communication.CallContext, expectedResponse interface{}) error {
	connection := c.session.Connection()
	var requestParameterList []RequestParam
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}
	if requestHeaders == nil {
		requestHeaders = []communication.Header{}
	}

	var requestJSON string
	if requestBody != nil {
		jsonHeader, _ := communication.NewHeader("Content-Type", "application/json")
		requestHeaders = append(requestHeaders, *jsonHeader)

		requestJSON, err = c.marshaller.Marshal(requestBody)

		if err != nil {
			return err
		}
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodPost, uri, requestHeaders, context)
	if err != nil {
		return err
	}
	response, err := connection.Post(uri, requestHeaders, requestJSON)
	if err != nil {
		return err
	}
	return c.processResponse(response, relativePath, context, expectedResponse)
}

// Put corresponds to the HTTP Put method
//
// relativePath is the Path to call, relative to the base URI
// requestHeaders is an optimal list of request headers
// requestParameters is an optional set of request parameters. If not used set it to nil
// requestBody is the body of the request. If not used set to nil
// context is an optional Call context which can be used. If not used set it to nil
// expectedObject is a reference to the expected response object
//
// Possibly returns an error
func (c *Communicator) Put(relativePath string, requestHeaders []communication.Header, requestParameters ParamRequest, requestBody interface{}, context communication.CallContext, expectedObject interface{}) error {
	connection := c.session.Connection()
	var requestParameterList []RequestParam
	var err error
	if requestParameters != nil {
		requestParameterList = requestParameters.ToRequestParameters()
	}
	uri, err := c.toAbsoluteURI(relativePath, requestParameterList)
	if err != nil {
		return err
	}
	if requestHeaders == nil {
		requestHeaders = []communication.Header{}
	}

	var requestJSON string
	if requestBody != nil {
		jsonHeader, _ := communication.NewHeader("Content-Type", "application/json")
		requestHeaders = append(requestHeaders, *jsonHeader)
		requestJSON, err = c.marshaller.Marshal(requestBody)

		if err != nil {
			return err
		}
	}

	requestHeaders, err = c.addGenericHeaders(http.MethodPut, uri, requestHeaders, context)
	if err != nil {
		return err
	}
	response, err := connection.Put(uri, requestHeaders, requestJSON)
	if err != nil {
		return err
	}
	return c.processResponse(response, relativePath, context, expectedObject)
}

// CloseExpiredConnections is a utility method that delegates the call to this communicator's session's connection.
// Also see Connection.CloseExpiredConnections
func (c *Communicator) CloseExpiredConnections() {
	c.session.Connection().CloseExpiredConnections()
}

// CloseIdleConnections is a utility method that delegates the call to this communicator's session's connection.
// The duration argument is a specification of how long the connection has to be Idle.
// Also see Connection.CloseIdleConnections
func (c *Communicator) CloseIdleConnections(duration time.Duration) {
	c.session.Connection().CloseIdleConnections(duration)
}

func (c *Communicator) toAbsoluteURI(relativePath string, requestParameters []RequestParam) (url.URL, error) {
	apiEndpoint := c.session.APIEndpoint()

	if apiEndpoint.Path != "" {
		return url.URL{}, ErrPathEndpoint
	}
	if apiEndpoint.User != nil || apiEndpoint.Query().Encode() != "" || apiEndpoint.Fragment != "" {
		return url.URL{}, ErrUserInfo
	}

	var absolutePath string
	if strings.Index(relativePath, "/") == 0 {
		absolutePath = relativePath
	} else {
		absolutePath = "/" + relativePath
	}

	var rawQuery = ""
	for index, element := range requestParameters {
		if index != 0 {
			rawQuery += "&"
		}
		id := url.QueryEscape(element.Name())
		value := url.QueryEscape(element.Value())
		rawQuery += id + "=" + value
	}

	var url = url.URL{Scheme: apiEndpoint.Scheme,
		Host:     apiEndpoint.Host,
		Path:     absolutePath,
		RawQuery: rawQuery}

	return url, nil
}

// Adds the necessary headers to the given list of headers. This includes the authorization header,
// which uses other headers, so when you need to override this method,
// make sure to call AddGenericHeaders at the end of your overridden method.
func (c *Communicator) addGenericHeaders(httpMethod string, url url.URL, requestHeaders []communication.Header, context communication.CallContext) ([]communication.Header, error) {
	//add server meta info headers
	requestHeaders = append(requestHeaders, c.session.MetaDataProvider().MetaDataHeaders()...)

	//add date header
	header, err := communication.NewHeader("Date", getHeaderDateString())
	if err != nil {
		return nil, err
	}
	requestHeaders = append(requestHeaders, *header)

	//add content specific headers
	if context != nil && context.GetIdempotenceKey() != "" {
		header, err = communication.NewHeader("X-GCS-Idempotence-Key", context.GetIdempotenceKey())
		if err != nil {
			return nil, err
		}
		requestHeaders = append(requestHeaders, *header)
	}

	//add signature
	authenticator := c.session.Authenticator()
	authenticationSignature, err := authenticator.CreateSimpleAuthenticationSignature(httpMethod, url, requestHeaders)
	if err != nil {
		return nil, err
	}

	header, err = communication.NewHeader("Authorization", authenticationSignature)
	if err != nil {
		return nil, err
	}
	requestHeaders = append(requestHeaders, *header)
	return requestHeaders, nil
}

// Gets the date in the preferred format for the HTTP date header (RFC1123).
func getHeaderDateString() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

// Checks the Response for errors and creates an error if necessary.
func (c *Communicator) createErrorIfNecessary(response *communication.Response, requestPath string) error {
	var statusCode = response.StatusCode()
	if statusCode < http.StatusOK || statusCode >= http.StatusMultipleChoices {
		var body = response.Body()
		if body != "" && !isJSON(response) {
			var cause = sdkErrors.NewResponseError(*response)
			if statusCode == http.StatusNotFound {
				err1, err2 := sdkErrors.NewNotFoundErrorVerbose("The requested resource was not found; invalid path: " + requestPath, cause)
				if err2 != nil {
					return err2
				}
				return err1
			}
			err1, err2 := sdkErrors.NewCommunicationError(cause)
			if err2 != nil {
				return err2
			}
			return err1
		}
		return sdkErrors.NewResponseError(*response)
	}
	return nil
}

func (c *Communicator) processResponse(response *communication.Response, requestPath string, context communication.CallContext, expectedObject interface{}) error {
	if context != nil {
		timestamp, err := getIdempotenceTimestamp(response)
		if err != nil {
			return err
		}
		context.SetIdempotenceRequestTimestamp(timestamp)
	}

	err := c.createErrorIfNecessary(response, requestPath)
	if err != nil {
		return err
	}

	return c.marshaller.Unmarshal(response.Body(), expectedObject)
}

func getIdempotenceTimestamp(response *communication.Response) (*int64, error) {
	header := response.GetHeader("X-GCS-Idempotence-Request-Timestamp")
	if header == nil {
		return nil, nil
	}

	retValue, err := strconv.ParseInt(header.Value(), 10, 64)
	return &retValue, err
}

func isJSON(response *communication.Response) bool {
	header := response.GetHeader("Content-Type")
	if header == nil {
		return true
	}

	contentType := strings.ToLower(header.Value())

	if contentType == "application/json" {
		return true
	}

	return strings.HasPrefix(contentType, "application/json")
}

// NewCommunicator creates a communicator with the given session and marshaller
func NewCommunicator(session *Session, marshaller Marshaller) (*Communicator, error) {
	if session == nil {
		return nil, ErrNilSession
	}
	if marshaller == nil {
		return nil, ErrNilMarshaller
	}

	return &Communicator{session, marshaller}, nil
}
