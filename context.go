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

// Code writes header with HTTP code
func (ctx *CTX) Code(code int) (err error) {
	ctx.Response.WriteHeader(code)
	return nil
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

// GetHeader returns specified header
func (ctx *CTX) GetHeader(header string) string {
	return ctx.Request.Header.Get(header)
}

// GetParam return specified paramater
func (ctx *CTX) GetParam(i string) string {
	return ctx.Params[i]
}

// GetParams returns all stored parameters
func (ctx *CTX) GetParams() map[string]string {
	return ctx.Params
}

// HasParam checks if param is set
func (ctx *CTX) HasParam(param string) bool {
	_, isSet := ctx.Params[param]
	return isSet
}

// HTTPError returns a text/html error with requested code
func (ctx *CTX) HTTPError(code int, message string) (err error) {
	ctx.Response.Header().Set("Content-Type", "text/html;charset=utf-8")
	ctx.Response.WriteHeader(code)
	_, err = ctx.Response.Write([]byte(message))
	return
}

// Redirect returns a HTTP code
func (ctx *CTX) Redirect(code int, uri string) (err error) {
	http.Redirect(ctx.Response.Writer, ctx.Request, uri, code)
	return nil
}

// SetHeader adds header to response
func (ctx *CTX) SetHeader(k string, v string) {
	ctx.Response.Header().Set(k, v)
}

func (ctx *CTX) String(code int, s string) (err error) {
	ctx.Response.Header().Set("Content-Type", "text/html;charset=utf-8")
	ctx.Response.WriteHeader(code)
	_, err = ctx.Response.Write([]byte(s))
	return
}
