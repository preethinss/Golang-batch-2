package main

import (
	"fmt"
	"log"
	"os"
)

/*type Rectangle struct {
	Width  float32
	Height float32
}

func (r *Rectangle) Area() float32 {
	r.Height = 3
	return r.Height * r.Width
}
*/
/*type myint int

func (i myint) double(a myint) myint {
	i = a
	return i * 2
}*/

func main() {
	/*r1 := Rectangle{6.0, 5.0}
	result := r1.Area()
	fmt.Println(r1.Height)
	fmt.Println(result)*/

	/*i1 := myint(10)
	result := i1.double(12)
	fmt.Println(result)*/

	/*result, err := add(10, 20)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}*/

	file, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Println("the file is created")
	content := []byte("hello")

	n, err1 := file.Write(content)

	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Println(n)

}

/*func add(a, b int) (int, error) {
	if a == 0 {
		return 0, errors.New("a should not be zero for performing the addition")
	}
	return a + b, nil
}*/
