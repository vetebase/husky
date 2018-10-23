package husky

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupCreation(t *testing.T) {
	h := New()

	g := h.Group("/group")

	assert.True(t, reflect.TypeOf(g).String() == "*husky.Group")
}

func TestGroupGetRoute(t *testing.T) {
	h := New()

	g := h.Group("/group")
	g.GET("/path", func(c *CTX) error {
		return c.String(200, "This is a test")
	})

	found := h.Router.GetRoutes("GET")

	assert.True(t, reflect.TypeOf(found["GET/group/path"]).String() == "husky.Route")
}

func TestGroupPostRoute(t *testing.T) {
	h := New()

	g := h.Group("/group")
	g.POST("/path", func(c *CTX) error {
		return c.String(200, "This is a test")
	})

	found := h.Router.GetRoutes("POST")

	assert.True(t, reflect.TypeOf(found["POST/group/path"]).String() == "husky.Route")
}

func TestGroupPatchRoute(t *testing.T) {
	h := New()

	g := h.Group("/group")
	g.PATCH("/path", func(c *CTX) error {
		return c.String(200, "This is a test")
	})

	found := h.Router.GetRoutes("PATCH")

	assert.True(t, reflect.TypeOf(found["PATCH/group/path"]).String() == "husky.Route")
}

func TestGroupPutRoute(t *testing.T) {
	h := New()

	g := h.Group("/group")
	g.PUT("/path", func(c *CTX) error {
		return c.String(200, "This is a test")
	})

	found := h.Router.GetRoutes("PUT")

	assert.True(t, reflect.TypeOf(found["PUT/group/path"]).String() == "husky.Route")
}

func TestGroupDeleteRoute(t *testing.T) {
	h := New()

	g := h.Group("/group")
	g.DELETE("/path", func(c *CTX) error {
		return c.String(200, "This is a test")
	})

	found := h.Router.GetRoutes("DELETE")

	assert.True(t, reflect.TypeOf(found["DELETE/group/path"]).String() == "husky.Route")
}

func TestAddGroupMiddleware(t *testing.T) {
	h := New()

	g := h.Group("/group")
	gmw1 := func(handler Handler) Handler {
		return handler
	}

	gmw2 := func(handler Handler) Handler {
		return handler
	}

	g.Middleware(gmw1)
	g.Middleware(gmw2)

	assert.True(t, len(g.MiddlewareHandlers) == 2)
}
