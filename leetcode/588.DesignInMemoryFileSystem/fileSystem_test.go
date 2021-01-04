package main

import (
	"fmt"
	"testing"
)

func TestFileSystem(t *testing.T) {
	obj := Constructor()

	// param_1 := obj.Ls("/")
	// fmt.Println(param_1)
	// obj.Mkdir("/a/b/c")

	// obj.AddContentToFile("/a/b/c/d.text", "something")
	// param_4 := obj.ReadContentFromFile("/a/b/c/d.text")
	// if param_4 != "something" {
	// 	t.Errorf("expected 'something' got %q", param_4)
	// }
	// fmt.Println(param_4)
	// fmt.Println(obj.Ls("/"))

	obj.Mkdir("/goowmfn")
	obj.Mkdir("/z")
	obj.AddContentToFile("/goowmfn/c", "shetopcy")
	fmt.Println("/goowmfn/c", obj.Ls("/goowmfn/c"))
	fmt.Println("/goowmfn", obj.Ls("/goowmfn"))
}
