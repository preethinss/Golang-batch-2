package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*fmt.Println("hello")

	file, err := os.Create("sample.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("file is created")
	}

	defer file.Close()

	content := []byte("this is content")
	_, err1 := file.Write(content)

	if err1 != nil {
		fmt.Println(err1)
	}*/

	/*fileinfo, err := os.Stat("sample.txt")

	if os.IsNotExist(err) {
		fmt.Println("file not exist")
	} else {
		content, err := os.ReadFile(fileinfo.Name())
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(content))
		}
	}*/

	/*err := os.Rename("sample.txt", "dest.txt")
	if err != nil {
		fmt.Println(err)
	}*/

	/*err := os.Remove("dest.txt")
	if err != nil {
		fmt.Println(err)
	}*/

	/*router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"meesage": "hello",
		})
	})
	router.Run(":8080")*/

	sum := 0
	for _, v := range os.Args[1:] {
		fmt.Printf("%T", v)
		result, _ := strconv.Atoi(v)
		sum += result

	}

	fmt.Println(sum)

}
