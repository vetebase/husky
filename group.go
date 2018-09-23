package husky

// Group holds information about the route group
type Group struct {
	Husky      *Husky
	Middleware []MiddlewareHandler
	Prefix     string
}

// GET adds a HTTP Get method to the group
func (g *Group) GET(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	g.add("GET", endpoint, handler, middleware)
}

func (g *Group) add(verb string, endpoint string, handler Handler, middleware []MiddlewareHandler) {
	g.Husky.add(verb, g.Prefix+endpoint, handler, middleware)
}
