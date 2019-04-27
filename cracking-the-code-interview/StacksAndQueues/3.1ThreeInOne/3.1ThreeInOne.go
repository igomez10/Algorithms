package main

import "fmt"

type hack struct {
	value interface{}
}

func main() {

	threeInOne := [3]hack{}
	threeInOne[0].value = []string{"hello"}
	threeInOne[1].value = 1
	threeInOne[2].value = func(par interface{}) { fmt.Printf("%+v", par) }

	for i := range threeInOne {
		fmt.Println(threeInOne[i])
	}

	threeInOne[2].value.(func(interface{}))("hey")

}
