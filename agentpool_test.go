package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAgentPoolSuccess(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			w.WriteHeader(http.StatusOK)
			return
		}
		if r.Method == http.MethodGet {
			w.WriteHeader(http.StatusConflict)
			return
		}
	}))
	defer ts.Close()

	// override the global variable with the test server client
	client = ts.Client()

	pool := Pools{
		Name:          poolName,
		AutoProvision: true,
		IsHosted:      false,
	}

	err := createAgentPool(pool)
	if err != nil {
		if err.Error() != "error: 409 Conflict" {
			t.Errorf("Unexpected error: %v", err)
		}
	}
}

func TestCreateAgentPoolInvalidInput(t *testing.T) {
	// create an invalid agent pool with an empty name
	pool := Pools{
		Name:          "",
		AutoProvision: true,
		IsHosted:      false,
	}

	err := createAgentPool(pool)
	if err == nil {
		t.Errorf("Expected error for invalid input, got nil")
	}

	// check that the error message is as expected
	expectedError := "error: Invalid input: pool name cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}
