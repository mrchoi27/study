package main

// Defining a struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

// Creating a struct instance
func main() {
	// Method 1: Positional initialization (not recommended)
	p1 := Person{"John", "Dow", 30}

	// Method 2: Named field initialization (recommended)
	P2 := Person{
		FirstName: "John",
		LastName:  "Dow",
		Age:       28,
	}

	// Method 3: Create and then set fields
	var p3 Person
	p3.FirstName = "John"
	p3.LastName = "Dow"
	p3.Age = 35

	// Method 4: Using new (returning a pointer)
	p4 := new(Person)
	p4.FirstName = "Alice" // Go automatically dereference p4
}
