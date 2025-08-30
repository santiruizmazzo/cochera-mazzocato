package utils

import "testing"

func AssertResponseContains[C comparable](response map[string]any, expectedKey string, expectedContent C, t *testing.T) {
	if receivedContent, ok := response[expectedKey]; !ok || receivedContent != expectedContent {
		t.Fatalf("Expected %s %v, got %v", expectedKey, expectedContent, receivedContent)
	}
}

func AssertStatusCodeIs(expectedCode, receivedCode int, t *testing.T) {
	if expectedCode != receivedCode {
		t.Fatalf("Expected status code %d, got %d", expectedCode, receivedCode)
	}
}
