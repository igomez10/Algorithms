package main

import (
	"fmt"
	"strconv"
)

func FlattenDictionary(dict map[string]interface{}) map[string]string {
	finalMap := map[string]string{}
	recursiveFlatten(dict, "", &finalMap)
	return finalMap
}

func recursiveFlatten(dict map[string]interface{}, suffix string, finalMap *map[string]string) {

	for key, value := range dict {
		if suffix != "" && suffix[len(suffix)-1] != '.' && key != "" {
			suffix = suffix + "."
		}
		switch value.(type) {

		case int:
			(*finalMap)[suffix+key] = strconv.Itoa(value.(int))

		case string:
			(*finalMap)[suffix+key] = value.(string)

		case map[string]interface{}:
			childDict := dict[key].(map[string]interface{})
			recursiveFlatten(childDict, suffix+key, finalMap)

		default:
			fmt.Printf("ERROR type %T", value)
		}
	}
}
