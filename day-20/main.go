package main

import (
	"fmt"
	"net/http"
)

/*var (
	balance int
	mutex   sync.Mutex
)*/

func main() {
	/*var wg sync.WaitGroup
	ch := make(chan int)
	sendonlychan := (chan<- int)(ch)
	recieveonlych := (<-chan int)(ch)

	wg.Add(2)

	go senddata(sendonlychan, 42, &wg)
	go recievedata(recieveonlych, &wg)

	wg.Wait()*/
	/*var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go incrementcounter(&wg)
	}
	wg.Wait()*/

	/*var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			deposit(100)
			time.Sleep(time.Millisecond * 300)
			withdraw(50)
		}(i)
	}
	wg.Wait()*/

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("server listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error:", err)
	}

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to home page")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to about page")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "contact us at contact@gmail.com")
}

/*func deposit(amount int) {
	mutex.Lock()
	defer mutex.Unlock()
	balance += amount
	fmt.Printf("deposite:%d,total_cmount:%d\n", amount, balance)
}
func withdraw(amount int) {
	mutex.Lock()
	defer mutex.Unlock()
	if balance >= amount {
		balance -= amount
		fmt.Printf("withdraw:%d,new_balance:%d\n", amount, balance)
	} else {
		fmt.Println("Insufficient amount")
	}
}*/

/*func incrementcounter(wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()

	counter++
	fmt.Println("counter value:", counter)

}*/

/*func senddata(ch chan<- int, data int, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- data

}
func recievedata(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	result := <-ch
	fmt.Println("result:", result)

}*/
