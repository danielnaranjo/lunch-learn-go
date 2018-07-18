package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math/rand"
	"time"
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

func throwDice() int {
	return rand.Intn(6) + 1
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Printf("Let's play 101/Pig. You will throw dice and see how much you can accumulate without getting 1.\n")
	accumulated := 0
	for {
		// `break` is used to exit any for loop
		// ***EXERCISE***: PLEASE help me coding this
		break
	}
	fmt.Printf("You scored: %v.\n", accumulated)
}
