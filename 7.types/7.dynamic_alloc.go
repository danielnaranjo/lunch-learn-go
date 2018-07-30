package main

import "fmt"

type Person struct {
	FirstName string
	LastName string
}

func (person *Person) String() string {
	return fmt.Sprintf("Person(first name=%v, last name=%v)", person.FirstName, person.LastName)
}

func NewPerson(firstName, lastName string) *Person {
	// There will be no dangling pointer here!
	return &Person{firstName, lastName}
}

func main() {
	fmt.Printf("New person: %v\n", NewPerson("Nestor", "Kirchner"))
}
