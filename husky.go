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

// Group creates a route group with a common prefix
func (husky *Husky) Group(prefix string, middleware ...MiddlewareHandler) *Group {
	group := &Group{Prefix: prefix, Husky: husky}
	group.Middleware = append(group.Middleware, middleware...)
	return group
}

func (husky *Husky) add(action string, endpoint string, handler Handler, middleware []MiddlewareHandler) {
	path := strings.Split(endpoint, "?")
	husky.Router.Add(action, path[0], func(c Context) error {
		handler := handler
		return handler(c)
	}, middleware)
}
