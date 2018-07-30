package main

import "fmt"

type Foo struct {
	A string
	B string
}

func fancyPrint(element interface{}) {
	switch element.(type) {
	case int, int8, int16, int32, int64:
		fmt.Printf("Signed number: %v\n", element)
	case uint, uint8, uint16, uint32, uint64:
		fmt.Printf("Unsigned number: %v\n", element)
	case string:
		fmt.Printf("String: %v\n", element.(string))
	default:
		fmt.Printf("Other: %v\n", element)
	}
}


func main() {
	fancyPrint(uint8(1))
	fancyPrint(-1)
	fancyPrint("Foo")
	fancyPrint(false)
	fancyPrint(Foo{"bar", "baz"})
}
