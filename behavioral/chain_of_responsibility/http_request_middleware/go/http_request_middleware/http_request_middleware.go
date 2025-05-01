package httprequestmiddleware

import (
	"errors"
	"fmt"
)

// User represents user info added during authentication
type User struct {
	ID   int
	Role string
}

// Request represents the incoming HTTP request (simplified)
type Request struct {
	Path         string
	Headers      map[string]string
	User         *User // Pointer to allow nil for unauthenticated
	RequiredRole string
	// Add Body, Method etc. as needed
}

// Response represents the outgoing HTTP response (simplified)
type Response struct {
	StatusCode int
	Body       string
}

// Error constants for middleware failures
var (
	ErrAuthenticationFailed = errors.New("authentication failed")
	ErrAuthorizationFailed  = errors.New("authorization failed")
	ErrAuthRequired         = errors.New("authentication required for authorization check")
)

// Middleware interface defines the contract for handlers
type Middleware interface {
	SetNext(Middleware) Middleware
	Handle(*Request) (*Response, error) // Return Response for success/handled, error for failure/stop chain
}

// BaseMiddleware provides common linking functionality
type BaseMiddleware struct {
	next Middleware
}

func (b *BaseMiddleware) SetNext(m Middleware) Middleware {
	b.next = m
	return m // Allow chaining
}

// Handle passes to the next middleware if it exists
func (b *BaseMiddleware) Handle(req *Request) (*Response, error) {
	if b.next != nil {
		return b.next.Handle(req)
	}
	// End of chain default behavior (can be customized)
	// Returning nil response and nil error indicates chain completed without a specific final handler action
	// Or potentially return a default "Not Found" response here
	return nil, nil
}

// --- Concrete Middleware Implementations ---

// LoggingMiddleware logs the request
type LoggingMiddleware struct {
	BaseMiddleware // Embed BaseMiddleware
}

func (l *LoggingMiddleware) Handle(req *Request) (*Response, error) {
	fmt.Printf("[Log] Received request: %s\n", req.Path)
	// Continue chain
	return l.next.Handle(req) // Explicitly call next handler's Handle
}

// AuthenticationMiddleware checks authentication
type AuthenticationMiddleware struct {
	BaseMiddleware
}

func (a *AuthenticationMiddleware) Handle(req *Request) (*Response, error) {
	fmt.Println("[Auth] Checking authentication...")
	token, ok := req.Headers["Authorization"]

	if ok && token == "valid_token" {
		fmt.Println("[Auth] Authentication successful.")
		// Add user info to the request
		req.User = &User{ID: 789, Role: "admin"}
		return a.next.Handle(req) // Continue chain
	}

	fmt.Println("[Auth] Authentication failed. Aborting request.")
	// Stop chain, return error response
	return &Response{StatusCode: 401, Body: "Unauthorized"}, ErrAuthenticationFailed
}

// AuthorizationMiddleware checks permissions
type AuthorizationMiddleware struct {
	BaseMiddleware
}

func (a *AuthorizationMiddleware) Handle(req *Request) (*Response, error) {
	fmt.Println("[ACL] Checking authorization...")

	if req.User == nil {
		fmt.Println("[ACL] No user found in request (likely due to failed auth). Aborting.")
		return &Response{StatusCode: 401, Body: "Authentication Required"}, ErrAuthRequired
	}

	if req.RequiredRole != "" && req.User.Role == req.RequiredRole {
		fmt.Printf("[ACL] Authorization successful (Role: %s).\n", req.User.Role)
		return a.next.Handle(req) // Continue chain
	} else if req.RequiredRole == "" {
		fmt.Println("[ACL] No specific role required. Proceeding.")
		return a.next.Handle(req) // Continue chain
	} else {
		fmt.Printf("[ACL] Authorization failed. User role '%s' does not match required role '%s'. Aborting.\n", req.User.Role, req.RequiredRole)
		return &Response{StatusCode: 403, Body: "Forbidden"}, ErrAuthorizationFailed
	}
}

// FinalHandler is the last step in the chain
type FinalHandler struct {
	BaseMiddleware // Embed BaseMiddleware although it won't use 'next'
}

func (f *FinalHandler) Handle(req *Request) (*Response, error) {
	fmt.Println("[Final] Processing request...")
	userID := "anonymous"
	if req.User != nil {
		userID = fmt.Sprintf("%d", req.User.ID)
	}
	body := fmt.Sprintf("Successfully processed request for path: %s for user %s", req.Path, userID)
	fmt.Println("[Final] Request processed.")
	// Successful processing, return final response
	return &Response{StatusCode: 200, Body: body}, nil
}

// MakeRequest is a helper to create requests
func MakeRequest(path string, token string, requiresRole string) *Request {
	req := &Request{
		Path:    path,
		Headers: make(map[string]string),
	}
	if token != "" {
		req.Headers["Authorization"] = token
	}
	if requiresRole != "" {
		req.RequiredRole = requiresRole
	}
	return req
}
