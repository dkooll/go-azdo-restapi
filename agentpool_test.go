package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCreateAgentPoolBadRequest(t *testing.T) {
	// create a mock server that returns a non-200 status code
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
	}))
	defer mockServer.Close()

	poolname := Pools{Name: "Selfhosted"}
	poolBytes, err := json.Marshal(poolname)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", mockServer.URL, bytes.NewBuffer(poolBytes))
	if err != nil {
		t.Error(err)
	}

	req.Header.Add("Authorization", "Basic "+os.Getenv("AZURE_TOKEN"))
	req.Header.Add("Content-Type", "application/json")

	err = createAgentPool(poolname)

	if err == nil {
		t.Error("Expected error, got nil")
	}
}

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
	pool := Pools{
		Name:          "",
		AutoProvision: true,
		IsHosted:      false,
	}

	err := createAgentPool(pool)
	if err == nil {
		t.Errorf("Expected error for invalid input, got nil")
	}

	expectedError := "error: Invalid input: pool name cannot be empty"
	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}
}
