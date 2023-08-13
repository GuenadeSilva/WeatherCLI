package api

import (
	"testing"
)

func TestApiCall(t *testing.T) {
	// Assuming you have set up the necessary environment variables
	datasets := ApiCall()

	// Test that the results are not empty
	if len(datasets) == 0 {
		t.Error("Expected non-empty datasets, but got empty")
	}

	// Add more specific tests if needed
}
