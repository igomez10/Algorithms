package main

import (
	"sort"
)

type sortableString []rune

func (s sortableString) Len() int {
	return len(s)
}

func (s sortableString) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortableString) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func checkPermutation(A, B string) bool {
	if len(A) != len(B) {
		return false
	}
	rA := []rune(A)
	rB := []rune(B)
	sort.Sort(sortableString(rA))
	sort.Sort(sortableString(rB))
	for i := 0; i < len(rA); i++ {
		if rA[i] != rB[i] {
			return false
		}
	}
	return true
}
