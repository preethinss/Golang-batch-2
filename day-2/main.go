package main

import (
	"fmt"

	pa1 "github.com/PreethiNS/day-2/package1"
)

var package_level_var = 30

func main() {
	/*//Declaration
	var a int

	//Intialization
	a = 20
	fmt.Println(a)

	//decalaration and intialization
	var b int = 30
	fmt.Println(b)

	//Type inference
	var c = 40
	fmt.Println(c)
	fmt.Printf("%T", c)
	fmt.Println()

	//short hand declaration
	d := 50
	fmt.Println(d)
	fmt.Printf("%T", d)*/

	/*var name32 string = "xxx"
	fmt.Println(name32)

	var Name string= "yyy"
	fmt.Println(Name)*/

	/*var a, b int = 10, 20
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)

	c, name := 3.14, "xxx"
	fmt.Println(c, name)*/

	//local variable
	/*var name string = "yyy"
	fmt.Println(name)*/

	//print package level variable
	fmt.Println(package_level_var)

	//print Exported variable
	fmt.Println(pa1.PI)

	var a int
	a = 10
	fmt.Println(a)

	a = 20

	fmt.Println(a)

	//constants
	const PI = 3.145

	fmt.Println(PI)

	//Pointers
	var y int = 10

	/*var x *int = &y

	fmt.Println(&y)
	fmt.Println(x)

	fmt.Println(y)
	fmt.Println(*x)

	fmt.Println(&x)
	*/

	modify(&y)
	fmt.Println(y)

}

func modify(z *int) {
	*z = 30
	fmt.Println(*z)
}
