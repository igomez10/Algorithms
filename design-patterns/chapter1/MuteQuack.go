package main

type MuteQuack struct{}

func (e MuteQuack) Quack() string {
	return "-"
}
