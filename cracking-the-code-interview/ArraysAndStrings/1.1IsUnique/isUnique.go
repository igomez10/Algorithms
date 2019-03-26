package main

func IsUnique(S string) bool {
	uniques := make([]bool, 128) //assuming ASCII encoding
	for i := range S {
		if uniques[S[i]] {
			return false
		}
		uniques[S[i]] = true
	}
	return true
}
