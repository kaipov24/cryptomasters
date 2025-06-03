package api_test

import (
	"testing"

	"github.com/kaipov24/cryptomsters/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error is not found")
	}
}
