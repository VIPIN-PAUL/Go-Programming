/*
The Script is for Adding two numbers given by the User.
*/

// Adding main package
package main

// Importing "fmt" package
import "fmt"

// Declaring the main function of Script
func main() {
	// Declaring the variables
	var x int
	var y int
	var z int

	// Getting values from user
	fmt.Println("Enter first value : ")
	fmt.Scan(&x)
	fmt.Println("Enter second value : ")
	fmt.Scan(&y)

	// Adding two values
	z = x + y

	// Printing the Output
	fmt.Println("Addition of ", x, " and ", y, " is ", z)
}
