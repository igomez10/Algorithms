package twosum

import (
	"testing"
)

func twoSum(nums []int, target int) []int {
	// wantedValue, complementIndex
	wantedNumbers := map[int]int{}
	for i := range nums {
		complement := target - nums[i]
		if firstSeenAt, exists := wantedNumbers[complement]; exists {
			return []int{firstSeenAt, i}
		}

		wantedNumbers[nums[i]] = i
	}

	return []int{}
}

func Test_twoSum(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Example 1",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Example 2",
			nums:   []int{3, 2, 4},
			target: 6,
			want:   []int{1, 2},
		},
		{
			name:   "Example 3",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if len(got) != len(tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
				return
			}
			for i := range got {
				if got[i] != tt.want[i] {
					t.Errorf("twoSum() = %v, want %v", got, tt.want)
					return
				}
			}
		})
	}
}
