package main

import "fmt"

func main() {
	myArray := [10]int{1, 2, 3, 4, 5}
	fmt.Printf("%v\n", myArray)

	myArray[4] = 6
	fmt.Printf("%v\n", myArray)

	var mySlice = myArray[1:9]
	mySlice[3] = 7
	fmt.Printf("%v\n", myArray)

	copy(myArray[5:10], []int{-1, -2, -3, -4, -5})
	fmt.Printf("%v\n", myArray)

	for index, value := range myArray {
		fmt.Printf("Index: %v, Value: %v\n", index, value)
	}

	for index := range myArray {
		fmt.Printf("Index: %v\n", index)
	}
}

