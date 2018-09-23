package husky

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
func (h *Husky) After(middleware ...MiddlewareHandler) {
	for i := 0; i < len(middleware); i++ {
		h.AfterMiddleware = append(h.AfterMiddleware, middleware[i])
	}
}

// Before adds a handler to be executed before the route handler
func (h *Husky) Before(middleware ...MiddlewareHandler) {
	for i := 0; i < len(middleware); i++ {
		h.BeforeMiddleware = append(h.BeforeMiddleware, middleware[i])
	}
}

// Middlware adds a handler to be executed before the route handler
func (h *Husky) Middlware(middleware MiddlewareHandler) {
	h.Middleware = append(h.Middleware, middleware)
}
