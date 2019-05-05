package main

import "testing"

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

func PrintLRU(t *testing.T, obj LRUCache) {

	t.Log("Keys in cache")
	for k, v := range obj.cache {
		t.Logf("\t %d - %d\n", k, v)
	}

	t.Log("Entries in MRU")
	for iterativeHead := obj.MRU; iterativeHead != nil; iterativeHead = iterativeHead.next {
		t.Logf("\t %d \n", iterativeHead.value)
	}
}
