package main

import "fmt"

func sub() func(x, y int) int {
	funcVariable := func(x, y int) int {
		return x - y
	}
	return funcVariable
}

// *******************
func testing(demofunc func(a, b int) int) {
	fmt.Println(demofunc(4, 5))

}

// ***********************
func add1(x int, y int) int {
	return x + y
}
func main() {
	addition := add1(4, 5)
	fmt.Println(addition)

	test1 := func(x int, y int) int {
		return x * y
	}
	testing(test1)

	storeSub := sub()
	fmt.Println(storeSub(5, 6))
}
