package main

func main() {}

func CountInversions(arr []int) int {

	if len(arr) <= 2 {
		return -1
	}
	counter := 0
	biggerTrend := true
	if arr[0] > arr[1] {
		biggerTrend = !biggerTrend
	}

	for i := 2; i < len(arr); i++ {
		if arr[i] >= arr[i-1] && biggerTrend {

		} else if arr[i] <= arr[i-1] && !biggerTrend {

		} else {
			biggerTrend = !biggerTrend
			counter++
		}
	}
	return counter
}

func CountInversionsNaive(inputArray []int) int {
	counter := 0
	for i := 0; i < len(inputArray); i++ {
		for j := i; j < len(inputArray); j++ {
			if inputArray[i] > inputArray[j] {
				counter++
			}
		}
	}
	return counter
}
