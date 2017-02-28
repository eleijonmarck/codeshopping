package api

import (
	"testing"
)

func TestGetCart(t *testing.T) {
	expectedResult := "hej"
	actualResult := ""
	if actualResult != "" {
		t.Fatalf("Expected %s but got the acutal %s", expectedResult, actualResult)
	}
}
