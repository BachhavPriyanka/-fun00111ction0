package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

//**********************
// func test() func(x, y int) int {
// 	f := func(x, y int) int {
// 		return x + y
// 	}
// 	return f
// }

func main() {
	// t := test()
	// fmt.Println(t(17, 7))
	//************************
	// num := 108
	// func() {
	// 	fmt.Println(num)
	// }()
	//***********************
	testSlice := []int{10, 10, 30, 30}

	sum := adder()
	for i, val := range testSlice {
		fmt.Println(i, val, sum(val))
	}

}
