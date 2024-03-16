package main

import (
	"fmt"
)

func main() {
	/*max(1, 2, 3, 4, 5)
	max(20, 30, 40, 10)
	max()*/

	//comm, ok := divide(10, 3)

	/*if ok != nil {
		fmt.Println(ok)
	} else {
		fmt.Println(comm)
	}*/
	a := map[string]string{"apple": "2"}

	comm1, ok1 := a["apple"]

	if ok1 {
		fmt.Println(comm1)
	} else {
		fmt.Println("key not found")
	}

	

}

/*func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("a cannot divide by zero")
	}
	return a / b, nil
}*/

/*func divide(a, b int) (int, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic: ", r)
		}
	}()
	if b == 0 {
		panic("a cannot divide by zero")
	}

	return a / b, nil

}*/

/*func max(a ...int) {
	if len(a) == 0 {
		fmt.Println("no elements to compare")
	}else{
	maxValue := a[0]
	for _, val := range a[1:] {
		if val > maxValue {
			maxValue = val
		}
	}
	fmt.Println(maxValue)
}
}
*/
