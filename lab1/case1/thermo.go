package main

import "fmt"

func celsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func celsiusToKelvin(c float64) float64 {
	return c + 273.15
}

func kelvinToCelsius(k float64) float64 {
	return k - 273.15
}

func fahrenheitToKelvin(f float64) float64 {
	return celsiusToKelvin(fahrenheitToCelsius(f))
}

func kelvinToFahrenheit(k float64) float64 {
	return celsiusToFahrenheit(kelvinToCelsius(k))
}

func main() {
	var c, f, k float64

	fmt.Print("Temperature in °C? ")
	fmt.Scan(&c)
	fmt.Printf("%.2f°C is %.2f°F\n", c, celsiusToFahrenheit(c))

	fmt.Print("Temperature in °F? ")
	fmt.Scan(&f)
	fmt.Printf("%.2f°F is %.2f°C\n", f, fahrenheitToCelsius(f))

	fmt.Print("Temperature in °C? ")
	fmt.Scan(&c)
	fmt.Printf("%.2f°C is %.2fK\n", c, celsiusToKelvin(c))

	fmt.Print("Temperature in K? ")
	fmt.Scan(&k)
	fmt.Printf("%.2fK is %.2f°C\n", k, kelvinToCelsius(k))

	fmt.Print("Temperature in °F? ")
	fmt.Scan(&f)
	fmt.Printf("%.2f°F is %.2fK\n", f, fahrenheitToKelvin(f))

	fmt.Print("Temperature in K? ")
	fmt.Scan(&k)
	fmt.Printf("%.2fK is %.2f°F\n", k, kelvinToFahrenheit(k))
}
