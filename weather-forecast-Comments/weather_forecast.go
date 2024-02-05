// Package weather provides tools to describe weather.
package weather

// CurrentCondition represents a current weather condition.
var CurrentCondition string

// CurrentLocation represents a current weather condition.
var CurrentLocation string

// Forecast returns an string describing the current weather conditions for a given city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
