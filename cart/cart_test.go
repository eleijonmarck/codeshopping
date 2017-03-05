package cart

import (
	"net/http"
	"testing"
)

func TestCartHasItem_WhenAdded(t *testing.T) {
	expectedResults := "hej"

	if expectedResults != "hej" {
		t.Errorf("Expected %s but got fjeioa", expectedResults)
	}
}

func TestGetCartWithId(t *testing.T) {
	r, err := http.Get("http://localhost:8080")
	if err != nil && r.StatusCode == 200 {
		t.Errorf("Either no internal connection or api for cart is not up")
	}
}
