package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Hello, I am a function
func readFloat(msg, errMsg string) float64 {
	fmt.Print(msg)
	for {
		var str, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		var value, err = strconv.ParseFloat(strings.Trim(str, " \n"), 64)
		if err != nil {
			print("Error! %v", err.Error())
			fmt.Print(errMsg)
		} else {
			return value
		}
	}
}

// And so do I, but still need to be... coded
func quadratic(a, b, c float64) {}

func main() {
	// TODO this source has an error because the function's body is missing
	a := readFloat("Enter A (any non-zero will do the job): ", "Please ensure a number: ")
	b := readFloat("Enter B: ", "Please ensure a number: ")
	c := readFloat("Enter C: ", "Please ensure a number: ")
	for _, v := range quadratic(a, b, c) {
		fmt.Printf("Result: %v\n", v)
	}
}
