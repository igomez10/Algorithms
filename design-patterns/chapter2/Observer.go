package main

import "fmt"

type Display interface {
	Render()
}

type ForecastDisplay struct {
	observable  Subject
	Temperature float64
	Pressure    float64
	Humidity    float64
}

type StatisticsDisplay struct {
	observable   Subject
	Temperatures []float64
	Pressures    []float64
	Humidities   []float64
}

type CurrentDisplay struct {
	observable  Subject
	Temperature float64
	Pressure    float64
	Humidity    float64
}

func (f *ForecastDisplay) Observe(s Subject) {
	s.RegisterObserver(f)
}
func (f *CurrentDisplay) Observe(s Subject) {
	s.RegisterObserver(f)
}
func (f *StatisticsDisplay) Observe(s Subject) {
	s.RegisterObserver(f)
}

func (f *ForecastDisplay) Update(temp float64, press float64, humid float64) {
	f.Temperature = temp + 1
	f.Pressure = press + 1
	f.Humidity = humid + 1
	f.Render()
}
func (f *CurrentDisplay) Update(temp float64, press float64, humid float64) {
	f.Temperature = temp
	f.Pressure = press
	f.Humidity = humid
	f.Render()
}
func (f *StatisticsDisplay) Update(temp float64, press float64, humid float64) {
	f.Temperatures = append(f.Temperatures, temp)
	f.Pressures = append(f.Pressures, press)
	f.Humidities = append(f.Humidities, humid)
	f.Render()
}

func (f *StatisticsDisplay) Render() {
	avgTemp := average(f.Temperatures)
	fmt.Println("this is the average temperature: ", avgTemp)
	avgPressure := average(f.Pressures)
	fmt.Println("this is the average pressure: ", avgPressure)
	avgHumidity := average(f.Humidities)
	fmt.Println("this is the average humidity: ", avgHumidity)
}

func (f *ForecastDisplay) Render() {
	fmt.Println("this is the forecast: TODO")
}

func (c *CurrentDisplay) Render() {
	fmt.Printf("this is the temperature: %f", c.Temperature)
	fmt.Printf("this is the pressure: %f", c.Pressure)
	fmt.Printf("this is the humidity: %f", c.Humidity)
}

func average(elements []float64) float64 {
	var avg float64
	for i := range elements {
		avg += elements[i]
	}
	avg = avg / float64(len(elements))
	return avg
}
