package main

import (
	"flag"
	"fmt"
	"github.com/krisfollegg/funtemps/conv"
)

// Definisjon for flag-variablene i hoved-"scope"
var fahr float64
var out string
var cels float64
var kelv float64
var temp string

// Bruker init for å sikre at flagvariablene er initialisert
func init() {

	// Definerer og initialiserer flagg-variablene
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celcius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")

	flag.StringVar(&out, "out", "C", "beregne temperatur i C - celsius, F - farhenheit, K- Kelvin")
	flag.StringVar(&temp, "t", "C", "tempraturskale i C - celsius, F - farhenheit, K- Kelvin")
	

}

func main() {

	flag.Parse()

	if out == "C" && isFlagPassed("F") {
		fmt.Print(formatNumber(fahr), "F is ", formatNumber(conv.FahrenheitToCelcius(fahr)), "C\n")
	}

	if out == "C" && isFlagPassed("K") {
		fmt.Print(formatNumber(kelv), "K is ", formatNumber(conv.KelvinToCelcius(kelv)), "C\n")
	}

	if out == "F" && isFlagPassed("C") {
		fmt.Print(formatNumber(cels), "C is ", formatNumber(conv.CelciusToFahrenheits(cels)), "F\n")
	}

	if out == "F" && isFlagPassed("K") {
		fmt.Print(formatNumber(kelv), "K is ", formatNumber(conv.KelvinToFahrenheit(kelv)), "F\n")
	}

	if out == "K" && isFlagPassed("C") {
		fmt.Print(formatNumber(cels), "C is ", formatNumber(conv.CelciusToKelvins(cels)), "K\n")
	}

	if out == "K" && isFlagPassed("F") {
		fmt.Print(formatNumber(fahr), "F is ", formatNumber(conv.FahrenheitToKelvin(fahr)), "K\n")
	}

}

//Sjekker om flagget er spesifisert på kommandolinje
func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}

//Formaterer tall til formatet som er ønsket
func formatNumber(convValue float64) string {
	convValue := fmt.Sprintf("%2f, convValue")

	sign := ""
	if convValueString[0:1] == "-" {
		sign = "-"
		convValueString = convValueString[1:]
	}

	numSlice := strings.Split(convValueString, ".")
	mainNum := numSlice[0]

	outputString := ""
	for i := len(mainNum); i > 0; i = i - 3 {
		if i - 3 < 0 {
			outputString = mainNum[0:i] + " " + outputString
		} else {
			outputString = mainNum[i-3:i] + " " + outputString
		}
	}

	outputString = outputString[0 : len(outputString)-1]
	if numSlice[1] != "00" {
		outputString = outputString + "." + numSlice[1]
	}

	outputString = sign + outputString

	return outputString
}
