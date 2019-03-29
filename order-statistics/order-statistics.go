package main

import "fmt"

// Interview : https://www.youtube.com/watch?v=D35llNtkCps
// Given an unrordered set of numbers, return the Nth smallest ELEMENT IN THE LIST IF IT WERE SORTED
// it becomes on**2 if i want to get the n element
// solution: sorting the array and get the n element
// 1 6 3 9 8 5

// ---------------- Implementing QuickSelect ---------------

/**
Maybe expanding and creating a list of all the mins up to n,
the problem with this is the big o notation would be: storing them all and checking all of them:
the runtime would become on**2, (checking all them )

SPECIAL QUESTION: THERE ARE NO DUPLICATES!!!

we can do nlogn if we sort the array
Can we do bewtter than n*log(n) ??
if we have 2,3,5,6 and m = 3

start ordering a random element, if its final position is == m then this is our desired value,
if the final indfex for the value is smaller than m then our value is to the right , if it's bigger then it's to the left


**/
func main() {
	arr := []int{73, 90, 85, 58, 69, 77, 90, 85, 18, 35}
	nthNumber := 0
	lo, hi := 0, len(arr)-1
	var answer int
	for {
		finalIndexForFirst := partition(&arr, lo, hi)
		fmt.Println(arr)
		if finalIndexForFirst < nthNumber {
			lo = finalIndexForFirst + 1
		} else if finalIndexForFirst > nthNumber {
			hi = finalIndexForFirst - 1
		} else {
			answer = arr[finalIndexForFirst]
			break
		}
	}
	fmt.Printf("The %d smallest number is %d\n", nthNumber, answer)
	fmt.Printf(" %+v\n", arr)
}

//quickselect modifies the inital array where the first element
//of the array arr[0] is in it's final postion and return the final
//index of the initial element
// we start by picking a[0] as the partitioning element and a pointer i, j
// increment i as long as its poingting to smething smaller than k
// decrement j as long as its pointing to somehting biger than k
// when both stop, switch a[i] and a[j]

func partition(arr *[]int, lo, hi int) int {
	pivot := hi
	i := lo - 1
	for j := lo; j <= hi-1; j++ {
		if (*arr)[j] <= (*arr)[pivot] {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}
	(*arr)[i+1], (*arr)[hi] = (*arr)[hi], (*arr)[i+1]
	return i + 1
}

func findSmallest(arr []int, n int) int {
	smallest := arr[0]
	var secondSmallest int
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			secondSmallest = smallest
			smallest = arr[i]
		} else if arr[i] < secondSmallest {
			secondSmallest = arr[i]
		}
	}
	return secondSmallest
}

func testingArr(arr *[]int) {
	(*arr)[0] = 9999
}
