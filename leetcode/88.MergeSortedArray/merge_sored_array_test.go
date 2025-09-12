package mergesortedarray

import "testing"

func merge(nums1 []int, m int, nums2 []int, n int) {
	// return in new array
	// keep two indexes
	// check current value in both indexes, pick the smallest and move index by 1
	// iterate m+n times
	// idx1 should only reach to len(nums1)
	// idx2 should only reach to len(nums2)
	// if by end of iteration idx1 < len(nums1)-1 or idx2 < len(nums2)-1
	// append correspoinding missingnnumbers to the end
	if m == 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}

	idxTotal := 0
	idx1 := m - 1
	idx2 := n - 1

	for idx1 >= 0 && idx2 >= 0 {
		if nums1[idx1] > nums2[idx2] {
			nums1[len(nums1)-1-idxTotal] = nums1[idx1]
			idx1--
			idxTotal++
		} else {
			nums1[len(nums1)-1-idxTotal] = nums2[idx2]
			idx2--
			idxTotal++
		}
	}

	for idx1 >= 0 {
		nums1[len(nums1)-1-idxTotal] = nums1[idx1]
		idx1--
		idxTotal++
	}

	for idx2 >= 0 {
		nums1[len(nums1)-1-idxTotal] = nums2[idx2]
		idx2--
		idxTotal++
	}
}

func TestMerge(t *testing.T) {
	tests := []struct {
		name         string
		nums1        []int
		m            int
		nums2        []int
		n            int
		expectedNums []int
	}{
		{
			name:         "Example 1",
			nums1:        []int{1, 2, 3, 0, 0, 0},
			m:            3,
			nums2:        []int{2, 5, 6},
			n:            3,
			expectedNums: []int{1, 2, 2, 3, 5, 6},
		},
		{
			name:         "Example 2",
			nums1:        []int{1},
			m:            1,
			nums2:        []int{},
			n:            0,
			expectedNums: []int{1},
		},
		{
			name:         "Example 3",
			nums1:        []int{0},
			m:            0,
			nums2:        []int{1},
			n:            1,
			expectedNums: []int{1},
		},
	}
	for _, tt := range tests {
		merge(tt.nums1, tt.m, tt.nums2, tt.n)
		for i := range tt.nums1 {
			if tt.nums1[i] != tt.expectedNums[i] {
				panic("fail")
			}
		}
	}
}
