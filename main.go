package main

import (
	"flag"
	"fmt"
	"github.com/krisfollegg/funtemps/conv"
)

func main() {
	// Define the input flags
	f := flag.Float64("F", 0.0, "Temperature in degrees Fahrenheit")
	c := flag.Float64("C", 0.0, "Temperature in degrees Celsius")
	k := flag.Float64("K", 0.0, "Temperature in degrees Kelvin")
	output := flag.String("out", "C", "Output unit: C (Celsius), F (Fahrenheit), K (Kelvin)")

	// Parse the command-line arguments
	flag.Parse()

	// Perform the desired conversion based on the input parameters
	switch {
	case *f != 0.0:
		if *output == "C" {
			celsius := conv.FahrenheitToCelsius(*f)
			fmt.Printf("%g°F is %g°C\n", *f, celsius)

		} else if *output == "K" {
			kelvin := conv.FahrenheitToKelvin(*f)
			fmt.Printf("%g°F is %gK\n", *f, kelvin)

		} else {
			fmt.Printf("Invalid output unit: %s\n", *output)
		}

	case *c != 0.0:
		if *output == "F" {
			fahrenheit := conv.CelsiusToFahrenheit(*c)
			fmt.Printf("%g°C is %g°F\n", *c, fahrenheit)

		} else if *output == "K" {
			kelvin := conv.CelciusToKelvin(*c)
			fmt.Printf("%g°C is %gK\n", *c, kelvin)

		} else {
			fmt.Printf("Invalid output unit: %s\n", *output)
		}

	case *k != 0.0:
		if *output == "C" {
			celsius := conv.KelvinToCelcius(*k)
			fmt.Printf("%gK is %g°C\n", *k, celsius)

		} else if *output == "F" {
			fahrenheit := conv.KelvinToFahrenheit(*k)
			fmt.Printf("%gK is %g°F\n", *k, fahrenheit)

		} else {
			fmt.Printf("Invalid output unit: %s\n", *output)
		}
		
	default:
		fmt.Println("Please provide a valid input temperature")
	}
}


