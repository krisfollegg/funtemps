package main

import (
	"fmt"
	"github.com/krisfollegg/funtemps/conv"
)

func main() {
	f := 134.0
	c := conv.FahrenheitToCelsius(f)
	fmt.Printf("%v degrees Fahrenheit is %v degrees Celsius\n", f, c)

	c = 56.7
	f = conv.CelciusToFahrenheit(c)
	fmt.Printf("%v degrees Celsius is %v degrees Fahrenheit\n", c, f)

	k := 9392.0
	f = conv.KelvinToFahrenheit(k)
	fmt.Printf("%v degrees Kelvin is %v degrees Fahrenheit\n", k, f)

	f = 134.0
	k = conv.FahrenheitToKelvin(f)
	fmt.Printf("%v degrees Fahrenheit is %v degrees Kelvin\n", f, k)

	c = 56.7
	k = conv.CelciusToKelvin(c)
	fmt.Printf("%v degrees Celsius is %v degrees Kelvin\n", c, k)

	k = 329.82
	c = conv.KelvinToCelcius(k)
	fmt.Printf("%v degrees Kelvin is %v degrees Celsius\n", k, c)
}
