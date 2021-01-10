package main

import (
	"testing"
)

func TestCompareStrings(t *testing.T) {

	type TestCase struct {
		Name     string
		A        string
		B        string
		Expected []int
	}

	cases := []TestCase{
		{
			Name:     "1",
			A:        "abcd,aabc,bd",
			B:        "aaa,aa",
			Expected: []int{3, 2},
		},
	}

	for i := range cases {
		t.Run(cases[i].Name, func(t *testing.T) {
			res := CompareStrings(cases[i].A, cases[i].B)
			if len(res) != len(cases[i].Expected) {
				t.Fatal("unexpected length", len(res), cases[i].Expected)
			}

			for j := range res {
				if res[j] != cases[i].Expected[j] {
					t.Error("unexpected value ", res[i], cases[i].Expected[j])
					t.Log(res)
					t.Log(cases[i].Expected)
				}
			}
		})
	}
}
