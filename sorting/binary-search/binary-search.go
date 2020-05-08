package main

func binarySearch(haystack []int, needle int) int {
	lo := 0
	hi := len(haystack) - 1

	//it's ok if the indexes are equal, after the next cycle, lo|hi will move +-1, preventing constant values as lo==hi
	for lo < hi {
		middle := int(uint(hi+lo) >> 1)

		if haystack[middle] < needle {
			lo = middle + 1 // why include the middle element if we already know that it's not our searchedElement??, skip it with a +1
		} else {
			hi = middle
		}
	}

	if haystack[lo] != needle {
		return -1
	}

	return lo

}
