package communicator

import (
	"errors"
	"net/url"
)

var (
	// ErrNilEndpoint occurs when the given endpoint is nil
	ErrNilEndpoint = errors.New("apiEndpoint is required")
	// ErrPathEndpoint occurs when the endpoint contains a path
	ErrPathEndpoint = errors.New("apiEndpoint should not contain a path")
	// ErrUserInfo occurs when the endpoint contains any userinfo
	ErrUserInfo = errors.New("apiEndpoint should not contain user info, query or fragment")
	// ErrNilConnection occurs when a nil connection is supplied
	ErrNilConnection = errors.New("connection is required")
	// ErrNilAuthenticator occurs when a nil authenticator is supplied
	ErrNilAuthenticator = errors.New("authenticator is required")
	// ErrNilMetaDataProvider occurs when a nil metaDataProvider is supplied
	ErrNilMetaDataProvider = errors.New("metaDataProvider is required")
)

// Session Contains the components needed to communicate with the GlobalCollect platform. Thread-safe.
type Session struct {
	apiEndpoint      *url.URL
	connection       Connection
	metaDataProvider *MetaDataProvider
	authenticator    Authenticator
}

// NewSession creates a new session with the given apiEndpoint, connection and authenticator
func NewSession(apiEndpoint *url.URL, connection Connection, authenticator Authenticator,
	metaDataProvider *MetaDataProvider) (*Session, error) {
	if apiEndpoint == nil {
		return nil, ErrNilEndpoint
	}
	if len(apiEndpoint.Path) > 0 && apiEndpoint.Path != "/" {
		return nil, ErrPathEndpoint
	}
	if apiEndpoint.User != nil || apiEndpoint.RawQuery != "" || apiEndpoint.Fragment != "" {
		return nil, ErrUserInfo
	}
	if connection == nil {
		return nil, ErrNilConnection
	}
	if authenticator == nil {
		return nil, ErrNilAuthenticator
	}
	if metaDataProvider == nil {
		return nil, ErrNilMetaDataProvider
	}
	return &Session{apiEndpoint: apiEndpoint,
		connection:       connection,
		metaDataProvider: metaDataProvider,
		authenticator:    authenticator}, nil
}

// APIEndpoint returns the apiEndpoint of the session
func (s *Session) APIEndpoint() *url.URL {
	return s.apiEndpoint
}

// Connection returns the connection of the session
func (s *Session) Connection() Connection {
	return s.connection
}

// MetaDataProvider returns the metaDataProvider of the session
func (s *Session) MetaDataProvider() *MetaDataProvider {
	return s.metaDataProvider
}

// Authenticator returns the authenticator of the session
func (s *Session) Authenticator() Authenticator {
	return s.authenticator
}
