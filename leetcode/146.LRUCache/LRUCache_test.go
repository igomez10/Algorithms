package main

import (
	"fmt"
	"log"
	"strconv"
	"testing"
)

func TestGetLRUCache(t *testing.T) {
	obj := Constructor(5)

	param_1 := obj.Get(1)

	if param_1 != -1 {
		t.Errorf("Unexpected valid entry in cache, expected -1 as nil element in cache but got %+v", param_1)
	}

	obj.Put(2, 2)

	param_2 := obj.Get(2)
	if param_2 != 2 {
		t.Errorf("Unexpected entry in cache, expected 2 in cache - but got %+v", param_2)
	}
}

func TestPutLRUCache(t *testing.T) {
	obj := Constructor(5)

	obj.Put(1, 1)
	obj.Put(2, 2)

	if obj.Get(1) != 1 || obj.Get(2) != 2 {
		t.Errorf("Error retrieving value from cache, expected 1,2 - got %d,%d", obj.Get(1), obj.Get(2))
	}

	obj.Put(99, 99)
	if obj.Get(99) != 99 {
		t.Errorf("Error retrieving value from cache, expected %d - got %d", 99, obj.Get(99))
		PrintLRU(t, obj)
	}

	obj.Put(4, 4)
	obj.Put(5, 5)
	obj.Put(6, 6)

	if obj.Get(6) != 6 {
		t.Errorf("Error retrieving value from cache, expected %d - got %d", 6, obj.Get(6))
		PrintLRU(t, obj)
	}

	if obj.Get(1) != -1 {
		t.Errorf("Error retrieving expelled value, expected %d but got %d", -1, obj.Get(1))
		PrintLRU(t, obj)
	}
}

func TestCache(t *testing.T) {
	type TestCase struct {
		input    []string
		capacity int
		values   [][]int
		expected []int
	}

	cases := []TestCase{
		// {
		// 	input:    []string{"get", "put", "get", "put", "put", "get", "get"},
		// 	values:   [][]int{{2}, {2, 6}, {1}, {1, 5}, {1, 2}, {1}, {2}},
		// 	expected: []int{-1, 0, -1, 0, 0, 2, 6},
		// 	capacity: 2,
		// },
		{
			input:    []string{"put", "put", "get", "put", "get", "put", "get", "get", "get"},
			values:   [][]int{{1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}},
			expected: []int{0, 0, 1, 0, -1, 0, -1, 3, 4},
			capacity: 2,
		},
	}

	for i, c := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			obj := Constructor(c.capacity)
			for in := range c.input {
				fmt.Println("running", c.input[in], c.values[in])
				res := Interpreter(&obj, c.input[in], c.values[in])
				if res != c.expected[in] {
					t.Fatal(c.input[in], c.values[in], "expected", c.expected[in], "got", res)
				}
			}
		})
	}

}

func PrintLRU(t *testing.T, obj LRUCache) {

	t.Log("Keys in cache")
	for k, v := range obj.dict {
		t.Logf("\t %d - %+v\n", k, v)
	}

	t.Log("Entries in MRU")
	for k, v := range obj.dict {
		t.Log(k, v)
	}
}

func Interpreter(obj *LRUCache, command string, values []int) int {
	switch command {
	case "get":
		return obj.Get(values[0])
	case "put":
		obj.Put(values[0], values[1])
		return 0
	default:
		log.Fatal("Unexpected command", command)
	}

	return 0
}
