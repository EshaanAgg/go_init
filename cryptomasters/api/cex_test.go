package api_test

import (
	"eshaanagg/go/cryptomasters/api"
	"testing"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Errorf("Error was not found.")
	}
}
