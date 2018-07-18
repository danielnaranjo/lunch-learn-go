package main

import "fmt"

func fibonacciGenerator() func() uint64 {
	var a uint64 = 0
	var b uint64 = 1
	return func() uint64 {
		sum := a + b
		a = b
		b = sum
		return a
	}
}

func main() {
	generator := fibonacciGenerator()
	fmt.Printf("Term: %v\n", generator())
	fmt.Printf("Term: %v\n", generator())
	fmt.Printf("Term: %v\n", generator())
	fmt.Printf("Term: %v\n", generator())
	fmt.Printf("Term: %v\n", generator())
}
