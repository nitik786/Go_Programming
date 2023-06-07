package main

import "fmt"

func modifyValue(x int) {
	x = 10
	fmt.Println(x)
}

func main() {
	num := 5
	modifyValue(num)
	fmt.Println(num) // Output: 5
}
