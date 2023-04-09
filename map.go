package main

import "fmt"

func imap(data []int, myFuncVariable func(int) int) []int {

	mapped := make([]int, len(data))

	for i, value := range data {
		mapped[i] = myFuncVariable(value)
	}

	return mapped
}

func main() {
	numbers := []int{4, 6, 8, 1}

	mapCalling := imap(numbers, func(input int) int {
		return 2 * input
	})
	fmt.Println(mapCalling)

}
