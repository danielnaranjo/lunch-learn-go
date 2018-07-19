package main

import "fmt"

type HttpErrorCode uint16

var NOT_FOUND = HttpErrorCode(404)

func main() {
	var myFloat float32 = 1.0
	var myFloat64 = 2.0
	myBool := 13 > 15
	myRune := 'c'
	myString := "Hello, World"
	var myInt = 1
	var myUint uint = 2
	var myComplex = 1 + 2i
	fmt.Printf("Values: %v %v %v %v %c %s %v %v %v %v", myFloat, myFloat64, myBool, myRune, myRune, myString, myInt, myUint, myComplex, NOT_FOUND)
}
