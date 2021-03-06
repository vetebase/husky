package husky

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddingRoute(t *testing.T) {
	h := New()

	h.GET("/path", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/path"]).String() == "husky.Route")
}

func TestAddingRouteWithMiddleware(t *testing.T) {
	h := New()

	h.GET("/path", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	}, middlware)

	found := h.Router.GetRoutes("GET")

	assert.Equal(t, 1, len(found["GET/path"].Middleware))
}

func TestHandlerReturnsHandler(t *testing.T) {
	h := New()

	h.GET("/path", func(ctx *CTX) error {
		return ctx.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("GET")
	route := found["GET/path"]

	assert.True(t, reflect.TypeOf(route.Handler).String() == "husky.Handler")
}

func TestGetRoutesReturnsRoutes(t *testing.T) {
	h := New()

	h.GET("/path", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/path"]).String() == "husky.Route")
}

func TestGetRoutesReturnsEmptyIfRouteNotFound(t *testing.T) {
	h := New()

	h.GET("/path", func(c *CTX) error {
		return c.JSON(200, "This is a test")
	})

	found := h.Router.GetRoutes("POST")

	assert.Empty(t, found)
}

func TestRootPathReturnsEmptyMap(t *testing.T) {
	h := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	h.Context = h.NewContext(w, r)

	assert.Empty(t, h.Context.GetParams())
	assert.True(t, reflect.TypeOf(h.Context.GetParams()).String() == "map[string]string")
}
