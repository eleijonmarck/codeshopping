package cart

import (
	"testing"
)

func TestCartHasItem_WhenAdded(t *testing.T) {
	expectedResults := "hej"

	if expectedResults != "hej" {
		t.Errorf("Expected %s but got fjeioa")
	}
}
