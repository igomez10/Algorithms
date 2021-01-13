package main

import (
	"container/list"
	"sort"
)

func wordBreak(s string, dict []string) bool {
	sort.Slice(dict, func(i, j int) bool {
		return len(dict[i]) < len(dict[j])
	})

	l := list.New()
	l.PushBack(0)
	for l.Len() > 0 {
		index := l.Remove(l.Front()).(int)
		for i := range dict {

			if len(dict[i]) <= len(s)-index &&
				s[index:index+len(dict[i])] == dict[i] {

				if len(dict[i])+index == len(s) {
					return true
				}
				l.PushBack(index + len(dict[i]))

			}
		}
	}

	return false
}
