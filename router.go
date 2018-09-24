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
	Handler    Handler             // main handler
	Endpoint   string              // endpoint for route
	Middleware []MiddlewareHandler // array of middleware handlers
	Verb       string              // http verb
}

// Add will add a new route to the Router.Routes map
func (router *Router) Add(verb string, endpoint string, handler Handler, middleware []MiddlewareHandler) {
	route := Route{
		Endpoint: endpoint,
		Handler:  handler,
		Verb:     verb,
	}

	// add middleware handler(s)
	for _, v := range middleware {
		route.Middleware = append(route.Middleware, v)
	}

	// initialize router routes
	if router.Routes == nil {
		router.Routes = make(map[string]map[string]Route)
	}

	// initialize verb
	if router.Routes[verb] == nil {
		router.Routes[verb] = make(map[string]Route)
	}

	router.Routes[verb][verb+endpoint] = route
}
