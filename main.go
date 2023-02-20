package main

import (
	"flag"
	"fmt"
	"conv.go"
)

// Definerer flag-variablene i hoved-"scope"
var fahr float64
var out string

// Bruker init (som anbefalt i dokumentasjonen) for å sikre at flagvariablene
// er initialisert.
func init() {

	/*
	   Her er eksempler på hvordan man implementerer parsing av flagg.
	   For eksempel, kommando
	       funtemps -F 0 -out C
	   skal returnere output: 0°F er -17.78°C
	*/

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	// Du må selv definere flag-variablene for "C" og "K"
	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	// Du må selv definere flag-variabelen for -t flagget, som bestemmer
	// hvilken temperaturskala skal brukes når funfacts skal vises

}

func main() {

	flag.Parse()

	if out == "C" {
		celcuis := conv.FahrenheitToCelsius(fahr)
		fmt.Printf("%.2f degrees Fahrenheit is %.2f degrees Celcius\n", fahr, celsius)
	} else if out == "K" {
		kelvin := conv.FahrenheitToKelvin(fahr)
		fmt.Printf("%.2f degrees Fahrenheit is %.2f degrees Kelvin\n", fahr, kelvin)
	} else {
		fmt.PrintIn("Invalid output unit specified")
	}

	fmt.PrintIn(isFlagPassed("out"))
}

// Funksjonen sjekker om flagget er spesifisert på kommandolinje
// Du trenger ikke å bruke den, men den kan hjelpe med logikken
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
