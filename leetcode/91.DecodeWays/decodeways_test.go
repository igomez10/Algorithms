package main

import "testing"

func TestDecodeWays(t *testing.T) {
	cases := map[string]int{
		"12":  2,
		"226": 3,
		"0":   0,
		"1":   1,
	}

	for k, v := range cases {
		t.Run(k, func(t *testing.T) {
			res := numDecodings(k)
			if res != v {
				t.Error("got", numDecodings(k), "expected", v)
			}
		})
	}
}
