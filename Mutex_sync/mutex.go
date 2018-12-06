/*
To do multiple process at a time without occuring DEAD-LOCK.
Here the code depends upon the transaction of a Bank Account.
*/
package main

// Importing the needed packages
import (
	"fmt"
	"sync"
)

// Global declaration of variables
var (
	mutex   sync.Mutex
	balance int
)

// Function with the main balance in the account
func init() {
	balance = 2000
}

// Func. is for depositing a certain amount to account
func deposit(value int, done chan bool) {
	mutex.Lock()
	fmt.Printf("Depositing %d to account with balance: %d\n", value, balance)
	balance += value
	mutex.Unlock()
	done <- true
}

// Func. is to Withdraw a certain amount from account
func withdraw(value int, done chan bool) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from account with balance: %d\n", value, balance)
	balance -= value
	mutex.Unlock()
	done <- true
}

func main() {
	fmt.Println("Go Mutex Example")

	// Channel is being created for secure execution.
	done := make(chan bool)
	go withdraw(700, done)
	go deposit(500, done)

	// Using channel for sending values
	<-done
	<-done

	// After all process, it displays the current Balance
	fmt.Printf("New Balance %d\n", balance)
}
