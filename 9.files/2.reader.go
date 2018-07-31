package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.OpenFile("sample.dat", os.O_RDONLY, 0400)
	if err != nil {
		fmt.Printf("Could not open the file sample.dat in read mode")
	} else {
		data := make([]byte, 255)
		n, _ := file.Read(data)
		file.Sync()
		file.Close()
		fmt.Printf("Data read! %v\n", string(data[:n]))
	}
}
