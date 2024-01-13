package auth

import (
	"net/http/httptest"
	"testing"
)

func TestGetApiKeyValid(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8000/", nil)
	req.Header.Set("Authorization", "ApiKey valid")
	input := req.Header
	want := "valid"
	got, _ := GetAPIKey(input)
	if want != got {
		t.Fatalf("expected: %s, got: %s", want, got)
	}
}

func TestGetApiKeyInvalid(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8000/", nil)
	req.Header.Set("Authorization", "DiferentKey valid")
	input := req.Header
	want := "a"
	got, _ := GetAPIKey(input)
	if want != got {
		t.Fatalf("expected: %s, got: %s", want, got)
	}
}
