package main

import (
	"math/rand"
	"time"
	"fmt"
)

func sampleGoroutine(channel chan<- string, prefix string) {
	counter := 0
	for {
		time.Sleep(time.Second * time.Duration(2 + rand.Intn(3)))
		counter++
		channel <- fmt.Sprintf("%v %v", prefix, counter)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	chan1 := make(chan string)
	chan2 := make(chan string)
	go sampleGoroutine(chan1, "first channel")
	go sampleGoroutine(chan2, "second channel")
	for {
		select {
		case x := <- chan1:
			fmt.Printf("arrived: %v\n", x)
		case y := <- chan2:
			fmt.Printf("arrived: %v\n", y)
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("channel not ready")
		}
	}
}
