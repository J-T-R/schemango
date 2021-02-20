package main

import "testing"

func TestAddress(t *testing.T) {
	a := Address{"test", "test_again", 99}
	postString := a.createPostString()
	expectedString := "test://test_again:99"

	if postString != expectedString {
		t.Errorf("Expected: %s, got: %s", postString, expectedString)
	}
}
