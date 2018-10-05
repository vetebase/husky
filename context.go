package husky

import (
	"encoding/json"
	"net/http"
)

// CTX (Context) struct
type CTX struct {
	Request  *http.Request
	Response *Response
	Params   map[string]string
}

// Context interface
type Context interface {
	AddParams(map[string]string)
	HTTPError(int, string) error
	JSON(int, interface{}) error
	Request() *http.Request
	Response() *http.Response
}

// AddParams adds parameters to context
func (ctx *CTX) AddParams(params map[string]string) {
	if ctx.Params == nil {
		ctx.Params = make(map[string]string)
	}

	for k, v := range params {
		ctx.Params[k] = v
	}

	return
}

// JSON returns response as serialized JSON
func (ctx *CTX) JSON(code int, i interface{}) (err error) {
	b, err := json.Marshal(i)

	if err != nil {
		ctx.HTTPError(500, err.Error())
		return
	}

	ctx.Response.Header().Set("Content-Type", "application/json")
	ctx.Response.WriteHeader(code)
	_, err = ctx.Response.Write([]byte(b))
	return
}

// Code writes header with HTTP code
func (ctx *CTX) Code(code int) (err error) {
	ctx.Response.WriteHeader(code)
	return nil
}

// HTTPError returns a text/html error with requested code
func (ctx *CTX) HTTPError(code int, message string) (err error) {
	ctx.Response.Header().Set("Content-Type", "text/html;charset=utf-8")
	ctx.Response.WriteHeader(code)
	_, err = ctx.Response.Write([]byte(message))
	return
}
