package main

import "fmt"

// Complete the activityNotifications function below.

func activityNotifications(expenditure []int32, d int32) int32 {
	notificationCounter := int32(0)
	hasEasyMedian := d%2 == 1

	if hasEasyMedian {
		medianIndex := d / 2
		for i := int32(0); i+d < int32(len(expenditure)); i++ {
			if expenditure[i+d] >= 2*expenditure[medianIndex+i] {
				notificationCounter++
			}
		}
	} else {
		for i := int32(0); i+d < int32(len(expenditure)); i++ {
			if float64(expenditure[i+d]) >= 2*float64(expenditure[(d/2)+i]+expenditure[(d/2)+1+i]) {
				notificationCounter++
			}
		}
	}
	return notificationCounter
}

func main() {
	arr := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}
	d := int32(5)

	fmt.Println(activityNotifications(arr, d))
}
