package husky

// Router holds all defined routes
// map[string]map[string] = [VERB][PATTERN] = [GET][/users]
type Router struct {
	routes map[string]map[string]Route
}

// Route holds all information about a defined route
type Route struct {
	method     string
	handler    Handler
	middleware []MiddlewareHandler
	path       string
}
