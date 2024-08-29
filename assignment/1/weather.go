package main

import "fmt"

func getWeatherRecommendation(temp float64) string {
	switch {
	case temp < 10:
		return "Wear a heavy jacket."
	case temp <= 20:
		return "Wear a light jacket."
	default:
		return "Wear a t-shirt."
	}
}

func main() {
	var temperature float64

	fmt.Print("What's the temperature around you(in °C)? ")
	fmt.Scan(&temperature)

	recommendation := getWeatherRecommendation(temperature)

	fmt.Println("Hey! I'd recommend that you:", recommendation)
}
