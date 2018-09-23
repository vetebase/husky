package husky

const pattern = `([aA-zZ0-9_-]+)`
const query = `[^&?]*?=[^&?]*`

// Router holds all defined routes
// map[string]map[string] = [VERB][PATTERN] = [GET][/users]
type Router struct {
	Routes map[string]map[string]Route
}

// Route holds all information about a defined route
type Route struct {
	Method     string
	Handler    Handler
	Middleware []MiddlewareHandler
	Path       string
}

// Add will add a new route to the Router.Routes map
func (router *Router) Add(method string, path string, handler Handler, middleware []MiddlewareHandler) {
	route := Route{
		Handler: handler,
		Method:  method,
		Path:    path,
	}

	// add middleware handler(s)
	for _, v := range middleware {
		route.Middleware = append(route.Middleware, v)
	}

	// initialize router routes
	if router.Routes == nil {
		router.Routes = make(map[string]map[string]Route)
	}

	// initialize method
	if router.Routes[method] == nil {
		router.Routes[method] = make(map[string]Route)
	}

	router.Routes[method][method+path] = route
}
