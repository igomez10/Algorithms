package main

import "strings"
import "testing"

func TestGroupingDishes(t *testing.T) {

	initialInput := [][]string{
		[]string{"Salad", "Tomato", "Cucumber", "Salad", "Sauce"},
		[]string{"Pizza", "Tomato", "Sausage", "Sauce", "Dough"},
		[]string{"Quesadilla", "Chicken", "Cheese", "Sauce"},
		[]string{"Sandwich", "Salad", "Bread", "Tomato", "Cheese"},
	}

	result := groupDished(initialInput)

	expectedResult := [][]string{
		[]string{"Cheese", "Quesadilla", "Sandwich"},
		[]string{"Salad", "Salad", "Sandwich"},
		[]string{"Sauce", "Pizza", "Quesadilla", "Salad"},
		[]string{"Tomato", "Pizza", "Salad", "Sandwich"},
	}

	for i := range result {
		if len(result[i]) != len(expectedResult[i]) {
			t.Errorf("Error in lengths of arrays, expected %d - got %d", len(expectedResult[i]), len(result[i]))
			for x := range result {
				t.Log(result[x])
			}
		}

		for j := range result[i] {
			if strings.Compare(result[i][j], expectedResult[i][j]) != 0 {
				t.Errorf("Unexpected string, expected %s - got %s", expectedResult[i][j], result[i][j])
			}
		}
	}
}
