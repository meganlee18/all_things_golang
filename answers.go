package main

import (
	"fmt"
)

// What's the value of the expression (true &&
// 	false) || (false && true) || !(false &&
// 	false)?
func findExpression() bool {
	return (true && false) || (false && true) || !(false && false)
}

// Although overpowered for the task you can use
// Go as a calculator. Write a program that computes
// 321325 × 424521 and prints it to the terminal.
func compute() int {
	result := 321325 * 424521
	return result
}

// What is a string? How do you find its length?
func findLength(str string) int {
	return len(str)
}

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

//Print numbers from 1 to 10
func printNumbers() {
	for number := 1; number <= 10; number++ {
		// if number%2 == 0 {
		// 	fmt.Println(number, "-> even")
		// } else {
		// 	fmt.Println(number, "-> odd")
		// }

		switch number {
		case 1:
			fmt.Println(number, "-> even")
		case 2:
			fmt.Println(number, "-> odd")
		default:
			fmt.Println(number, "Unknown Number")
		}
	}
}

// Write a program that prints out all the numbers
// evenly divisible by 3 between 1 and 100. (3, 6, 9,
// etc.)
func listNumbersDivisibleByThree() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

//Fizzbuzz challenge:
// //Write a program that prints the numbers from 1
// to 100. But for multiples of three print "Fizz" instead
// of the number and for the multiples of five
// print "Buzz". For numbers which are multiples
// of both three and five print "FizzBuzz".
func fizzBuzz() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, "-> FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println(i, "-> Fizz")
		} else if i%5 == 0 {
			fmt.Println(i, "-> Buzz")
		}
	}
}

//
func findAverage() {
	var x [5]float64
	var total float64

	x[0] = 98
	x[1] = 93
	x[2] = 77
	x[3] = 82
	x[4] = 83

	// for i := 0; i < 5; i++ {
	// 	total += x[i]
	// }

	// average := total / float64(len(x))
	// fmt.Println(average)

	for _, value := range x {
		total += value
		//fmt.Println(index, value)
	}
	fmt.Println(total / float64(len(x)))
}

func printArrayAndSlice() {
	//array: length is fixed and part of array type name
	array := [5]float64{
		98,
		23,
		56,
		11,
		//89,
	}

	//array slices are indexable and have a length
	//cannot be longer than array but can be smaller
	slice := make([]float64, 5, 6)

	fmt.Println(array)
	fmt.Println(slice)
}

// Write a program that finds the smallest number
// in this list:
func findSmallestNumber() {

	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	number := x[0]

	for _, value := range x {
		if value < number {
			number = value
		}
	}
	fmt.Println(number)
}

// sum is a function which takes a slice of numbers
// and adds them together. What would its function
// signature look like in Go?
func sumOfNumbers(slice []int) int {
	number := 0

	for _, value := range slice {
		number += value
	}
	return number
}

// Write a function which takes an integer and
// halves it and returns true if it was even or false
// if it was odd. For example half(1) should return
// (0, false) and half(2) should return (1,
// true).
func half(number int) string {
	halved := number / 2

	if number%2 == 0 {
		//even: return (1, true)
		fmt.Println(halved)
		return "(1, true)"
	}
	//odd: return (0, false)
	fmt.Println(halved)
	return "(0, false)"
}

//variadic functions - take more than int
// Write a function with one variadic parameter
// that finds the greatest number in a list of numbers.
func findLargestNumber(args ...int) int {
	number := args[0]

	for _, value := range args {
		if value > number {
			number = value
		}
	}
	return number
}

//Writing closure functions
// Using makeEvenGenerator as an example, write a
// makeOddGenerator function that generates odd
// numbers.
func makeEvenGenerator() func() uint {
	i := uint(0)

	return func() (result uint) {
		result = i
		i += 3
		return
	}
}

//Writing recursive functions
// The Fibonacci sequence is defined as: fib(0) =
// 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2).
// Write a recursive function which can find
// fib(n).
func fib(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
}

// Write a program that can swap two integers
// (x := 1; y := 2; swap(&x, &y) should give you
// x=2 and y=1).
func swap(x *int, y *int) {
	*x = 2
	*y = 1
}

func main() {
	//fmt.Println(findExpression())
	//fmt.Println(compute())
	//fmt.Println(findLength("how long is this?"))
	//fmt.Println(trueOrFalse("hello", "world"))
	//fmt.Println(multipleByTwo()
	//fmt.Println(convertToCelsius(324.23))
	//fmt.Println(convertToMeters(20))
	//printNumbers()
	//listNumbersDivisibleByThree()
	//fizzBuzz()
	//findAverage()
	//printArrayAndSlice()
	//findSmallestNumber()
	//fmt.Println(sumOfNumbers([]int{111, 3325, 444}))
	//fmt.Println(half(1))
	//fmt.Println(findLargestNumber(341, 7845, 456))

	// nextOdd := makeEvenGenerator()
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())
	// fmt.Println(nextOdd())

	//fmt.Println(fib(5))

	//Run swap function and dereference pointers
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Printf("x=%v, y=%v", x, y)

}
