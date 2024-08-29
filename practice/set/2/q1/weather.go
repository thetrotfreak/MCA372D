package main

import (
	"fmt"
	"time"
)

// Store coordinates
// perhaps we should store a string name too
type Coord struct {
	latitude  float64
	longitude float64
}

// Weather structure to hold weather measurements
type Weather struct {
	location    Coord
	tinstance   time.Time
	temperature float64
	humidity    int
	windspeed   float64
	// what about the direction?
}

func main() {
	// Creating Weather instance
	currentWeather := Weather{
		location: Coord{
			latitude:  12.934371,
			longitude: 77.605703,
		},
		temperature: 25.5,
		humidity:    60,
		windspeed:   12.3,
		tinstance:   time.Now(),
	}

	// that specific date is Go's time reference
	fmt.Println("Weather At  :", currentWeather.location)
	fmt.Printf("DateTime    : %s\n", currentWeather.tinstance.Format("2006-01-02 15:04:05"))
	fmt.Printf("Temperature : %.2f°C\n", currentWeather.temperature)
	fmt.Printf("Humidity    : %d%%\n", currentWeather.humidity)
	fmt.Printf("Wind Speed  : %.2f m/s\n", currentWeather.windspeed)
}
