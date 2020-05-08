package main

import (
	"sort"
	"testing"
)

var arr []int = []int{0, 1, 2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 5, 6, 7, 8, 9, 10, 100, 200, 300, 400, 555, 600, 700, 800, 900, 1000}
var searchedElement int = 1005

func TestIsInArray(t *testing.T) {

	if index := binarySearch(arr, 4); index != 4 {
		t.Errorf("Wrong index found %d != %d", index, 4)
	}

	if binarySearch(arr, searchedElement) != -1 {
		t.Error("found unexisting element", searchedElement, arr)
	}
}

func TestSortPackage(t *testing.T) {
	itemToSearch := 4

	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= itemToSearch
	})

	if index >= len(arr) || arr[index] != itemToSearch {
		t.Fatal("item not in array")
	}

	if index < len(arr) && arr[index] == 4 {
		t.Logf("Success found %d at index %d in %v\n", 4, index, arr)
	} else {
		t.Errorf("%d not found in %v\n", index, arr)
	}

}
