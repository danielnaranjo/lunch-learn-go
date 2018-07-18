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

func readCount() int {
	level := readUint("Let's shoot a magic missile. What's your character's level? ", "No, seriously: ")
	var count int = int((level - 1) / 2) + 1
	if count > 5 {
		return 5
	} else {
		return count
	}
}

func throwOne() int {
	return rand.Intn(4) + 1
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	count := readCount()
	damage := 0
	for i := 0; i < count; i++ {
		partialDamage := throwOne()
		fmt.Printf("Throwing: %v...\n", partialDamage)
		damage += partialDamage
	}
	fmt.Printf("Total damage: %v\v", damage)
}
