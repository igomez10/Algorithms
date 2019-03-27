package main

import "testing"

func TestURLify(t *testing.T) {
	initial := "Mr John Smith"
	expected := "Mr%20John%20Smith"
	if URLify(initial) != expected {
		t.Errorf("Expected %s got %s", expected, URLify(initial))
	}
}
