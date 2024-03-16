package main

import (
	"fmt"
	"sync"
)

func main() {
	/*fmt.Println("hello")

	unbuffered := make(chan int)

	go func() {
		unbuffered <- 42

	}()



	fmt.Println(<-unbuffered)
	fmt.Println(<-unbuffered)*/

	/*buffered := make(chan int, 4)

	go func() {
		buffered <- 23
		buffered <- 34
		buffered <- 56
		buffered <- 66
		buffered <- 76

	}()
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)*/

	//ch1 := make(chan string)
	//ch2 := make(chan string)

	/*go func() {
		ch1 <- "hello"
	}()

	go func() {
		ch2 <- "world"
	}()*/

	/*select {
	case msg1 := <-ch1:
		fmt.Println(msg1)
	case msg2 := <-ch2:
		fmt.Println(msg2)
	default:
		fmt.Println("no data on channel")
	}*/
	var wg sync.WaitGroup
	sendOnlych := make(chan<- int, 2)
	recieveOnlych := make(<-chan int, 1)

	wg.Add(1)
	go sendData(sendOnlych, 42, &wg)

	wg.Add(1)
	go recievData(recieveOnlych, &wg)

	wg.Wait()
}

func sendData(ch chan<- int, data int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- data
}

func recievData(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := <-ch
	fmt.Println("result:", result)
}
