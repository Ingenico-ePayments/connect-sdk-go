package connectsdk

// CallContext can be used to send extra information with a request, and to receive extra information from a response.
// Please note that this type is not thread-safe. Each request should get its own call context instance.
type CallContext struct {
	IdempotenceKey              string
	IdempotenceRequestTimestamp *int64
}

// GetIdempotenceKey returns the idempotence key
func (c *CallContext) GetIdempotenceKey() string {
	return c.IdempotenceKey
}

// GetIdempotenceRequestTimestamp returns the idempotence timestamp
func (c *CallContext) GetIdempotenceRequestTimestamp() *int64 {
	if c.IdempotenceRequestTimestamp == nil {
		return nil
	}

	timestamp := *c.IdempotenceRequestTimestamp

	return &timestamp
}

// SetIdempotenceRequestTimestamp sets the idempotence timestamp
func (c *CallContext) SetIdempotenceRequestTimestamp(timestamp *int64) {
	c.IdempotenceRequestTimestamp = timestamp
}

// NewCallContext creates a CallContext using the given idempotenceKey
func NewCallContext(idempotenceKey string) (*CallContext, error) {
	return &CallContext{idempotenceKey, nil}, nil
}
