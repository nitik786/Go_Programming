package main

import "fmt"

func main() {
	// Variable declaration and initialization
	var age int = 25
	var name string = "John"

	// Variable declaration with type inference
	country := "USA"
	height := 175.5

	// Multiple variable declaration and initialization
	var x, y int = 10, 20

	// Constant declaration
	const pi = 3.14

	// Printing the values of variables
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Country:", country)
	fmt.Println("Height:", height)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("Pi:", pi)
	fmt.Println("Value of x is :", x)
}
