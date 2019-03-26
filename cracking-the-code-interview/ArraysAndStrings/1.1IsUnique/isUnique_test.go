package main

import "testing"

func TestIsUnique(t *testing.T) {
	unique := "abcdefghi"
	if !IsUnique(unique) {
		t.Error("Expecting true got false")
	}
	notUnique := "abcdaaefghi"
	if IsUnique(notUnique) {
		t.Error("Expecting false got true")
	}
}
