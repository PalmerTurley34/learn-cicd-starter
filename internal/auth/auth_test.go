package auth

import (
	"net/http/httptest"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	_, err := GetAPIKey(req.Header)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("No fail for missing auth header")
	}
	req.Header.Add("Authorization", "123")
	_, err = GetAPIKey(req.Header)
	if err == nil {
		t.Fatalf("No fail for bad auth header")
	}
	if err.Error() != "malformed authorization header" {
		t.Fatalf("Unexpected error for bad auth header")
	}
	req.Header.Set("Authorization", "ApiKey 123")
	apiKey, err := GetAPIKey(req.Header)
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}
	if apiKey != "123" {
		t.Fatalf("Expected apiKey 123. Got: %s", apiKey)
	}
}
