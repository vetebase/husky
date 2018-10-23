[![Build Status](https://travis-ci.org/vetebase/husky.svg?branch=master)](https://travis-ci.org/vetebase/husky)
[![codecov](https://codecov.io/gh/vetebase/husky/branch/master/graph/badge.svg)](https://codecov.io/gh/vetebase/husky)
[![GoDoc](https://godoc.org/github.com/olebedev/config?status.png)](https://godoc.org/github.com/vetebase/husky)
[![Go Report Card](https://goreportcard.com/badge/github.com/vetbase/husky)](https://goreportcard.com/report/github.com/vetebase/husky)

# Husky

Microservice framework written in Go.

## Example

```go
func main() {
    h := husky.New()

    // index handler
    h.GET("/", func(c husky.Context) error {
        return c.JSON(200, "Hello World!")
    })

    h.Start()
}
```

## Config

Husky provides the ability to configure your service from a central config file.
Copy `.env.example` to `.env` into the source of your service. You can place any
environment variables into this file. The `.env` is loaded into a Config struct
which can be read from anywhere in the service.

## Routes

### Add Routes

```go
// Add GET route
h.GET('/endpoint', handler)

// Add POST route
h.POST('/endpoint', handler)

// Add PATCH route
h.PATCH('/endpoint', handler)

// Add PUT route
h.PUT('/endpoint', handler)

// Add DELETE route
h.DELETE('/endpoint', handler)
```

## Middleware

### Included Middleware

Husky includes a couple of middleware (more to come) handlers which can be used
right out of the box.

#### JWT Middleware

```go
// the JWT middleware gets the key/secret from the config which is set in .env
h.GET("/endpoint", handler, middleware.JWT)
```

#### CORS Middleware

```go
h.GET("/endpoint", handler, middleware.CORSConfig{
  AllowedOrigins: "http://domain.com",  // default "*"
  AllowedMethods: "GET,POST,PUT,PATCH", // default "*"
  AllowedHeaders: "application/json",   // default "*"
  ExposedHeaders: "*",                  // default "*"
})
```

### Custom Middleware

Husky allows you to define your own custom middleware that can be used throughout
your service.

To create custom middleware:

```go
// Define middleware handler
middleware := func(c husky.Context, handler husky.Handler) husky.Handler {
    // code here
    return handler
}

// Add middleware to route
h.GET('/path', handler, middleware)
```

### How to Implement Middleware

#### Middleware

Adds a middleware handler to be executed after route is found but before
the handler is executed.

```go
h.Middleware(middleware1)
h.Middleware(middleware2)
-- or --
h.Middleware(middleware1, middleware2)
```

#### Before

Adds a middleware function to be executed before the route handler is executed.

```go
h.Before(middleware1)
h.Before(middleware2)
-- or --
h.Before(middleware1, middleware2)
```

#### After

Adds a middleware function to be executed after the route handler is executed.

```go
h.After(middleware1)
h.After(middleware2)
-- or --
h.After(middleware1, middleware2)
```

## Route Groups

```go
// Create Route Group
g := h.Group('/prefix')
g.GET('/endpoint', handler) // GET /prefix/endpoint
g.POST('/endpoint', handler) // POST /prefix/endpoint
```

### Middleware

Groups can have defined middleware. These middleware handlers will be executed for every route within the group:

```go
g := h.Group('/prefix', middlewareHandler1, middlewareHandler1)
```

Groups can have middleware handlers that are executed for every route within the group:

```go
g := h.Group('/prefix')
g.Middleware(middlwareHandler1)
g.Middleware(middlwareHandler2)
```

If you need middleware to be executed on specific routes you can add middleware to the route definition:

```go
g := h.Group('/prefix')
g.GET('/endpoint', handler, middleware)
```

## Development

Husky uses golang's [dep](https://github.com/golang/dep) for dependency management. Make sure dep is installed on your local development machine.

To pull in the required dependencies, run the following command: `dep ensure`.