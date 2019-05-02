package main

import (
	"sort"
)

func main() {}

func groupDished(arr [][]string) [][]string {
	mapIngredients := make(map[string]int)
	dishes := make([]string, len(arr))

	// initialize dishes
	for i := range arr {
		dishes[i] = arr[i][0]
	}

	//initialize ingredients
	for i := range arr {
		for j := 1; j < len(arr[i]); j++ {
			if _, exists := mapIngredients[arr[i][j]]; !exists {
				mapIngredients[arr[i][j]]++
			}
		}
	}

	ingredients := []string{}

	for k := range mapIngredients {
		if mapIngredients[k] > 2 {
			ingredients = append(ingredients, k)
		}
	}

	directory := make(map[string][]string)

	//initialize all the ingredients entries in directory
	for i := range ingredients {
		directory[ingredients[i]] = []string{}
	}

	for dishIndex := range arr {
		currentDish := arr[dishIndex][0]
		for ingredientIndex := 1; ingredientIndex < len(arr[dishIndex]); ingredientIndex++ {
			directory[arr[dishIndex][ingredientIndex]] = append(directory[arr[dishIndex][ingredientIndex]], currentDish)
		}
	}

	sort.Strings(ingredients)

	// sort dishes entries in all ingredients
	for currentIngredient := range directory {
		sort.Strings(directory[currentIngredient])
	}

	// build final response array of arrays of strings
	finalResponse := [][]string{}

	for currentIngredientIndex := range ingredients {
		finalResponse = append(finalResponse, []string{})
		finalResponse[currentIngredientIndex] = append(finalResponse[currentIngredientIndex], ingredients[currentIngredientIndex])
		finalResponse[currentIngredientIndex] = append(finalResponse[currentIngredientIndex], directory[ingredients[currentIngredientIndex]]...)
	}

	return finalResponse
}
