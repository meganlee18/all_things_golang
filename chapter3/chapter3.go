package main

import "fmt"

/* What's the value of the expression (true && false) || (false && true) || !(false && false)? */
func findExpression() bool {
	return (true && false) || (false && true) || !(false && false)
}

/* Although overpowered for the task you can use Go as a calculator. Write a program that computes 321325 × 424521 and prints it to the terminal. */
func compute() int {
	result := 321325 * 424521
	return result
}

/* What is a string? How do you find its length? */
func findLength(str string) int {
	return len(str)
}

func main() {
	fmt.Println(findExpression())
	fmt.Println(compute())
	fmt.Println(findLength("test this piece of string"))
}
