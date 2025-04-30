package request_pipeline

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBaseHandler(t *testing.T) {
	request := Request{
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    map[string]interface{}{"data": "test"},
		Path:    "/api/test",
		Method:  "POST",
	}

	handler := &BaseHandler{}
	response := handler.Handle(request)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Request processed successfully", response.Body["message"])
}

func TestAuthMiddleware(t *testing.T) {
	request := Request{
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    map[string]interface{}{"data": "test"},
		Path:    "/api/test",
		Method:  "POST",
	}

	// Test unauthorized request
	handler := NewAuthMiddleware(&BaseHandler{})
	response := handler.Handle(request)

	assert.Equal(t, 401, response.StatusCode)
	assert.Equal(t, "Unauthorized", response.Body["error"])

	// Test authorized request
	request.Headers["Authorization"] = "Bearer valid-token"
	response = handler.Handle(request)

	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Request processed successfully", response.Body["message"])
}

func TestRateLimitMiddleware(t *testing.T) {
	request := Request{
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    map[string]interface{}{"data": "test"},
		Path:    "/api/test",
		Method:  "POST",
	}

	handler := NewRateLimitMiddleware(&BaseHandler{}, 2)

	// First request should succeed
	response := handler.Handle(request)
	assert.Equal(t, 200, response.StatusCode)

	// Second request should succeed
	response = handler.Handle(request)
	assert.Equal(t, 200, response.StatusCode)

	// Third request should fail
	response = handler.Handle(request)
	assert.Equal(t, 429, response.StatusCode)
	assert.Equal(t, "Too many requests", response.Body["error"])
}

func TestMiddlewareChain(t *testing.T) {
	request := Request{
		Headers: map[string]string{"Content-Type": "application/json"},
		Body:    map[string]interface{}{"data": "test"},
		Path:    "/api/test",
		Method:  "POST",
	}

	// Create a chain: Auth -> RateLimit -> BaseHandler
	handler := NewAuthMiddleware(
		NewRateLimitMiddleware(
			&BaseHandler{},
			2,
		),
	)

	// Test with valid auth
	request.Headers["Authorization"] = "Bearer valid-token"
	response := handler.Handle(request)
	assert.Equal(t, 200, response.StatusCode)

	// Test with invalid auth
	delete(request.Headers, "Authorization")
	response = handler.Handle(request)
	assert.Equal(t, 401, response.StatusCode)
} 