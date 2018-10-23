package husky

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockedHusky struct {
	mock.Mock
}

func handler(c *CTX) error {
	return nil
}

func middlware(handler Handler) Handler {
	return handler
}

func TestNewHusky(t *testing.T) {
	h := New()
	assert.NotNil(t, h)
}

func TestRouterIsHuskyRouter(t *testing.T) {
	h := New()
	assert.True(t, reflect.TypeOf(h.Router).String() == "*husky.Router")
}

func TestAddBeforeMiddleware(t *testing.T) {
	h := New()

	mw1 := func(handler Handler) Handler {
		return handler
	}

	mw2 := func(handler Handler) Handler {
		return handler
	}

	h.Before(mw1, mw2)

	assert.True(t, len(h.BeforeMiddleware) == 2)
}

func TestAddAfterMiddleware(t *testing.T) {
	h := New()

	mw1 := func(handler Handler) Handler {
		return handler
	}

	mw2 := func(handler Handler) Handler {
		return handler
	}

	h.After(mw1, mw2)

	assert.True(t, len(h.AfterMiddleware) == 2)
}

func TestAddMiddleware(t *testing.T) {
	h := New()

	mw1 := func(handler Handler) Handler {
		return handler
	}

	h.Middlware(mw1)

	assert.True(t, len(h.Middleware) == 1)
}

func TestCreateGroup(t *testing.T) {
	h := New()

	g := h.Group("/group")
	g.GET("/test", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("GET")

	assert.True(t, len(found) == 1)
}

func TestGroupRouteAddedCorrectly(t *testing.T) {
	h := New()

	g := h.Group("/group")
	g.GET("/test", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/group/test"]).String() == "husky.Route")
}

// func TestGroupRouteMiddlewareAddedCorrectly(t *testing.T) {
// 	h := New()

// 	g := h.Group("/group", middlware)
// 	g.GET("/test", func(c *CTX) error {
// 		return c.JSON(200, "This is a test")
// 	})

// 	assert.Equal(t, len(g.MiddlewareHandlers), 1)
// }

func TestAddGetRoute(t *testing.T) {
	h := New()

	h.GET("/test", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/test"]).String() == "husky.Route")
}

func TestAddPostRoute(t *testing.T) {
	h := New()

	h.POST("/test", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("POST")

	assert.True(t, reflect.TypeOf(found["POST/test"]).String() == "husky.Route")
}

// func TestConfigLoad(t *testing.T) {
// 	config := config.Load()

// 	assert.True(t, reflect.TypeOf(config).String() == "map[string]string")
// }

func TestAddPatchRoute(t *testing.T) {
	h := New()

	h.PATCH("/test", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("PATCH")

	assert.True(t, reflect.TypeOf(found["PATCH/test"]).String() == "husky.Route")
}

func TestAddPutRoute(t *testing.T) {
	h := New()

	h.PUT("/test", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("PUT")

	assert.True(t, reflect.TypeOf(found["PUT/test"]).String() == "husky.Route")
}

func TestAddDeleteRoute(t *testing.T) {
	h := New()

	h.DELETE("/test", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("DELETE")

	assert.True(t, reflect.TypeOf(found["DELETE/test"]).String() == "husky.Route")
}

func TestAddOptionsRoute(t *testing.T) {
	h := New()

	h.OPTIONS("/test", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("OPTIONS")

	assert.True(t, reflect.TypeOf(found["OPTIONS/test"]).String() == "husky.Route")
}

func TestGetContext(t *testing.T) {
	h := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	h.Context = h.NewContext(w, r)

	assert.True(t, reflect.TypeOf(h.GetContext()).String() == "*husky.CTX")
}

func TestNotFoundHandler(t *testing.T) {
	h := New()

	h.GET("/", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	r, _ := http.NewRequest("GET", "/blah", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	h.ServeHTTP(w, r)

	if assert.Equal(t, 404, w.Code) {
		assert.Equal(t, "\"Not Found\"", w.Body.String())
	}
}

func TestNewServerReturnsHTTPServer(t *testing.T) {
	h := New()
	server := h.server()

	assert.True(t, reflect.TypeOf(server).String() == "*http.Server")
}

func TestNewServerRunsOnCorrectPort(t *testing.T) {
	h := New()
	server := h.server()

	assert.Equal(t, ":8080", server.Addr)
}
