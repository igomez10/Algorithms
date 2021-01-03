package main

import "testing"

var True = true
var False = false

func TestLoggerRateLimiter(t *testing.T) {
	events := []string{
		"Logger",
		"shouldPrintMessage",
		"shouldPrintMessage",
		"shouldPrintMessage",
		"shouldPrintMessage",
		"shouldPrintMessage",
		"shouldPrintMessage",
	}
	logs := []Log{
		{},
		{1, "foo"},
		{2, "bar"},
		{3, "foo"},
		{8, "bar"},
		{10, "foo"},
		{11, "foo"},
	}

	ans := Answer(events, logs)

	expected := []*bool{nil, &True, &True, &False, &False, &False, &True}

	for i := 1; i < len(expected); i++ {
		currentAns := *ans[i]
		currentExpected := *expected[i]
		if currentAns != currentExpected {
			t.Errorf("index %d expected %t but got %t", i, *expected[i], *ans[i])
		}
	}

}
