package main

import "strings"

func CompareStrings(A string, B string) []int {
	arrA := strings.Split(A, ",")
	arrB := strings.Split(B, ",")

	ans := make([]int, len(arrB))
	for i := range arrB {
		for j := range arrA {
			count1 := GetFrequency(arrA[j], GetSmallestChar(arrA[j]))
			count2 := GetFrequency(arrB[i], GetSmallestChar(arrB[i]))
			if count1 < count2 {
				ans[i]++
			}

		}
	}
	return ans
}

func GetSmallestChar(S string) string {
	if len(S) <= 1 {
		return S
	}

	var smallest rune
	for _, v := range S {
		smallest = v
	}

	for _, v := range S {
		if v < smallest {
			smallest = v
		}
	}

	return string(smallest)
}

func GetFrequency(S string, char string) int {
	if char == "" || len(S) == 0 {
		return 0
	}
	var r rune
	for _, c := range char {
		r = c
	}

	counter := 0
	for _, v := range S {
		if v == r {
			counter++
		}
	}
	return counter
}

// sm1 := GetSmallestChar(A)
// count1 := CountChar(sm1, A)

// sm2 := GetSmallestChar(B)
// count2 := CountChar(sm2, B)

// return count1 < count2
