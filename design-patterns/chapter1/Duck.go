package main

type Duck struct {
	flyBehavior   FlyBehavior
	quackBehavior QuackBehavior
}

func (d Duck) PerformQuack() string {
	return d.quackBehavior.Quack()
}

func (d Duck) PerformFly() string {
	return d.flyBehavior.Fly()
}

func (d Duck) Swim() string {
	return "duck is swimming"
}

func (d Duck) Display() string {
	return "This is an image of a duck 🦆"
}
