package exception

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-playground/validator"
)

func Test_notFoundError(t *testing.T) {
	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/", nil)

	// Create a new validation error
	err := validator.ValidationErrors{}

	// Call the notFoundError function
	result := notFoundError(rr, req, err)

	// Check that the function returns true
	if !result {
		t.Errorf("Expected notFoundError to return true, but it returned false")
	}

	// Check the HTTP status code
	if rr.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, rr.Code)
	}

	// Check the HTTP response body
	expected := `{"code":404,"status":"Not Found","data":""}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Expected body %s, but got %s", expected, rr.Body.String())
	}

	// // Create a new validation error
	// err = validator.ValidationErrors{}

	// // Call the notFoundError function
	// result := notFoundError(rr, req, err)

	// // Check that the function returns true
	// if !result {
	// 	t.Errorf("Expected notFoundError to return true, but it returned false")
	// }
}

func Test_validationError(t *testing.T) {
	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/", nil)

	// Create a new validation error
	err := validator.ValidationErrors{}

	// Call the validationError function
	result := validationError(rr, req, err)

	// Check that the function returns true
	if !result {
		t.Errorf("Expected validationError to return true, but it returned false")
	}

	// Check the HTTP status code
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, but got %d", http.StatusBadRequest, rr.Code)
	}

	// Check the HTTP response body
	expected := `{"code":400,"status":"BAD REQUEST","data":""}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Expected body %s, but got %s", expected, rr.Body.String())
	}
}

func Test_internalServerError(t *testing.T) {
	// Create a new HTTP response recorder
	rr := httptest.NewRecorder()

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/", nil)

	// Create a new error
	err := errors.New("something went wrong")

	// Call the internalServerError function
	internalServerError(rr, req, err)

	// Check the HTTP status code
	if rr.Code != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, rr.Code)
	}

	// Check the HTTP response body
	expected := `{"code":500,"status":"INTERNAL SERVER ERROR","data":{}}`
	if strings.TrimRight(rr.Body.String(), "\n") != expected {
		t.Errorf("Expected body %s, but got %s", expected, rr.Body.String())
	}
}

func TestErrorHanlder(t *testing.T) {
	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/", nil)

	ErrorHanlder(writer, request, "Test error")

	response := writer.Result()
	defer response.Body.Close()

	if response.StatusCode != http.StatusInternalServerError {
		t.Errorf("Expected status code %d, but got %d", http.StatusInternalServerError, response.StatusCode)
	}

	contentType := response.Header.Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected content type application/json, but got %s", contentType)
	}

	var body struct {
		Code   int    `json:"code"`
		Status string `json:"status"`
		Data   string `json:"data"`
	}

	err := json.NewDecoder(response.Body).Decode(&body)
	if err != nil {
		t.Errorf("Error decoding response body: %s", err.Error())
	}

	if body.Code != http.StatusInternalServerError {
		t.Errorf("Expected response code %d, but got %d", http.StatusInternalServerError, body.Code)
	}

	if body.Status != "INTERNAL SERVER ERROR" {
		t.Errorf("Expected response status 'INTERNAL SERVER ERROR', but got '%s'", body.Status)
	}

	if body.Data != "Test error" {
		t.Errorf("Expected response data 'Test error', but got '%s'", body.Data)
	}
}
