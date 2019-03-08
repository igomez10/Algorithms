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
	arr := []int{8, 11, 6, 5, 4, 10, 2, 3, 1, 9}
	nthNumber := 3
	lo, hi := 0, len(arr)-1
	finalIndexForFirst := partition(&arr, lo, hi)
	fmt.Println(finalIndexForFirst)
	var answer int
	for {
		if finalIndexForFirst < nthNumber {
			lo = finalIndexForFirst + 1
			finalIndexForFirst = partition(&arr, lo, hi)
		} else if finalIndexForFirst > nthNumber {
			hi = finalIndexForFirst - 1
			finalIndexForFirst = partition(&arr, lo, hi)
		} else {
			answer = arr[finalIndexForFirst]
			break
		}
		fmt.Println(finalIndexForFirst)
	}
	fmt.Println(finalIndexForFirst)
	fmt.Printf("The %d smallest number is %d", nthNumber, answer)
	fmt.Printf(" %+v", arr)
}

//quickselect modifies the inital array where the first element
//of the array arr[0] is in it's final postion and return the final
//index of the initial element
// we start by picking a[0] as the partitioning element and a pointer i, j
// increment i as long as its poingting to smething smaller than k
// decrement j as long as its pointing to somehting biger than k
// when both stop, switch a[i] and a[j]

func partition(arr *[]int, lo, hi int) int {
	pivot := lo
	smallPointer := lo + 1
	bigPointer := hi
	for bigPointer > smallPointer {
		for (*arr)[smallPointer] < (*arr)[pivot] && smallPointer <= hi {
			smallPointer++
		}

		for (*arr)[bigPointer] > (*arr)[pivot] && bigPointer >= lo {
			bigPointer--
		}

		if bigPointer <= smallPointer {
			break
		}

		(*arr)[smallPointer], (*arr)[bigPointer] = (*arr)[bigPointer], (*arr)[smallPointer]
	}
	(*arr)[bigPointer], (*arr)[pivot] = (*arr)[pivot], (*arr)[bigPointer]
	return bigPointer
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
