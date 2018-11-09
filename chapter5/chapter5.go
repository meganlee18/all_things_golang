package main

import (
	"fmt"
)

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

func main() {
	//printNumbers()
	//listNumbersDivisibleByThree()
	//fizzBuzz()
}
