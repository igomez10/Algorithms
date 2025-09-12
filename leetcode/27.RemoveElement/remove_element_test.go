package removeelement

import "testing"

func removeElement(nums []int, val int) int {
	// find first non val element from the right to left and create indexRight = i
	// traverse nums left to right, if find a nums[i] = val, swap with nums[indexRight]
	// move indexRight-- until nums[indexRight] is not a val

	//  find first non val from right to left
	nonValsCounter := 0
	indexRight := len(nums) - 1
	for i := len(nums) - 1; i >= 0 && nums[indexRight] == val && indexRight >= 0; i-- {
		indexRight--
		nonValsCounter++
	}

	for i := 0; i < len(nums) && i < indexRight; i++ {
		if nums[i] == val {
			// swap i with indexRight
			nums[i], nums[indexRight] = nums[indexRight], nums[i]
			// move indexRight until next non val
			for indexRight >= 0 && nums[indexRight] == val {
				indexRight--
				nonValsCounter++
			}
		}
	}

	return len(nums) - nonValsCounter
}
func Test_removeElement(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		val  int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 2, 2, 3},
			val:  3,
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
			val:  2,
			want: 5,
		},
		{
			name: "Empty",
			nums: []int{},
			val:  0,
			want: 0,
		},
		{
			name: "No elements to remove",
			nums: []int{1, 2, 3, 4, 5},
			val:  6,
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElement(tt.nums, tt.val)
			if got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
