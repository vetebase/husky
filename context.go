package husky

import "net/http"

// Context interface
type Context interface {
	Request() *http.Request
	Response() *http.Response
}
