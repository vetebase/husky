package husky

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
