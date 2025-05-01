package httprequestmiddleware

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

// Helper function to capture stdout
func captureOutput(f func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // Execute the function whose output we want to capture

	w.Close()
	os.Stdout = old // restoring the real stdout
	out, _ := io.ReadAll(r)
	return string(out)
}

// Test suite setup
func setupTestChain() Middleware {
	logger := &LoggingMiddleware{}
	authenticator := &AuthenticationMiddleware{}
	authorizer := &AuthorizationMiddleware{}
	finalHandler := &FinalHandler{}

	// Build chain: Log -> Auth -> Authorize -> Final
	logger.SetNext(authenticator).SetNext(authorizer).SetNext(finalHandler)
	return logger // Return the entry point
}

func TestMiddlewareChain(t *testing.T) {

	t.Run("Successful Request Admin", func(t *testing.T) {
		chain := setupTestChain()
		req := MakeRequest("/admin", "valid_token", "admin")

		var resp *Response
		var err error
		output := captureOutput(func() {
			resp, err = chain.Handle(req)
		})

		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if resp == nil {
			t.Fatal("Expected a response, but got nil")
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
		}
		if !strings.Contains(resp.Body, "Successfully processed") {
			t.Errorf("Response body mismatch: %s", resp.Body)
		}
		if !strings.Contains(resp.Body, "user 789") {
			t.Errorf("Response body does not contain correct user ID: %s", resp.Body)
		}

		// Check logs
		if !strings.Contains(output, "[Log] Received request") { t.Error("Missing Log output") }
		if !strings.Contains(output, "[Auth] Authentication successful") { t.Error("Missing Auth success output") }
		if !strings.Contains(output, "[ACL] Authorization successful") { t.Error("Missing ACL success output") }
		if !strings.Contains(output, "[Final] Processing request") { t.Error("Missing Final processing output") }
	})

	t.Run("Authentication Failure", func(t *testing.T) {
		chain := setupTestChain()
		req := MakeRequest("/secure", "invalid_token", "")

		var resp *Response
		var err error
		output := captureOutput(func() {
			resp, err = chain.Handle(req)
		})

		if !errors.Is(err, ErrAuthenticationFailed) {
			t.Errorf("Expected ErrAuthenticationFailed, but got: %v", err)
		}
		if resp == nil {
			t.Fatal("Expected a response, but got nil")
		}
		if resp.StatusCode != 401 {
			t.Errorf("Expected status code 401, but got %d", resp.StatusCode)
		}
		if resp.Body != "Unauthorized" {
			t.Errorf("Expected body 'Unauthorized', got '%s'", resp.Body)
		}

		// Check logs show stop at Auth
		if !strings.Contains(output, "[Log] Received request") { t.Error("Missing Log output") }
		if !strings.Contains(output, "[Auth] Authentication failed") { t.Error("Missing Auth failure output") }
		if strings.Contains(output, "[ACL] Checking authorization") { t.Error("ACL should not have been checked") }
		if strings.Contains(output, "[Final] Processing request") { t.Error("Final Handler should not have been reached") }
	})

	t.Run("Authorization Failure", func(t *testing.T) {
		chain := setupTestChain()
		req := MakeRequest("/admin/super", "valid_token", "superadmin")

		var resp *Response
		var err error
		output := captureOutput(func() {
			resp, err = chain.Handle(req)
		})

		if !errors.Is(err, ErrAuthorizationFailed) {
			t.Errorf("Expected ErrAuthorizationFailed, but got: %v", err)
		}
		if resp == nil {
			t.Fatal("Expected a response, but got nil")
		}
		if resp.StatusCode != 403 {
			t.Errorf("Expected status code 403, but got %d", resp.StatusCode)
		}
		if resp.Body != "Forbidden" {
			t.Errorf("Expected body 'Forbidden', got '%s'", resp.Body)
		}

		// Check logs show stop at ACL
		if !strings.Contains(output, "[Log] Received request") { t.Error("Missing Log output") }
		if !strings.Contains(output, "[Auth] Authentication successful") { t.Error("Missing Auth success output") }
		if !strings.Contains(output, "[ACL] Authorization failed") { t.Error("Missing ACL failure output") }
		if strings.Contains(output, "[Final] Processing request") { t.Error("Final Handler should not have been reached") }
	})

	t.Run("Successful Request No Role Required", func(t *testing.T) {
		chain := setupTestChain()
		req := MakeRequest("/public", "valid_token", "") // No role required

		var resp *Response
		var err error
		output := captureOutput(func() {
			resp, err = chain.Handle(req)
		})


		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}
		if resp == nil {
			t.Fatal("Expected a response, but got nil")
		}
		if resp.StatusCode != 200 {
			t.Errorf("Expected status code 200, but got %d", resp.StatusCode)
		}
		if !strings.Contains(resp.Body, "Successfully processed") {
			t.Errorf("Response body mismatch: %s", resp.Body)
		}

		// Check logs
		if !strings.Contains(output, "[Log] Received request") { t.Error("Missing Log output") }
		if !strings.Contains(output, "[Auth] Authentication successful") { t.Error("Missing Auth success output") }
        if !strings.Contains(output, "[ACL] No specific role required. Proceeding.") { t.Error("Missing ACL no role required output") }
		if !strings.Contains(output, "[Final] Processing request") { t.Error("Missing Final processing output") }
	})
}
