package main

import "fmt"

func modifySlice(s []int) {
	s[0] = 10
	fmt.Println(s)
}

func main() {
	numbers := []int{1, 2, 3}
	modifySlice(numbers)
	fmt.Println(numbers) // Output: [10, 2, 3]
}

// package main

// import "fmt"

// func modifyValue(x *int) {
//     *x = 10
// }

// func main() {
//     num := 5
//     modifyValue(&num)
//     fmt.Println(num) // Output: 10
// }
