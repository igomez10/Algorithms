package main

import "fmt"

type Dog struct{}
type Cat struct{}

type Animal interface {
	makeSound() string
}

func (c Cat) Meow() string {
	return "Meow"
}

func (d Dog) Bark() string {
	return "Guau"
}

func (c Cat) makeSound() string {
	return c.Meow()
}

func (d Dog) makeSound() string {
	return d.Bark()
}

func main() {
	fmt.Println("vim-go")
	var animal Animal
	animal = Dog{}
	fmt.Println(animal.makeSound())
}
