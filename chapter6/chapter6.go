package main

import (
	"fmt"
)

//Find the average of slice
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

func main() {
	findAverage()
	printArrayAndSlice()
	findSmallestNumber()
}
