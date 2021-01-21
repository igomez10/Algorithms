package main

import (
	"fmt"
	"testing"
)

func TestFlatten(t *testing.T) {
	demo := map[string]interface{}{
		"Key1": "1",
		"Key2": map[string]interface{}{
			"a": "2",
			"b": "3",
			"c": map[string]interface{}{
				"d": "3",
				"e": map[string]interface{}{
					"": "1",
				},
			},
		},
	}

	res := FlattenDictionary(demo)
	for k, v := range res {
		fmt.Println(k, v)
	}
}
