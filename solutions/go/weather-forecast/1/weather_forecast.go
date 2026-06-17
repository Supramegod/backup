// Package weather provides weather forecast information.
package weather

// CurrentCondition stores the current weather condition.
var CurrentCondition string

// CurrentLocation stores the current location.
var CurrentLocation string

// Forecast returns the current weather forecast for a city.
func Forecast(city, condition string) string {
	CurrentLocation = city
	CurrentCondition = condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}