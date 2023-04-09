package main

import "fmt"

func max(values ...int) int {
	maxValue := -1
	for _, val := range values {
		if val > maxValue {
			maxValue = val
		}
	}
	return maxValue
}
func main() {

	m := max(4, 3, 2, 8, 7, 1)
	fmt.Println(m)
}
