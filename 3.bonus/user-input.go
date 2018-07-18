package main

import (
	"bufio"
	"os"
	"fmt"
	"strconv"
	"strings"
)

func read(msg string) string {
	fmt.Print(msg)
	var ret, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	return ret
}

func readUint(msg, errMsg string) uint64 {
	fmt.Print(msg)
	for {
		var str, _ = bufio.NewReader(os.Stdin).ReadString('\n')
		var value, err = strconv.ParseUint(strings.Trim(str, " \n"), 10, 64)
		if err != nil {
			print("Error! %v", err.Error())
			fmt.Print(errMsg)
		} else {
			return value
		}
	}
}

func main() {
	myName := read("What's your name? ")
	fmt.Printf("Welcome, %v", myName)

	myAge := readUint("What's your age? ", "No, seriously: ")
	fmt.Printf("Got it: you are %v y.o", myAge)
}
