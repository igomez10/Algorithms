package main

import "fmt"

// Subject interface models the required methods to
// act as a subject object
type Subject interface {
	RegisterObserver()
	RemoveObserver()
	NotifyObservers()
}

// Observer interface models the required methods to
// act as an observer object
type Observer interface {
	Update()
}

func main() {
	fmt.Println("vim-go")
}
