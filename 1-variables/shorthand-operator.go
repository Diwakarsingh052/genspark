package main

import "fmt"

func main() {
	a, greet := 1, "hello" // we are creating two variables

	a, ok := 10, true // if a variable already exist it would update the existing one,
	// and create other variables which are not already present

	// a := 20 // this would not work // we need atleast one new variable at the left side

	fmt.Println(a, ok, greet)
}
