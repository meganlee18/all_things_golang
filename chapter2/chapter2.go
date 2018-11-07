package main

import "fmt"

/* Modify the program we wrote so that instead of
printing Hello World it prints Hello, my name
is followed by your name.*/

func printName(name string) string {
	return "Hello my name is " + name
}

func main() {
	fmt.Println(printName("Azkabar Germany"))
}
