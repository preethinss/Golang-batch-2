package main

import "fmt"

/*type Number interface {
	int | float64
}*/

func main() {
	fmt.Println("hello")
	result := add(3, 4)

	fmt.Println(result)

}
func add(a, b int) int {
	return a + b

}
