package main

type FlyNoWay struct {
}

func (e FlyNoWay) Fly() string {
	return "Cannot fly"
}
