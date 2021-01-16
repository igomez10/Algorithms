package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	low := 0
	high := len(nums1)

	for low <= high {
		partitionA := (low + high) / 2
		partitionB := ((len(nums1) + len(nums2) + 1) / 2) - partitionA

		maxLeftA := -math.MaxInt32
		if partitionA != 0 {
			maxLeftA = nums1[partitionA-1]
		}

		minRightA := math.MaxInt32
		if partitionA != len(nums1) {
			minRightA = nums1[partitionA]
		}

		maxLeftB := -math.MaxInt32
		if partitionB != 0 {
			maxLeftB = nums2[partitionB-1]
		}

		minRightB := math.MaxInt32
		if partitionB != len(nums2) {
			minRightB = nums2[partitionB]
		}

		if maxLeftA <= minRightB && maxLeftB <= minRightA {
			if (len(nums1)+len(nums2))%2 == 0 {
				return float64(max(maxLeftA, maxLeftB)+min(minRightA, minRightB)) / 2.0
			} else {
				return float64(max(maxLeftA, maxLeftB))
			}
		} else if maxLeftA > minRightB {
			high = partitionA - 1
		} else {
			low = partitionA + 1
		}
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
