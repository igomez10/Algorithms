package main

import "fmt"

func FindBusiestPeriod(data [][]int) int {
	// your code goes here
	currentTimestamp := data[0][0]
	currentNumberOfPeople := 0

	biggestNumberOfPeople := 0
	busiestTimestamp := currentTimestamp
	for i := 0; i < len(data); i++ {

		if currentTimestamp != data[i][0] {
			if biggestNumberOfPeople < currentNumberOfPeople {
				busiestTimestamp = currentTimestamp
				biggestNumberOfPeople = currentNumberOfPeople
			}

			currentNumberOfPeople = 0
			currentTimestamp = data[i][0]

		}

		if data[i][2] == 1 {
			currentNumberOfPeople += data[i][1]
		} else {
			currentNumberOfPeople -= data[i][1]
		}
	}

	return busiestTimestamp
}

func main() {
	demoData := [][]int{
		{1487799425, 14, 1},
		{1487799425, 4, 0},
		{1487799425, 2, 0},
		{1487800378, 10, 1},
		{1487801478, 18, 0},
		{1487801478, 18, 1},
		{1487901013, 1, 0},
		{1487901211, 7, 1},
		{1487901211, 7, 0},
	}
	expected := 1487800378
	if res := FindBusiestPeriod(demoData); res != expected {
		fmt.Println("fail", res)
	} else {
		fmt.Println("success")
	}
}

//timestamp, #visitors, enter/exit

/*
 [1487799425, 14, 1], +14
 [1487799425, 4,  0], -4
 [1487799425, 2,  0],
 [1487800378, 10, 1],
 [1487801478, 18, 0],
 [1487801478, 18, 1],
 [1487901013, 1,  0],
 [1487901211, 7,  1],
 [1487901211, 7,  0] ]

{
"1487799425": +14 -4 -2 = 8
.
.
.
"1487901211"
}


*/
