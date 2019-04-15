package main

import "fmt"

func main() {

}

// StringRotation returns true is s2 if by rotation the characters s1 can be obtained
func StringRotation(s1, s2 string) {
	//	longed := s1 + s1
}

func isSubstring(s1, s2 string) bool {

	longest := s1
	smallest := s2
	if len(s2) > len(s1) {
		longest = s2
		smallest = s1
	}

	var counter int
	for i := range longest {
		counter = 0
		for j := 0; j < len(smallest); j++ {
			fmt.Println(longest, smallest, j, i)
			if longest[i+j] == smallest[j] {
				counter++
			} else {
				break
			}
			if counter == len(smallest) {
				return true
			}
		}
	}
	return false
}
