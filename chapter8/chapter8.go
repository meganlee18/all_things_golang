package main

import (
	"fmt"
)

/* What is the value of x after running this program: */
/* Answer: 2.25? */
func square(w *float64) {
	*w = *w * *w
}

// Write a program that can swap two integers
// (x := 1; y := 2; swap(&x, &y) should give you
// x=2 and y=1).
func swap(x *int, y *int) {
	*x = 2
	*y = 1
}

func main() {
	w := 1.5
	square(&w)
	fmt.Println(w)

	//Run swap function and dereference pointers
	x := 1
	y := 2
	swap(&x, &y)
	fmt.Printf("x=%v, y=%v", x, y)

}
