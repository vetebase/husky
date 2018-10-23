package husky

// Group holds information about the route group
type Group struct {
	Husky              *Husky
	MiddlewareHandlers []MiddlewareHandler
	Prefix             string
}

// GET adds a HTTP Get method to the group
func (g *Group) GET(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	g.add("GET", endpoint, handler, middleware)
}

// POST adds a HTTP POST method to the group
func (g *Group) POST(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	g.add("POST", endpoint, handler, middleware)
}

// PATCH adds a HTTP PATCH method to the group
func (g *Group) PATCH(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	g.add("PATCH", endpoint, handler, middleware)
}

// PUT adds a HTTP PUT method to the group
func (g *Group) PUT(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	g.add("PUT", endpoint, handler, middleware)
}

// DELETE adds a HTTP DELETE method to the group
func (g *Group) DELETE(endpoint string, handler Handler, middleware ...MiddlewareHandler) {
	g.add("DELETE", endpoint, handler, middleware)
}

// Middleware adds a middleware handler to be executed after route is found
// but before the handler is executed
func (g *Group) Middleware(m MiddlewareHandler) {
	g.MiddlewareHandlers = append(g.MiddlewareHandlers, m)
}

func (g *Group) add(verb string, endpoint string, handler Handler, middleware []MiddlewareHandler) {
	g.Husky.add(verb, g.Prefix+endpoint, handler, middleware)
}
