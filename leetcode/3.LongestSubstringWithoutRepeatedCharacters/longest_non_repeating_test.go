package longestsubstringwithoutrepeatedcharacters

import "testing"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	lo := 0
	hi := 0

	seen := make([]bool, 256)
	maxSoFar := 1
	// traverse left to right using lo but not always increasing lo, sometimes only hi is increased
	for lo < len(s) {
		if hi < len(s) && seen[int(s[hi])] == false {
			// s in hi is new char, add it
			seen[int(s[hi])] = true
			currentLength := hi - lo + 1
			if maxSoFar < currentLength {
				maxSoFar = currentLength
			}

			hi++
			continue
		}

		// s in hi is not new, remove s in lo and increase lo by 1
		seen[int(s[lo])] = false
		lo++
	}

	return maxSoFar
}

func TestLengthLongest(t *testing.T) {

	tests := []struct {
		name string
		s    string
		want int
	}{
		{
			name: "Example 1",
			s:    "abcabcbb",
			want: 3,
		},
		{
			name: "Example 2",
			s:    "bbbbb",
			want: 1,
		},
		{
			name: "Example 3",
			s:    "pwwkew",
			want: 3,
		},
		{
			name: "Example 4",
			s:    "",
			want: 0,
		},
		{
			name: "Example 5",
			s:    "au",
			want: 2,
		},
		{
			name: "Example 6",
			s:    "aab",
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
