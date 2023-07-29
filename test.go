package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	// Create a request to the /hello endpoint
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which implements http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Call the handler function
	helloHandler(rr, req)

	// Check the status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status %d; got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedResponse := "hello!"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Expected response body '%s'; got '%s'", expectedResponse, rr.Body.String())
	}
}
