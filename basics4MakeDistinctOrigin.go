package main

import (
	"fmt"
	"math"
)

func makeDistOrigin(x_o, y_o float64) func(x, y float64) float64 {
	calculate := func(x, y float64) float64 {
		return math.Sqrt(math.Pow(x-x_o, 2) + math.Pow(y-y_o, 2))
	}
	return calculate
}

func main() {

	
	dist1 := makeDistOrigin(5, 10)
	Dist1 := dist1(3, 4)
	fmt.Println(Dist1)

	dist2 := makeDistOrigin(15, 20)
	Dist2 := dist2(4, 2)
	fmt.Println(Dist2)
}
