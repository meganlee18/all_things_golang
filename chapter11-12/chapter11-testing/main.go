package main

import (
	"fmt"
	"intro_to_golang/chapter11-12/chapter11-testing/math"
)

func main() {
	xs := []float64{1, 2}
	avg := math.Average(xs)
	fmt.Println(avg)

	max := math.Max(xs)
	min := math.Min(xs)

	fmt.Println(max)
	fmt.Println(min)
}
