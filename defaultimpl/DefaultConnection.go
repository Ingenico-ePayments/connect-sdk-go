package defaultimpl

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	"github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/logging"
)

// DefaultConnection is the default implementation for the connection interface. Supports Pooling, and is thread safe.
type DefaultConnection struct {
	client              http.Client
	underlyingTransport *http.Transport
	logger              logging.CommunicatorLogger
	proxyAuth           string
}

func (c *DefaultConnection) logRequest(id, body string, req *http.Request) error {
	if c.logger == nil {
		return nil
	}

	var url url.URL
	if req.URL != nil {
		url = *req.URL
	}

	reqMessage, err := logging.NewRequestLogMessageBuilder(id, req.Method, url)
	if err != nil {
		c.logError(id, err)
		return err
	}

	for k, v := range req.Header {
		for _, rv := range v {
			reqMessage.AddHeader(k, rv)
		}
	}

	reqMessage.SetBody(body, req.Header.Get("Content-Type"))

	message, err := reqMessage.BuildMessage()
	if err != nil {
		c.logError(id, err)
		return err
	}

	c.logger.LogRequestLogMessage(message)

	return nil
}

func (c *DefaultConnection) logResponse(id, body string, resp *http.Response, duration time.Duration) error {
	if c.logger == nil {
		return nil
	}

	respMessage, err := logging.NewResponseLogMessageBuilder(id, resp.StatusCode, duration)
	if err != nil {
		c.logError(id, err)
		return err
	}

	for k, v := range resp.Header {
		for _, rv := range v {
			respMessage.AddHeader(k, rv)
		}
	}

	respMessage.SetBody(body, resp.Header.Get("Content-Type"))

	message, err := respMessage.BuildMessage()
	if err != nil {
		c.logError(id, err)
		return err
	}

	c.logger.LogResponseLogMessage(message)

	return nil
}

func (c *DefaultConnection) logError(id string, err error) {
	if c.logger != nil {
		c.logger.LogError(id, err)
	}
}

// Close implements the io.Closer interface
func (c *DefaultConnection) Close() error {
	// No-op, because the http.Client connection's close automatically after the socket timeout passes
	// and they can't be closed manually
	return nil
}

// NewDefaultConnection creates a new object that implements Connection, and initializes it
func NewDefaultConnection(socketTimeout, connectTimeout, keepAliveTimeout, idleTimeout time.Duration, maxConnections int, proxy *url.URL) (*DefaultConnection, error) {
	dialer := net.Dialer{
		Timeout:   connectTimeout,
		KeepAlive: keepAliveTimeout,
	}

	transport := &http.Transport{
		DialContext:     dialer.DialContext,
		Proxy:           http.ProxyURL(proxy),
		IdleConnTimeout: idleTimeout,
		MaxIdleConns:    maxConnections,
	}

	client := http.Client{
		Transport: transport,
		Timeout:   socketTimeout,
	}

	proxyAuth := getProxyAuth(proxy)

	return &DefaultConnection{client, transport, nil, proxyAuth}, nil
}

// Get sends a GET request to the GlobalCollect platform and return the response.
func (c *DefaultConnection) Get(uri url.URL, headerList []communication.Header) (*communication.Response, error) {
	r := communication.NewRequest("GET", uri, headerList)
	return c.sendRequest(r, "")
}

// Delete sends a DELETE request to the GlobalCollect platform and return the response.
func (c *DefaultConnection) Delete(uri url.URL, headerList []communication.Header) (*communication.Response, error) {
	r := communication.NewRequest("DELETE", uri, headerList)
	return c.sendRequest(r, "")
}

// Put sends a PUT request to the GlobalCollect platform and return the response.
func (c *DefaultConnection) Put(uri url.URL, headerList []communication.Header, body string) (*communication.Response, error) {
	r := communication.NewRequest("PUT", uri, headerList)
	r.SetBodyString(body)
	return c.sendRequest(r, body)
}

// Post sends a POST request to the GlobalCollect platform and return the response.
func (c *DefaultConnection) Post(uri url.URL, headerList []communication.Header, body string) (*communication.Response, error) {
	r := communication.NewRequest("POST", uri, headerList)
	r.SetBodyString(body)
	return c.sendRequest(r, body)
}

func pseudoUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func getProxyAuth(proxy *url.URL) (string) {
	if proxy == nil || proxy.User == nil {
		return ""
	}

	return "Basic " + base64.StdEncoding.EncodeToString([]byte(proxy.User.String()))
}

func (c *DefaultConnection) sendRequest(r *communication.Request, body string) (*communication.Response, error) {
	id, err := pseudoUUID()
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(r.Method, r.Destination.String(), r.Body())
	if err != nil {
		c.logError(id, err)
		return nil, err
	}

	for _, h := range r.Headers {
		httpRequest.Header[h.Name()] = append(httpRequest.Header[h.Name()], h.Value())
	}
	if len(c.proxyAuth) > 0 {
		httpRequest.Header["Proxy-Authorization"] = append(httpRequest.Header["Proxy-Authorization"], c.proxyAuth)
	}

	start := time.Now()

	c.logRequest(id, body, httpRequest)

	resp, err := c.client.Do(httpRequest)
	switch ce := err.(type) {
	case *url.Error:
		{
			c.logError(id, ce)

			newErr, _ := errors.NewCommunicationError(ce)
			return nil, newErr
		}
	}
	if err != nil {
		c.logError(id, err)
		return nil, err
	}

	end := time.Now()

	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.logError(id, err)
		return nil, err
	}

	err = resp.Body.Close()
	if err != nil {
		c.logError(id, err)
		return nil, err
	}

	c.logResponse(id, string(bodyBuf), resp, end.Sub(start))

	respHeaders := []communication.Header{}
	for name, values := range resp.Header {
		if name == "X-Gcs-Idempotence-Request-Timestamp" {
			name = "X-GCS-Idempotence-Request-Timestamp"
		}

		header, err := communication.NewHeader(name, values[0])
		if err != nil {
			c.logError(id, err)
			return nil, err
		}
		respHeaders = append(respHeaders, *header)
	}
	retVal, err := communication.NewResponse(resp.StatusCode, string(bodyBuf), respHeaders)
	if err != nil {
		return nil, err
	}
	return retVal, nil
}

// CloseIdleConnections closes all HTTP connections that have been idle for the specified time. This should also include
// all expired HTTP connections.
// timespan represents the time spent idle
// Note: in the current implementation, it is only possible to close the connection after a predetermined time
// Therefore, this implementation ignores the parameter, and instead uses the preconfigured one
func (c *DefaultConnection) CloseIdleConnections(t time.Duration) {
	// Assume t is equal to configured value
	c.underlyingTransport.CloseIdleConnections()
}

// CloseExpiredConnections closes all expired HTTP connections.
func (c *DefaultConnection) CloseExpiredConnections() {
	// No-op, because this is done automatically for this implementation
}

// EnableLogging implements the logging.Capable interface
// Enables logging to the given CommunicatorLogger
func (c *DefaultConnection) EnableLogging(l logging.CommunicatorLogger) {
	c.logger = l
}

// DisableLogging implements the logging.Capable interface
// Disables logging
func (c *DefaultConnection) DisableLogging() {
	c.logger = nil
}
