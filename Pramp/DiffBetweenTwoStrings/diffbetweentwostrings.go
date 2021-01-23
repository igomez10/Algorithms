package main

import "fmt"

func DiffStrings(source, target string) []string {
	chain := LCS(source, target)
	return chain
}

func LCS(w1, w2 string) []string {
	if len(w1) == 0 {
		res := []string{}
		for i := range w2 {
			res = append(res, "+"+string(w2[i]))
		}
		return res
	} else if len(w2) == 0 {
		res := []string{}
		for i := range w1 {
			res = append(res, "-"+string(w1[i]))
		}
		return res
	}

	if w1[0] == w2[0] {
		res := LCS(w1[1:], w2[1:])
		return append([]string{string(w1[0])}, res...)
	}
	ans1 := LCS(w1[1:], w2)
	ans2 := LCS(w1, w2[1:])

	if len(ans1) <= len(ans2) {
		return append([]string{"-" + string(w1[0])}, ans1...)
	}

	return append([]string{"+" + string(w2[0])}, ans2...)
}

func main() {
	source := "ABCDEFG"
	target := "ABDFFGH"
	//  expected := []string{"A", "B", "-C", "D", "-E", "F", "+F", "G", "+H"}
	//  fmt.Println(expected)
	fmt.Println(DiffStrings(source, target))

}
