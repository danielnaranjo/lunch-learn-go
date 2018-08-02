package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.OpenFile("sample.dat", os.O_CREATE | os.O_APPEND | os.O_WRONLY, 0400)
	if err != nil {
		fmt.Println("Could not open the file sample.dat in write mode")
	} else {
		file.Write([]byte("Lorem ipsum dolor sit amet consectetur adipiscing elit."))
		file.Close()
		fmt.Println("Done!")
	}
}
