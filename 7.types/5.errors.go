package main

import (
	"errors"
	"fmt"
)

// Errors are just a matter of conventions
func divMod(a, b int64) (int64, int64, error) {
	if b == 0 {
		return 0, 0, errors.New("Cannot divide by zero")
	} else {
		return a / b, a % b, nil
	}
}

// Panics are like regular exceptions in other languages
func divModWithPanic(a, b int64) (int64, int64) {
	return a / b, a % b
}

func recoverFromPanic() {
	value := recover()
	if value != nil {
		fmt.Printf("Recovered from panic: %v", value)
	}
}

func main() {
	defer recoverFromPanic()
	a, b := divModWithPanic(1, 0)
	fmt.Printf("Results for a=%v b=%v", a, b)

	_, _, err := divMod(1, 0)
	fmt.Print("Error raised: " + err.Error())

	a2, b2, _ := divMod(12, 7)
	fmt.Printf("Results: 12/7=%v, 12%%7=%v", a2, b2)
}