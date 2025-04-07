package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Address struct {
	Street  string
	City    string
	Country string
}

type Employee struct {
	Person   // Embedded struct (no field name)
	Address  // Another embedded struct
	Position string
	Salary   float64
}

func main() {
	emp := Employee{
		Person: Person{
			FirstName: "John",
			LastName:  "Smith",
			Age:       35,
		},
		Address: Address{
			Street:  "123 Main St",
			City:    "Boston",
			Country: "USA",
		},
		Position: "Software Engineer",
		Salary:   90000,
	}

	// Fields of embedded structs can be accessed directly
	fmt.Println(emp.FirstName) // Output: John
	fmt.Println(emp.Street)    // Output: 123 Main st

	// Or through the embedded struct's field name
	fmt.Println(emp.Person.LastName) // Output: Smith
	fmt.Println(emp.Address.City)    // Output: Boston
}
