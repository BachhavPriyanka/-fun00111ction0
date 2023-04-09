package main

import "fmt"

// Define a function that takes a callback as an argument
func doWork(callback func()) {
    // Perform some work
  fmt.Println("Doing work...")
  calling this  https://reqres.in/api/users?page=2
  data in datavariable
  fmt.Println("I am calling IMDB database rest api.v  )

    // Call the callback
    callback(data)
}

func main() {
    // Define a function that will be passed as a callback
    done := func() {
        fmt.Println("Work is done!")
    }

    // Call the doWork function and pass the done function as a callback
    doWork(done) // Output: "Doing work...", "Work is done!"
}




=========================================================
  
package main

import "fmt"

// Define a function that takes a callback as an argument
func normalAPICall() {
    // Perform some work
  fmt.Println("Doing work...")
  calling this api  https://reqres.in/api/users?page=2   returns json data           /// it is taking 1 mins
  storing the data in datavariable
  
// On execution datavariable = null.

  // Main logic starts here for our program
	
 fmt.Println(datavariable[1].firstname)   // Index out of bound
 fmt.Println(datavariable[1].lastname)    // Index out of bound

}


================================================================
  
  

package main

import "fmt"

// Define a function that takes a closure as an argument
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
