// Package weather provides a tool for forecasting weather.
package weather

var (
    // CurrentCondition represent a current condition.
	CurrentCondition string
	// CurrentLocation represent a current location.
    CurrentLocation  string
)

// Forecast returns a string which is forecast current weather condition looking at city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
