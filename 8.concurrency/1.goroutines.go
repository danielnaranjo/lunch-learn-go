package main

import (
	"time"
	"fmt"
)

func eternalLoop(prefix string) {
	counter := 0
	for {
		// Beware! Other languages use `1` as Millisecond while here `1` means Nanosecond.
		time.Sleep(1500 * time.Millisecond)
		counter++
		fmt.Printf("Iteration %v: %v\n", prefix, counter)
	}
}

func main() {
	// Wait! Somehow my goroutine never execute
	go eternalLoop("foo")
	go eternalLoop("bar")

	// Perhaps if preventing my main() function finishing...
	// time.Sleep(15000 * time.Millisecond)
}