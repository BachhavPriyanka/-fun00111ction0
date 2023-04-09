package main

import "fmt"

// func main() {
// 	x := 10
// 	// addX := func(y int) int {
// 	// 	return x + y
// 	// }
// 	// result := addX(20)
// 	// fmt.Println(result) // Output: 30

// 	subX := subtraction()
// 	result := subX(5)
// 	fmt.Println(result)
// }

// func subtraction(y int) int {
// 	return x - y
// }

// ***********************************
// func doWork(callback func()) {
// 	// Perform some work
// 	fmt.Println("Doing work...")

// 	// Call the callback
// 	callback()
// }

// func main() {
// 	// Define a function that will be passed as a callback
// 	done := func() {
// 		fmt.Println("Work is done!")
// 	}

// 	// Call the doWork function and pass the done function as a callback
// 	doWork(done) // Output: "Doing work...", "Work is done!"
// }

func doLater(f func()) {
    // Store the closure in a variable for later use
    deferredFunc := f

    // Perform some work
    fmt.Println("Doing work...")

    // Execute the closure at a later time
    deferredFunc()
}

func main() {
    // Define a closure that will be passed to the doLater function
    done := func() {
        fmt.Println("Work is done!")
    }

    // Call the doLater function and pass the closure as an argument
    doLater(done) // Output: "Doing work...", "Work is done!"
}