package main

import (
	"container/list"
)

func wordBreak(s string, dict []string) bool {
	l := list.New()
	l.PushBack(0)
	memo := map[int]bool{}
	for l.Len() > 0 {
		index := l.Remove(l.Front()).(int)
		if index == len(s) {
			return true
		}

		if _, visited := memo[index]; visited {
			continue
		} else {
			memo[index] = visited
		}

		for i := range dict {
			if len(dict[i]) <= len(s)-index &&
				s[index:index+len(dict[i])] == dict[i] {
				l.PushBack(index + len(dict[i]))
			}
		}
	}

	return false
}
