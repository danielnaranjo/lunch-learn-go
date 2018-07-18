package main

import (
	"math/rand"
	"time"
	"fmt"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	die6 := rand.Intn(6) + 1
	var die4 int = rand.Intn(4) + 1
	var total = die6 + die4
	fmt.Printf("Results: %v + %v = %v", die4, die6, total)
}
