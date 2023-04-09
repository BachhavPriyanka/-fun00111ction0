package main

import "fmt"

func main() {
	arr := []int{12, 0, 42, 0, 4}
	var newArray []int
	for _, val := range arr {
		if val != 0 {
			newArray = append(newArray, val)
		}
	}
	for _, val := range arr {
		if val == 0 {
			newArray = append(newArray, val)
		}
	}
	fmt.Println(newArray)

}
