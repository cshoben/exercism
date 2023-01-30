// Package weather forecasts the current weather condition in various cities in Goblinocus.
package weather

// CurrentCondition holds a cities current weather condition.
var CurrentCondition string
// CurrentLocation holds the city name where the weather is being forecasted.
var CurrentLocation string

// Forecast takes in a city name and condition and returns it in a printed statement.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
