//Package weather forecast provides tools to retrieve weather data on a per city basis.
package weather

// CurrentCondition represents a specific whether state for a city.
var CurrentCondition string
// CurrentLocation represents a specific city.
var CurrentLocation string
// Forecast takes in a city (location) and a condition (weather state) and returns a string detailing the state of a cities weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
