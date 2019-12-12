package test

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetMockedContext(t *testing.T) {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:123/test", nil)
	assert.Nil(t, err)
	response := httptest.NewRecorder()
	request.Header = http.Header{
		"X-mock":{"true"},
	}

	c := GetMockedContext(request, response)

	assert.EqualValues(t, http.MethodGet, c.Request.Method)
	assert.EqualValues(t, "123", c.Request.URL.Port())
	assert.EqualValues(t, "/test", c.Request.URL.Path)
	assert.EqualValues(t, "http", c.Request.URL.Scheme)
}
