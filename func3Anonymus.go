package main

import "fmt"

type add func(x int, y int) int

func main() {

	mufunc := func() { //Anonymus Func
		fmt.Println("Welcome")
	}
	mufunc()

	first()
	second()
	//*****************
	func() {
		fmt.Println("This is Anonyus func 2")
	}()

	//*****************
	mufunc = first
	mufunc()
	//***************
	var p add = func(x int, y int) int {
		return x + y
	}
	addition := p(7, 8)
	fmt.Println("addition : ", addition)

	//***************
	crazyishFunction := func(f func()) func() {
		return f
	}

	crazyFunction(crazyishFunction)(first)()

}

func crazyFunction(f func(func()) func()) func(func()) func() {
	return f
}

func first() {
	fmt.Println("Inside first func")
}

func second() {
	fmt.Println("Inside second func")
}
