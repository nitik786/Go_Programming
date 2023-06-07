package main

import "fmt"

// Define a struct type for a Person
type Person struct {
	name string
	age  int
}

// Define a method on the Person struct
func (p Person) introduce() {
	fmt.Printf("Hi, my name is %s and I'm %d years old.\n", p.name, p.age)
}

func main() {
	// Create an instance of Person struct
	person := Person{
		name: "John",
		age:  25,
	}
	nitik := Person{
		name: "Nitik",
		age:  25,
	}
	// Access and print the fields of the Person struct
	fmt.Println("Name:", person.name)
	fmt.Println("Age:", person.age)
	fmt.Println("--------------------------------------------------")
	fmt.Println("Name:", nitik.name)
	fmt.Println("Age:", nitik.age)
	// Call the introduce() method on the Person struct
	person.introduce()
	nitik.introduce()
}
