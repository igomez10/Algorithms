package main

func main() {}

func isRotation(s1, s2 string) bool {
	longer := s1
	smaller := s2
	if longer < smaller {
		smaller, longer = longer, smaller
	}

	combined := longer + longer

	if isSubstring(smaller, combined) {
		return true
	}
	return false

}

func isSubstring(s1, s2 string) bool {

	longest := s1
	smallest := s2
	if len(s2) > len(s1) {
		longest = s2
		smallest = s1
	}

	var counter int
	for i := 0; i <= len(longest)-len(smallest); i++ {
		counter = 0
		for j := 0; j < len(smallest); j++ {
			longchar := longest[i+j]
			smallchar := smallest[j]
			if longchar == smallchar {
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
