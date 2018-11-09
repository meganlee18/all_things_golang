package main

import (
	"fmt"
)

// Determines if x is equal to y
func trueOrFalse(x, y string) bool {
	return x == y
}

//Allows user input and multiply that by 2
func multipleByTwo() float64 {
	fmt.Print("Enter a number: ")
	var input float64
	//Scanf fills input with the number we enter
	fmt.Scanf("%f", &input)
	output := input * 2
	return output
}

// Using the example program as a starting point,
// write a program that converts from Fahrenheit
// into Celsius. (C = (F - 32) * 5/9)
func convertToCelsius(fahrenheit float64) float64 {
	celsius := (fahrenheit - 32) * 5 / 9
	return celsius
}

//Write another program that converts from feet
// into meters. (1 ft = 0.3048 m)
func convertToMeters(feet float64) float64 {
	meter := feet * 0.3048
	return meter
}

func main() {
	fmt.Println(trueOrFalse("hello", "world"))
	fmt.Println(multipleByTwo())
	fmt.Println(convertToCelsius(324.23))
	fmt.Println(convertToMeters(20))
}
