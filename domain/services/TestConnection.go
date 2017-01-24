// This class was auto-generated from the API references found at
// https://developer.globalcollect.com/documentation/api/server/

package services

// TestConnection represents class TestConnection
// Documentation can be found at https://developer.globalcollect.com/documentation/api/server/#schema_TestConnection
type TestConnection struct {
	Result *string `json:"result,omitempty"`
}

// NewTestConnection constructs a new TestConnection
func NewTestConnection() *TestConnection {
	return &TestConnection{}
}
