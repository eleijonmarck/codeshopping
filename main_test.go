package main

import "testing"

func TeststoreData(t *testing.T) {
	actualResult := "hej"
	var expectedResult = "hej"

	if actualResult != expectedResult {
		t.Fatalf("Expected %s but got the acutal %s", expectedResult, actualResult)
	}
}
