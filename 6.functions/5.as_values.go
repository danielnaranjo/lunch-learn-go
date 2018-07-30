package main

import "fmt"

func mapValues(elements []uint64, mapper func(uint64) uint64) []uint64{
	result := make([]uint64, len(elements))
	for idx, value := range elements {
		result[idx] = mapper(value)
	}
	return result
}

func main() {
	fmt.Printf("Result: %v", mapValues([]uint64{1, 2, 3, 4, 5}, func(x uint64) uint64 { return x * x }))
}
