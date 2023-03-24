package middleware

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"ryan-test-backend/model/web"
	"testing"
)

func TestAuthMiddleware_ServeHTTP(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	middleware := NewAuthMiddleware(handler)

	t.Run("Authorized request", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)
		req.Header.Set("X-API-Key", "api-key-rahasia")

		recorder := httptest.NewRecorder()
		middleware.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusOK {
			t.Errorf("Expected status code %d, but got %d", http.StatusOK, recorder.Code)
		}
	})

	t.Run("Unauthorized request", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/", nil)

		recorder := httptest.NewRecorder()
		middleware.ServeHTTP(recorder, req)

		if recorder.Code != http.StatusUnauthorized {
			t.Errorf("Expected status code %d, but got %d", http.StatusUnauthorized, recorder.Code)
		}

		var webResponse web.WebResponse
		err := json.NewDecoder(recorder.Body).Decode(&webResponse)
		if err != nil {
			t.Errorf("Error decoding response body: %v", err)
		}

		if webResponse.Code != http.StatusUnauthorized {
			t.Errorf("Expected web response code %d, but got %d", http.StatusUnauthorized, webResponse.Code)
		}

		if webResponse.Status != "UNAUTHORIZED" {
			t.Errorf("Expected web response status %q, but got %q", "UNAUTHORIZED", webResponse.Status)
		}
	})
}
