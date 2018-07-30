package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"math/rand"
	"time"
)

// we read a string from console
func read(msg string) string {
	fmt.Print(msg)
	var ret, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	return ret
}

// we prompt for a valid RPS option
func readRPS() string {
	for {
		option := strings.ToLower(read("Choose rock, paper, scissors, or quit: "))
		switch option {
		case "rock", "paper", "scissors", "quit":
			return option
		default:
			fmt.Println("Invalid option")
		}
	}
}

// we randomly pick an option from computer
func randomRPS() string {
	options := [3]string{"rock", "paper", "scissors"}
	index := rand.Intn(3)
	return options[index]
}

func main() {
	score := 0
	rand.Seed(time.Now().UTC().UnixNano())
	game := make(map[string]map[string]int)
	// TODO help me end this part
	for {
		result := readRPS()
		if result == "quit" {
			fmt.Printf("Your final score: %v\n", score)
			return
		} else {
			opponentResult := randomRPS()
			// TODO how would I access the map, the game logic, and
			// TODO   print the result?
		}
	}
}