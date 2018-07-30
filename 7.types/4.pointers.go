package main

import "fmt"

func alterValue(stringPtr *string) {
	*stringPtr = *stringPtr + " World"
}

func main() {
	value := "Hello"
	alterValue(&value)
	fmt.Printf(value)
}
