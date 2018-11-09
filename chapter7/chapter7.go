package main

import (
	"fmt"
)

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

func main() {
	fmt.Println(sumOfNumbers([]int{111, 3325, 444}))
	fmt.Println(half(1))
	fmt.Println(findLargestNumber(341, 7845, 456))

	nextOdd := makeEvenGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())

	fmt.Println(fib(5))
}
