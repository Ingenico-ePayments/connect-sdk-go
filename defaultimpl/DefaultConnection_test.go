package defaultimpl

import (
	"net/http"
	"testing"
	"time"
)

const socketTimeout = 10000 * time.Microsecond
const connectTimeout = 100 * time.Microsecond
const keepAliveTimeout = 100 * time.Microsecond
const idleTimeout = 10 * time.Second
const maxConnections = 100

func TestConstructWithoutProxy(t *testing.T) {
	connection, _ := NewDefaultConnection(socketTimeout, connectTimeout, keepAliveTimeout, idleTimeout, maxConnections, nil)
	// Cannot recover connectTimeout, because it is embedded in the dial function

	// Cannot recover proxy, because it is embedded in a function that is never nil, and functions also cannot be compared in go
	if transport, ok := connection.client.Transport.(*http.Transport); ok {
		if transport.MaxIdleConns != maxConnections {
			t.Fatal("Max connections not used")
		}
	}

	if connection.client.Timeout != socketTimeout {
		t.Fatal("Socket timeout not used")
	}
}
