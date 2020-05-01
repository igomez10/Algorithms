package main

import (
	"strings"
	"testing"
)

func TestClassifyStrings(t *testing.T) {
	badWords := []string{
		"aaaaa",
		"abaaa",
		"ababaeee",
	}

	goodWords := []string{
		"cxcvbxcvb",
		"sadlfkjhsaldf",
		"qwrtpc",
	}

	for i := range badWords {
		if strings.Compare(classifyStrings(badWords[i]), "bad") != 0 {
			t.Errorf("Expected bad word, got %s", classifyStrings(badWords[i]))
		}
	}

	for i := range goodWords {
		if strings.Compare(classifyStrings(goodWords[i]), "good") != 0 {
			t.Errorf("Expected bad word, got %s", classifyStrings(badWords[i]))
		}
	}

	for i := range goodWords {
		if strings.Compare(classifyStrings(goodWords[i]), "good") != 0 {
			t.Errorf("Expected bad word, got %s", classifyStrings(badWords[i]))
		}
	}

}
