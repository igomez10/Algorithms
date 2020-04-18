package main

import (
	"bytes"
	"sort"
)

type sortString []byte

func isAnagram(s string, t string) bool {
	s1 := sortString(s)
	s2 := sortString(t)

	sort.Sort(s1)
	sort.Sort(s2)

	if !bytes.Equal(s1, s2) {
		return false
	}
	return true
}

func (s sortString) Less(i, j int) bool {
	if s[i] < s[j] {
		return true
	}

	return false
}

func (s sortString) Swap(i, j int) {
	tmp := s[i]
	s[i] = s[j]
	s[j] = tmp
}

func (s sortString) Len() int {
	return len(s)
}
