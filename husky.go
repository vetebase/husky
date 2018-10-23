package husky

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// Husky struct holds router and context for framework
type Husky struct {
	AfterMiddleware  []MiddlewareHandler
	BeforeMiddleware []MiddlewareHandler
	Config           Configuration
	Context          *CTX
	Middleware       []MiddlewareHandler
	Router           *Router
}

// Handler basic function to router handlers
type Handler func(*CTX) error

// MiddlewareHandler defines a function to process middleware
type MiddlewareHandler func(Handler) Handler

// NotFoundHandler default 404 handler for not found routes
func NotFoundHandler(ctx *CTX) (err error) {
	b, _ := json.Marshal("Not Found")

	ctx.Response.Header().Set("Content-Type", "application/json")
	ctx.Response.WriteHeader(404)
	_, err = ctx.Response.Write([]byte(b))

	return
}

// New creates a new service
func New() (husky *Husky) {
	return &Husky{
		Router: new(Router),
	}
}

// GetContext returns current context
func (husky *Husky) GetContext() *CTX {
	return husky.Context
}

// After adds a handler to be executed after the route handler
// Executed if route is found or not
func (husky *Husky) After(middleware ...MiddlewareHandler) {
	for i := 0; i < len(middleware); i++ {
		husky.AfterMiddleware = append(husky.AfterMiddleware, middleware[i])
	}
}

// Before adds a handler to be executed before the route handler
// Executed if route is found or not
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

// Start initates the framework to start listening for requests
func (husky *Husky) Start() {
	server := husky.server()
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Server error: %s", err)
	}
}

func (husky *Husky) server() *http.Server {
	config := husky.Config.Load()

	name := config["NAME"]
	port := config["PORT"]

	server := &http.Server{
		Addr:    ":8080",
		Handler: husky,
	}

	fmt.Println("==> Running " + name + " on port: " + port)

	return server
}

// NewContext creates new Context struct
func (husky *Husky) NewContext(w http.ResponseWriter, r *http.Request) *CTX {
	return &CTX{
		Request:  r,
		Response: NewResponse(w),
	}
}

func (husky *Husky) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// create context
	husky.Context = husky.NewContext(w, r)
	var handler Handler

	// execute BeforeMiddleware
	if len(husky.BeforeMiddleware) > 0 {
		for i := 0; i < len(husky.BeforeMiddleware); i++ {
			husky.BeforeMiddleware[i](handler)
		}
	}

	// execute handler
	if found, route := husky.Router.FindRoute(husky.Context); found {
		handler := func(*CTX) error {
			handler := route.Handler

			// execute middleware chain
			if len(husky.Middleware) > 0 {
				for i := 0; i < len(husky.Middleware); i++ {
					handler = husky.Middleware[i](handler)
				}
			}

			// execute route
			if err := handler(husky.Context); err != nil {
				panic(err)
			}

			return nil
		}

		handler(husky.Context)
	} else {
		// route was not found
		NotFoundHandler(husky.Context)
	}

	// execute AfterMiddleware
	if len(husky.AfterMiddleware) > 0 {
		for i := 0; i < len(husky.AfterMiddleware); i++ {
			husky.AfterMiddleware[i](handler)
		}
	}

	return
}

func (husky *Husky) add(verb string, endpoint string, handler Handler, middleware []MiddlewareHandler) {
	path := strings.Split(endpoint, "?")
	husky.Router.Add(verb, path[0], handler, middleware)
}
