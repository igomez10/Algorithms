package main

import (
	"fmt"
	"sort"
)

type Indexes struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

type Triples struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
	Z int `json:"z,omitempty"`
}

func threeSum(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

	answer := [][]int{}

	for i := 0; i < len(nums)-2; i++ {
		if i == 0 || (i > 0 && nums[i] != nums[i-1]) {
			low := i + 1
			high := len(nums) - 1

			for low < high {
				fmt.Println(low, high)
				if nums[low]+nums[high]+nums[i] == 0 {
					answer = append(answer, []int{nums[low], nums[high], nums[i]})

					for low < high && nums[low] == nums[low+1] {
						low++
					}
					for low < high && nums[high] == nums[high-1] {
						high--
					}

					low++
					high--

				} else if nums[low]+nums[high]+nums[i] < 0 {
					low++
				} else {
					high--
				}

			}
		}
	}
	return answer
}

func main() {
	fmt.Println("--------START-----------")
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum(nums))
	fmt.Println("--------END-----------")
}

func printMap(m map[Indexes]int) {
	for k, v := range m {
		fmt.Println(fmt.Sprintf("%+v, %+v\n", k, v))
	}
}

func printMapEntry(m map[Indexes]int, key Indexes) {
	fmt.Println(fmt.Sprintf("%+v, %+v\n", key, m[key]))
}
