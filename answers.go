package main

import (
	"fmt"
)

func findExpression() bool {
	return (true && false) || (false && true) || !(false && false)
}

func compute() int {
	result := 321325 * 424521
	return result
}

func findLength(str string) int {
	return len(str)
}

func trueOrFalse(x, y string) bool {
	return x == y
}

func multipleByTwo() float64 {
	fmt.Print("Enter a number: ")
	var input float64
	//Scanf fills input with the number we enter
	fmt.Scanf("%f", &input)
	output := input * 2
	return output
}

func convertToCelsius(fahrenheit float64) float64 {
	celsius := (fahrenheit - 32) * 5 / 9
	return celsius
}

func convertToMeters(feet float64) float64 {
	meter := feet * 0.3048
	return meter
}

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

func listNumbersDivisibleByThree() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

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

func sumOfNumbers(slice []int) int {
	number := 0

	for _, value := range slice {
		number += value
	}
	return number
}

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
func makeEvenGenerator() func() uint {
	i := uint(0)

	return func() (result uint) {
		result = i
		i += 3
		return
	}
}

//Writing recursive functions
func fib(x uint) uint {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return fib(x-1) + fib(x-2)
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

	fmt.Println(fib(5))
}
