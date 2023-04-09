package main

import "fmt"

func div() func(a int, b int) int {
	return func(a int, b int) int {
		return a / b
	}
}

func multi() func(a int, b int) int {
	m := func(a int, b int) int {
		return a * b
	}
	return m
}
func main() {

	multiVar := multi()
	fmt.Println(multiVar(3, 6))

	divVar := div()
	divVar2 := divVar(8, 2)
	fmt.Println(divVar2)
}
