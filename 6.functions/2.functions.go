package main

import "fmt"

// Functions can be taken as arguments
// TODO convert this function to have a generic behaviour instead of just multiplying by 2
func arrayTransform(elements []float64) []float64 {
	newElements := make([]float64, len(elements))
	for index, element := range elements {
		newElements[index] = element * 2
	}
	return newElements
}

func main() {
	// TODO also convert this call according to the function change
	fmt.Printf("%v", arrayTransform([]float64{1, 2, 3, 4, 5}))
}
