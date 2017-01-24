package communicator

import "net/url"

// SessionBuilder is the builder for Session objects
type SessionBuilder struct {
	APIEndpoint      *url.URL
	Connection       Connection
	MetaDataProvider *MetaDataProvider
	Authenticator    Authenticator
}

// WithAPIEndpoint sets the GlobalCollect platform API endpoint to be used by the Session
func (s *SessionBuilder) WithAPIEndpoint(endpoint *url.URL) *SessionBuilder {
	s.APIEndpoint = endpoint

	return s
}

// WithConnection sets the Connection to be used by the Session
func (s *SessionBuilder) WithConnection(connection Connection) *SessionBuilder {
	s.Connection = connection

	return s
}

// WithMetaDataProvider sets the MetaDataProvider to be used by the Session
func (s *SessionBuilder) WithMetaDataProvider(provider *MetaDataProvider) *SessionBuilder {
	s.MetaDataProvider = provider

	return s
}

// WithAuthenticator sets the Authenticator to be used by the Session
func (s *SessionBuilder) WithAuthenticator(auth Authenticator) *SessionBuilder {
	s.Authenticator = auth

	return s
}

// Build creates a Session object based on the builder parameters
func (s *SessionBuilder) Build() (*Session, error) {
	return NewSession(s.APIEndpoint, s.Connection, s.Authenticator, s.MetaDataProvider)
}

// NewSessionBuilder creates a SessionBuilder object
func NewSessionBuilder() (*SessionBuilder, error) {
	return &SessionBuilder{}, nil
}
