package main

import (
	"math/rand"
	"time"
	"fmt"
)

type Color int
const (
	GREEN Color = 0
	RED Color = 1
	BLACK Color = 2
)

var colors = [37]Color{GREEN, RED, BLACK, RED, BLACK, RED, BLACK, RED, BLACK, RED,
			                  BLACK, BLACK, RED, BLACK, RED, BLACK, RED, BLACK, RED,
			                  RED, BLACK, RED, BLACK, RED, BLACK, RED, BLACK, RED,
			                  BLACK, BLACK, RED, BLACK, RED, BLACK, RED, BLACK, RED}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	number := rand.Intn(37)

	// Standard value-switch sentence
	switch colors[number] {
	case GREEN:
		fmt.Println("No color")
	// TODO: Help me finish this one!
	}

	// Arbitrary switch sentence (more like a shorthand for several if-else/if-else/... conditions)
	switch {
	case number == 0:
		fmt.Println("No column")
	case number % 3 == 0:
		fmt.Println("First column")
	// TODO help me finish this one!
	}

	// TODO a switch sentence to compute whether the number is odd, even, or zero
	switch {
	}

	// TODO a switch sentence to compute major/minor (1-18, 19-36, 0)
	switch {
	}

	// TODO a switch sentence to compute grouping (1-12, 13-24, 35-46, 0)
	switch {
	}
}
