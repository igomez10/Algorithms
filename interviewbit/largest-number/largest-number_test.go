package main

import (
	"strings"
	"testing"
)

func TestLargestNumber(t *testing.T) {
	numbersToUse := []int{3, 30, 34, 5, 9}
	expectedResult := "9534330"

	result := findLargest(numbersToUse)
	if strings.Compare(result, expectedResult) != 0 {
		t.Errorf("Unexpected largest number, expected %s - got %s", expectedResult, result)
	}
}

func TestCompareByFindLargest(t *testing.T) {
	arr := arrLargest{9, 91, 123, 991, 992}
	if arr.Less(0, 1) != false {
		t.Errorf("%d is bigger than %d but was reported otherwise", arr[0], arr[1])
	}

	if arr.Less(1, 2) != false {
		t.Errorf("%d is bigger than %d but was reported otherwise", arr[1], arr[2])
	}

	if arr.Less(3, 4) != true {
		t.Errorf("%d is smaller than %d but was reported otherwise", arr[3], arr[4])
	}

}
