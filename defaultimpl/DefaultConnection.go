package defaultimpl

import (
	"bytes"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/textproto"
	"net/url"
	"strings"
	"time"

	"github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication"
	sdkErrors "github.com/Ingenico-ePayments/connect-sdk-go/errors"
	"github.com/Ingenico-ePayments/connect-sdk-go/logging"
	"github.com/Ingenico-ePayments/connect-sdk-go/logging/obfuscation"
)

// DefaultConnection is the default implementation for the connection interface. Supports Pooling, and is thread safe.
type DefaultConnection struct {
	client              http.Client
	underlyingTransport *http.Transport
	logger              logging.CommunicatorLogger
	proxyAuth           string
	bodyObfuscator      obfuscation.BodyObfuscator
	headerObfuscator    obfuscation.HeaderObfuscator
}

func (c *DefaultConnection) logRequest(id, body string, req *http.Request) error {
	if c.logger == nil {
		return nil
	}

	var url url.URL
	if req.URL != nil {
		url = *req.URL
	}

	reqMessage, err := logging.NewRequestLogMessageBuilderWithObfuscators(id, req.Method, url, c.bodyObfuscator, c.headerObfuscator)
	if err != nil {
		c.logError(id, err)
		return err
	}

	for k, v := range req.Header {
		for _, rv := range v {
			reqMessage.AddHeader(k, rv) // #nosec G104
		}
	}

	reqMessage.SetBody(body, req.Header.Get("Content-Type")) // #nosec G104

	message, err := reqMessage.BuildMessage()
	if err != nil {
		c.logError(id, err)
		return err
	}

	c.logger.LogRequestLogMessage(message)

	return nil
}

func (c *DefaultConnection) logResponse(id string, reader io.Reader, binaryResponse bool, resp *http.Response, duration time.Duration) error {
	if c.logger == nil {
		return nil
	}

	respMessage, err := logging.NewResponseLogMessageBuilderWithObfuscators(id, resp.StatusCode, duration, c.bodyObfuscator, c.headerObfuscator)
	if err != nil {
		c.logError(id, err)
		return err
	}

	for k, v := range resp.Header {
		for _, rv := range v {
			respMessage.AddHeader(k, rv) // #nosec G104
		}
	}

	if binaryResponse {
		respMessage.SetBinaryBody(resp.Header.Get("Content-Type")) // #nosec G104
	} else {
		bodyBuff, err := ioutil.ReadAll(reader)
		if err != nil {
			c.logError(id, err)
			return err
		}

		respMessage.SetBody(string(bodyBuff), resp.Header.Get("Content-Type")) // #nosec G104
	}

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
		TLSClientConfig: &tls.Config{
			MinVersion: tls.VersionTLS12,
		},
	}

	client := http.Client{
		Transport: transport,
		Timeout:   socketTimeout,
	}

	proxyAuth := getProxyAuth(proxy)

	return &DefaultConnection{client, transport, nil, proxyAuth, obfuscation.DefaultBodyObfuscator(), obfuscation.DefaultHeaderObfuscator()}, nil
}

// Get sends a GET request to the Ingenico ePayments platform and calls the given response handler with the response.
func (c *DefaultConnection) Get(uri url.URL, headerList []communication.Header, handler communication.ResponseHandler) (interface{}, error) {
	return c.sendRequest("GET", uri, headerList, nil, "", handler)
}

// Delete sends a DELETE request to the Ingenico ePayments platform and calls the given response handler with the response.
func (c *DefaultConnection) Delete(uri url.URL, headerList []communication.Header, handler communication.ResponseHandler) (interface{}, error) {
	return c.sendRequest("DELETE", uri, headerList, nil, "", handler)
}

// Post sends a POST request to the Ingenico ePayments platform and calls the given response handler with the response.
func (c *DefaultConnection) Post(uri url.URL, headerList []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error) {
	return c.sendRequest("POST", uri, headerList, strings.NewReader(body), body, handler)
}

// PostMultipart sends a multipart/form-data POST request to the Ingenico ePayments platform and calls the given response handler with the response.
func (c *DefaultConnection) PostMultipart(uri url.URL, headerList []communication.Header, body *communication.MultipartFormDataObject, handler communication.ResponseHandler) (interface{}, error) {
	r, err := c.createMultipartReader(body)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return c.sendRequest("POST", uri, headerList, r, "<binary content>", handler)
}

// Put sends a PUT request to the Ingenico ePayments platform and returns the response.
func (c *DefaultConnection) Put(uri url.URL, headerList []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error) {
	return c.sendRequest("PUT", uri, headerList, strings.NewReader(body), body, handler)
}

// PutMultipart sends a multipart/form-data POST request to the Ingenico ePayments platform and calls the given response handler with the response.
func (c *DefaultConnection) PutMultipart(uri url.URL, headerList []communication.Header, body *communication.MultipartFormDataObject, handler communication.ResponseHandler) (interface{}, error) {
	r, err := c.createMultipartReader(body)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return c.sendRequest("PUT", uri, headerList, r, "<binary content>", handler)
}

func pseudoUUID() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func getProxyAuth(proxy *url.URL) string {
	if proxy == nil || proxy.User == nil {
		return ""
	}

	return "Basic " + base64.StdEncoding.EncodeToString([]byte(proxy.User.String()))
}

func isBinaryContent(headers []communication.Header) bool {
	header := communication.Headers(headers).GetHeader("Content-Type")
	if header == nil {
		return false
	}

	contentType := strings.ToLower(header.Value())

	return !strings.HasPrefix(contentType, "text/") && !strings.Contains(contentType, "json") && !strings.Contains(contentType, "xml")
}

func (c *DefaultConnection) sendRequest(method string, uri url.URL, headerList []communication.Header, body io.Reader, bodyString string, handler communication.ResponseHandler) (interface{}, error) {
	id, err := pseudoUUID()
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		c.logError(id, err)
		return nil, err
	}

	for _, h := range headerList {
		httpRequest.Header[h.Name()] = append(httpRequest.Header[h.Name()], h.Value())
	}
	if len(c.proxyAuth) > 0 {
		httpRequest.Header["Proxy-Authorization"] = append(httpRequest.Header["Proxy-Authorization"], c.proxyAuth)
	}

	start := time.Now()

	c.logRequest(id, bodyString, httpRequest) // #nosec G104

	resp, err := c.client.Do(httpRequest)
	switch ce := err.(type) {
	case *url.Error:
		{
			c.logError(id, ce)

			newErr, _ := sdkErrors.NewCommunicationError(ce)
			return nil, newErr
		}
	}
	if err != nil {
		c.logError(id, err)
		return nil, err
	}

	end := time.Now()

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			c.logError(id, err)
		}
	}()

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

	bodyReader := resp.Body.(io.Reader)
	if isBinaryContent(respHeaders) {
		c.logResponse(id, nil, true, resp, end.Sub(start)) // #nosec G104
	} else {
		readBuffer := bytes.NewBuffer([]byte{})
		teeReader := io.TeeReader(resp.Body, readBuffer)
		bodyReader = teeReader
		defer c.logResponse(id, io.MultiReader(readBuffer, teeReader), false, resp, end.Sub(start))
	}

	return handler.Handle(resp.StatusCode, respHeaders, bodyReader)
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

// SetBodyObfuscator sets the body obfuscator to use.
func (c *DefaultConnection) SetBodyObfuscator(bodyObfuscator obfuscation.BodyObfuscator) {
	c.bodyObfuscator = bodyObfuscator
}

// SetHeaderObfuscator sets the header obfuscator to use.
func (c *DefaultConnection) SetHeaderObfuscator(headerObfuscator obfuscation.HeaderObfuscator) {
	c.headerObfuscator = headerObfuscator
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

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func (c *DefaultConnection) createMultipartReader(body *communication.MultipartFormDataObject) (io.ReadCloser, error) {
	r, w := io.Pipe()

	writer := multipart.NewWriter(w)
	err := writer.SetBoundary(body.GetBoundary())
	if err != nil {
		return nil, err
	}
	if writer.FormDataContentType() != body.GetContentType() {
		return nil, errors.New("multipart.Writer  did not create the expected content type")
	}

	go func() {
		for name, value := range body.GetValues() {
			err := writer.WriteField(name, value)
			if err != nil {
				w.CloseWithError(err)
				return
			}
		}
		for name, file := range body.GetFiles() {
			// Do not use writer.CreateFormFile because it does not allow a custom content type
			header := textproto.MIMEHeader{}
			header.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(name), escapeQuotes(file.GetFileName())))
			header.Set("Content-Type", file.GetContentType())
			pw, err := writer.CreatePart(header)
			if err != nil {
				w.CloseWithError(err)
				return
			}
			_, err = io.Copy(pw, file.GetContent())
			if err != nil {
				w.CloseWithError(err)
				return
			}
		}
		err = writer.Close()
		if err != nil {
			w.CloseWithError(err)
			return
		}
		w.Close() // #nosec G104
	}()

	return r, nil
}
