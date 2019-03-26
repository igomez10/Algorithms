package main

// WeatherStation defines the structure to be used as the weatehr station
type WeatherStation struct {
	Observers   []Observer
	Temperature float64
	Pressure    float64
	Humidity    float64
}

// RegisterObserver adds an observer to its array of observers
func (w WeatherStation) RegisterObserver(obs Observer) {
	w.Observers = append(w.Observers, obs)
}

// RemoveObserver removes the provided observer from the list of observers
func (w WeatherStation) RemoveObserver(obs Observer) {
	for i := 0; i < len(w.Observers); i++ {
		if w.Observers[i] == obs {
			w.Observers = append(w.Observers[:i], w.Observers[i+1:]...)
		}
	}
}

// NotifyObservers traverses the list of observers notifying them on the change
func (w WeatherStation) NotifyObservers() {
	for i := 0; i < len(w.Observers); i++ {
		w.Observers[i].Update()
	}
}
