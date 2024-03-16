package main

import (
	"fmt"
	"time"
)

func main() {

	channel := make(chan string)

	go func() {
		channel <- "hello"
	}()

	message := <-channel
	fmt.Println(message)

	/*go printNumbers()
	go printNumbers()*/

	/*var wg sync.WaitGroup
	urls := []string{"https://google.com", "https://google.com", "http://google.com"}
	startTime := time.Now()

	for _, url := range urls {
		wg.Add(1)
		go handleRequest(&wg, url)
	}
	wg.Wait()
	concurrenttime := time.Since(startTime)
	fmt.Printf("time taken for entire logic:%s\n", concurrenttime)*/
	/*urls := []string{"https://google.com", "https://google.com", "http://google.com"}
	startTime := time.Now()

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching %s: \n", err)
			continue
		}
		defer resp.Body.Close()
		fmt.Printf("Status code for %s and %d\n", url, resp.StatusCode)
	}

	nonconcurrenttime := time.Since(startTime)

	fmt.Printf("non concurrent approach took:%s\n", nonconcurrenttime)*/

}

func printNumbers() {

	for i := 1; i <= 5; i++ {
		fmt.Println("i value", i)
		time.Sleep(time.Millisecond * 300)
	}

}

/*func handleRequest(wg *sync.WaitGroup, url string) {
	defer wg.Done()

	starttime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching %s:%v\n", url, err)
	}
	defer resp.Body.Close()
	time_taken := time.Since(starttime)
	fmt.Printf("time taken: %s\n", time_taken)
}*/
