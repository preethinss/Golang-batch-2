package main

import (
	"fmt"
)

func main() {
	/*fmt.Println("hello world")

	var a float64 = 3.1456
	var c = int(a)

	fmt.Println(c)

	var s string = "42"
	fmt.Printf("%T", s)
	val, _ := strconv.Atoi(s)

	fmt.Printf("%T", val)*/

	result := add(3, 4)
	fmt.Println(result)

}

func add(a, b int)int {
	return a + b
}
