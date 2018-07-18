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
	if readUint("Tell me your age so I can tell whether you can vote or not: ", "No, seriously: ") >= 18 {
		fmt.Println("Yep, you can vote in next elections")
	} else {
		fmt.Println("Nope, you are too young")
	}
}
