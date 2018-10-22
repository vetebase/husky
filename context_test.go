package husky

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	id   int
	name string
}

const JSON = `{"id":1,"name":"John Adams"}`

func TestContext(t *testing.T) {
	h := New()

	r, _ := http.NewRequest("GET", "/", strings.NewReader(JSON))
	w := httptest.NewRecorder()

	c := h.NewContext(w, r)

	// Response
	assert.NotNil(t, c.Response)

	// Request
	assert.NotNil(t, c.Request)
}
