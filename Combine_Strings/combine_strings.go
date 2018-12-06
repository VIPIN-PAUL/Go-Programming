/*
The Script is for Concatenating Two Strings
*/
package main

// Importing "fmt" package
import "fmt"

func main() {
	//  Declaring 2 varibles with "string" type
	var str1, str2 string

	// Getting the input from user and storing it into varible
	fmt.Print("Enter first string : ")
	fmt.Scan(&str1)

	fmt.Print("Enter second string : ")
	fmt.Scan(&str2)

	// Concatenating both the strings in EASY way and displaying it.
	fmt.Print(str1 + " " + str2 + "\n")
}
