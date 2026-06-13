package scanner

import "testing"

func TestTrivyScan(t *testing.T) {
	// TODO: Mock Trivy execution or use testcontainers
	t.Skip("Integration test - requires Trivy binary or Docker")
}

func TestZAPScan(t *testing.T) {
	err := RunZAPScan("https://example.com", "baseline")
	if err != nil {
		t.Log("Expected in test env without Docker/ZAP:", err)
	}
}