package main

import "fmt"

//type multiply func(x int, y int) int

func area(a, b int) int {
	return a * b
}

func perimeter(a, b int) int {
	return 2 * (a + b)
}

func print_test(myf func(int, int) int, a, b int) {
	fmt.Println(myf(a, b))
}

func diffLenghBreadth(a, b int) int {
	return (a - b)
}

func main() {

	print_test(area, 5, 4)
	print_test(perimeter, 5, 4)
	print_test(diffLenghBreadth, 5, 4)

	//************************************
	// var p multiply = func(x int, y int) int {
	// 	return x * y
	// }
	// product := p(5, 10)
	// fmt.Println("Product", product)

	// ****************************************
	//
	//	func() {
	//		fmt.Println("hello world first class function")
	//	}()
	//
	// *************************
	//
	//	newFunc := func() {
	//		fmt.Println("hello world first class function")
	//	}
	//
	// newFunc()
	// fmt.Printf("%T", newFunc)
}
