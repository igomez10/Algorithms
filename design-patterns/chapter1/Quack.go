package main

type Quack struct{}

func (e Quack) Quack() string {
	return "quack"
}
