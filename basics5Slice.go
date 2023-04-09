package main

import "fmt"

func sumCal() func(val int) int {
	sum := 0
	storeFunc := func(val int) int {
		sum += val
		return sum
	}
	return storeFunc
}
func main() {
	numbers := []int{3, 4, 1, 5, 1}
	sumCalculation := sumCal()

	for index, val := range numbers {
		fmt.Println(index, val, sumCalculation(val))
	}
}
