package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	exponent := 16
	preppendToArray(exponent)
	appendToArray(exponent)
	preppendingWithoutArrayAlocation(exponent)
	appendingWithoutArrayAlocation(exponent)
}

func preppendToArray(n int) {
	counter := 1
	for i := 2; float64(i) < math.Pow(float64(2), float64(n)); i *= 2 {
		start := time.Now()
		arr := make([]int, 0, i)
		fmt.Printf("preppending 2^%d elements\n", counter)
		for j := 0; j < i; j++ {
			arr = append([]int{i}, arr...)
		}
		fmt.Printf("\t%d took %+v\n", i, time.Since(start))
		counter++
	}
}

func appendToArray(n int) {
	counter := 1
	for i := 2; float64(i) < math.Pow(float64(2), float64(n)); i *= 2 {
		start := time.Now()
		arr := make([]int, 0, i)
		fmt.Printf("appending 2^%d elements\n", counter)
		for j := 0; j < i; j++ {
			arr = append(arr, i)
		}
		fmt.Printf("\t%d took %+v\n", i, time.Since(start))
		counter++
	}
}

func preppendingWithoutArrayAlocation(n int) {
	counter := 1
	for i := 2; float64(i) < math.Pow(float64(2), float64(n)); i *= 2 {
		start := time.Now()
		arr := []int{} // KEY difference, capacity not allocated
		fmt.Printf("preppending 2^%d elements\n", counter)
		for j := 0; j < i; j++ {
			arr = append([]int{i}, arr...)
		}
		fmt.Printf("\t%d took %+v\n", i, time.Since(start))
		counter++
	}
}

func appendingWithoutArrayAlocation(n int) {
	counter := 1
	for i := 2; float64(i) < math.Pow(float64(2), float64(n)); i *= 2 {
		start := time.Now()
		arr := []int{} // KEY difference, capacity not allocated
		fmt.Printf("preppending 2^%d elements\n", counter)
		for j := 0; j < i; j++ {
			arr = append(arr, i)
		}
		fmt.Printf("\t%d took %+v\n", i, time.Since(start))
		counter++
	}
}
