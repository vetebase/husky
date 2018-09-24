package husky

import "strings"

// Husky struct holds router and context for framework
type Husky struct {
	AfterMiddleware  []MiddlewareHandler
	BeforeMiddleware []MiddlewareHandler
	Context          Context
	Middleware       []MiddlewareHandler
	Router           *Router
}

// Handler basic function to router handlers
type Handler func(Context) error

// MiddlewareHandler defines a function to process middleware
type MiddlewareHandler func(Handler) Handler

// New creates a new service
func New() (h *Husky) {
	return &Husky{
		Router: new(Router),
	}
}

// After adds a handler to be executed after the route handler
func (husky *Husky) After(middleware ...MiddlewareHandler) {
	for i := 0; i < len(middleware); i++ {
		husky.AfterMiddleware = append(husky.AfterMiddleware, middleware[i])
	}
}

// Before adds a handler to be executed before the route handler
func (husky *Husky) Before(middleware ...MiddlewareHandler) {
	for i := 0; i < len(middleware); i++ {
		husky.BeforeMiddleware = append(husky.BeforeMiddleware, middleware[i])
	}
}

// Middlware adds a handler to be executed before the route handler
func (husky *Husky) Middlware(middleware MiddlewareHandler) {
	husky.Middleware = append(husky.Middleware, middleware)
}

// DELETE adds a HTTP DELETE route to router
func (husky *Husky) DELETE(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	husky.add("DELETE", endpoint, handler, middleware)
}

// GET adds a HTTP GET route to router
func (husky *Husky) GET(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	husky.add("GET", endpoint, handler, middleware)
}

// OPTIONS adds a HTTP OPTIONS route to router
func (husky *Husky) OPTIONS(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	husky.add("OPTIONS", endpoint, handler, middleware)
}

// PATCH adds a HTTP PATCH route to router
func (husky *Husky) PATCH(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	husky.add("PATCH", endpoint, handler, middleware)
}

// POST adds a HTTP POST route to router
func (husky *Husky) POST(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	husky.add("POST", endpoint, handler, middleware)
}

// PUT adds a HTTP PUT route to router
func (husky *Husky) PUT(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	husky.add("PUT", endpoint, handler, middleware)
}

// Group creates a route group with a common prefix
func (husky *Husky) Group(prefix string, middleware ...MiddlewareHandler) *Group {
	group := &Group{Prefix: prefix, Husky: husky}
	group.Middleware = append(group.Middleware, middleware...)
	return group
}

func (husky *Husky) add(verb string, endpoint string, handler Handler, middleware []MiddlewareHandler) {
	path := strings.Split(endpoint, "?")
	husky.Router.Add(verb, path[0], func(c Context) error {
		handler := handler
		return handler(c)
	}, middleware)
}
