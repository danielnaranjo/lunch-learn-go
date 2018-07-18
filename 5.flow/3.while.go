package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

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
	number := readUint("Let's test a number into Collatz's: ", "No, seriously, a positive number: ")
	fmt.Printf("Ok. Number will start as %v.\n", number)
	for number > 1 {
		if number % 2 == 1 {
			number = number * 3 + 1
		} else {
			number = number / 2
		}
		fmt.Printf("Next step: %v.\n", number)
	}
	fmt.Println("Aaaand we're done.")
}
