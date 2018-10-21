package husky

import (
	"bytes"
	"regexp"
	"strings"
)

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

// FindRoute searches for requested route
func (router *Router) FindRoute(ctx *CTX) (bool, Route) {
	// by default route is nil, i.e. Not Found
	var route Route
	found := false

	httpMethod := ctx.Request.Method
	httpURI := strings.Split(ctx.Request.URL.String(), "?")

	// routes to search
	// routes := router.getRoutes(context.Request().Method)
	for k, v := range router.Routes[ctx.Request.Method] {
		formatted := format(k)

		regex := regexp.MustCompile(`^` + formatted + `/?$`)

		if regex.MatchString(httpMethod + httpURI[0]) {
			found = true
			route = v

			ctx.AddParams(parseURLParams(httpMethod, httpURI[0], formatted, k))

			ctx.Request.ParseForm()
			ctx.AddParams(parseFormParams(ctx.Request.Form))

			if len(httpURI) > 1 {
				ctx.AddParams(parseQueryParams(httpURI[1]))
			}
		}
	}

	return found, route
}

// func (router *Router) getRoutes(method string) map[string]Route {
// 	if val, exists := router.Routes[method]; exists {
// 		return val
// 	}

// 	return nil
// }

func format(route string) string {
	var formatted bytes.Buffer
	re := regexp.MustCompile(`:` + pattern)
	formatted.WriteString(re.ReplaceAllString(route, pattern))
	return formatted.String()
}

func parseURLParams(method string, url string, path string, route string) map[string]string {
	if path == "/" {
		return make(map[string]string)
	}

	// map of params to be returned
	params := make(map[string]string)

	// key regular expression (kre)
	kre := regexp.MustCompile(`:` + pattern)
	keys := kre.FindAllStringSubmatch(route, -1)

	// value regular express (vre)
	vre := regexp.MustCompile(path)
	values := vre.FindAllStringSubmatch(method+url, -1)[0][1:]

	// assign keys to values
	for i, v := range keys {
		params[v[1]] = values[i]
	}

	return params
}

func parseQueryParams(url string) map[string]string {
	params := make(map[string]string)

	qre := regexp.MustCompile(query)
	q := qre.FindAllStringSubmatch(url, -1)

	for _, query := range q {
		values := strings.Split(query[0], "=")
		params[values[0]] = values[1]
	}

	params["param"] = "1"
	return params
}

func parseFormParams(form map[string][]string) map[string]string {
	params := make(map[string]string)

	for k, v := range form {
		params[k] = v[0]
	}

	return params
}
