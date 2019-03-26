package main

import "fmt"

// Subject interface models the required methods to
// act as a subject object
type Subject interface {
	RegisterObserver(Observer)
	RemoveObserver(Observer)
	NotifyObservers()
}

// Observer interface models the required methods to
// act as an observer object
type Observer interface {
	Update(temp float64, pressure float64, humid float64)
	Observe(Subject)
}

func main() {

	station := WeatherStation{Observers: []Observer{}}

	var currentD CurrentDisplay
	var forecastD ForecastDisplay
	var statsD StatisticsDisplay

	currentD.Observe(&station)
	statsD.Observe(&station)
	forecastD.Observe(&station)

	station.SetWeather(10.0, 1.0, 1.0)

	fmt.Println("vim-go")
}
