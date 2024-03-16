package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)
	out := make(chan int)

	//Producer
	wg.Add(1)
	go producer(ch, &wg)

	numconsumer := 3
	for i := 0; i < numconsumer; i++ {
		wg.Add(1)
		go consumer(ch, out, &wg)
	}
	go collector(out, &wg)
	wg.Wait()
}

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range in {
		result := num * 2
		out <- result
	}
}

func collector(out <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for result := range out {
		fmt.Println("result:", result)
	}
}
