// Package weather provides cosas.
package weather

// CurrentCondition represents cosas.
var CurrentCondition string

// CurrentLocation represents cosas.
var CurrentLocation string

// Forecast return a weather prediction.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
