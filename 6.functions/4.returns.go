package main

import "fmt"

func doSomething(something string) {
	fmt.Printf("This function returns no value, but prints: %v\n", something)
}

func returnSomething() string {
	fmt.Println("This function does something, and returns a value")
	return "Hello, World"
}

func returnTwoValues() (string, string) {
	fmt.Println("This function does something, and returns two values")
	return "Hello", "World"
}

func returnNamedSomething() (rv string) {
	fmt.Println("This function returns a value in a variable")
	rv = "Hello, World"
	return
}

func returnTwoNamedValues() (rv1, rv2 string) {
	fmt.Println("This function does something, and returns two values in variables")
	rv1, rv2 = "Hello", "World"
	return
}


func main() {
	doSomething("Hello")
	fmt.Printf("returned: %v", returnSomething())
	foo, bar := returnTwoValues()
	fmt.Printf("returned: %v, %v", foo, bar)
	fmt.Printf("returned: %v", returnNamedSomething())
	foo, bar = returnTwoNamedValues()
	fmt.Printf("returned: %v, %v", foo, bar)
}
