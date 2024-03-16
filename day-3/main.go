package main

import (
	"fmt"
)

func main() {

	var a uint8 = 255
	fmt.Println(a)
	fmt.Printf("a:%08b", a)
	fmt.Println()

	var f float64 = 123.45
	fmt.Printf("a:%064b", f)
	fmt.Println()

	var b bool = true
	fmt.Println(b)

	var byt byte = 'B'
	fmt.Printf("%d", byt)
	fmt.Println()

	var ru rune = 'A'
	fmt.Printf("%d", ru)
	fmt.Println()

	var str string = "hello"
	for _, v := range str {
		fmt.Printf("%d\n", v)
	}

}
