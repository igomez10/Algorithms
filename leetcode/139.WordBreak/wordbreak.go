package main

func wordBreak(s string, dict []string) bool {
	memo := &map[int]bool{}
	return wordBreakMemo(s, dict, 0, memo)
}

func wordBreakMemo(s string, dict []string, index int, memo *map[int]bool) bool {
	if index == len(s) {
		return true
	}
	if val, exist := (*memo)[index]; exist {
		return val
	}

	for i := range dict {
		if len(dict[i]) <= len(s)-index &&
			s[index:index+len(dict[i])] == dict[i] &&
			wordBreakMemo(s, dict, index+len(dict[i]), memo) {
			(*memo)[index] = true
			return true
		}
	}

	(*memo)[index] = false
	return false
}
