package main

import (
	"fmt"
	"time"
)

func main() {
	go taskA()
	go taskB()

	time.Sleep(time.Second * 5)
}
func taskA() {
	for i := 1; i <= 8; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}
func taskB() {
	for i := 9; i <= 12; i++ {
		fmt.Println(i)
		time.Sleep(time.Millisecond * 300)
	}
}
