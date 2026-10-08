// Package weather is a package for weather.
package weather

var (
    // CurrentCondition is a string.
	CurrentCondition string
    // CurrentLocation is a string.
	CurrentLocation  string
)
// Forecast is a function.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
