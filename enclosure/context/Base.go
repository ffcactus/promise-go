package context

import (
	"net/http"
)

// Base containes the basic context info
type Base struct {
	ID      string
	Request *http.Request
}
