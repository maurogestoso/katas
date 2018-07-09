package katas

// CelsiusToFahrenheit converts the passed temperature int in Celsius to Fahrenheit
func CelsiusToFahrenheit(celsius int) int {
	fahrenheit := celsius*9/5 + 32
	return fahrenheit
}
