package main

import (
	"fmt"
	"time"
)

func grabDataFromInternet(channel chan<- string) {
	elements := []string{"Lorem", "Ipsum", "Dolor", "Sit", "Amet", "Consectetur", "Adipiscing", "Elit"}
	for _, element := range elements {
		fmt.Printf("Sending element: %v\n", element)
		channel <- element
	}
	close(channel)
}

func main() {
	// Change it to a buffered channel to see what happens!
	channel := make(chan string, 4)  // Add ,4 to see what happens
	go grabDataFromInternet(channel)
	for {
		// Read data while the channel is open
		data, stillOpen := <-channel
		if !stillOpen {
			fmt.Println("Channel closed")
			break;
		} else {
			fmt.Printf("Received element: %v\n", data)
			time.Sleep(1500 * time.Millisecond)
		}
	}
}