package main

import "fmt"

// function as a parameter
func returnFunction(x int, f func(int) int) int {
	return f(x)
}

// parameter and return value
func returnParameter(i int) int {
	return i
}
func main() {
	fmt.Println("Hello")
	fmt.Printf("%T\n", returnParameter)
	fmt.Printf("%T\n", returnParameter(5))

	fmt.Printf("%T\n", returnFunction)
	fmt.Printf("%d\n", returnFunction(1, returnParameter))
}
