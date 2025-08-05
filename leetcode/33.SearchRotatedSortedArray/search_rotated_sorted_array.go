package searchrotatedsortedarray

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	} else if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	lo := 0
	hi := len(nums) - 1
	middle := (lo + hi) / 2
	if nums[lo] < nums[hi] {
		middle = 0
	} else {
		for lo <= hi {
			middle = (lo + hi) / 2
			if nums[middle] > nums[middle+1] {
				middle = middle + 1
				break
			} else {
				if nums[middle] < nums[lo] {
					hi = middle - 1
				} else {
					lo = middle + 1
				}
			}
		}
	}

	lo = 0
	hi = middle - 1
	for lo <= hi {
		middle := (lo + hi) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			lo = middle + 1
			continue
		}
		hi = middle - 1
	}

	lo = middle
	hi = len(nums) - 1
	for lo <= hi {
		middle := (lo + hi) / 2
		if nums[middle] == target {
			return middle
		}
		if nums[middle] < target {
			lo = middle + 1
			continue
		}
		hi = middle - 1
	}
	return -1
}
