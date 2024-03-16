package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var item []Item

func main() {
	r := gin.Default()

	r.GET("/items", getAllItems)
	r.POST("/items/add", addItem)
	r.DELETE("/items/delete", deleteItem)

	fmt.Println("server is listening on port 8080")
	r.Run(":8080")

}
func getAllItems(c *gin.Context) {
	c.JSON(200, item)
}

func addItem(c *gin.Context) {
	var newitem Item
	err := c.BindJSON(&newitem)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	item = append(item, newitem)
	c.JSON(201, newitem)
}
func deleteItem(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.JSON(400, gin.H{"error": "please provifde id parameter"})
		return
	}
	var deleteditem Item
	for i, it := range item {
		if it.ID == id {
			deleteditem = item[i]
			item = append(item[:i], item[i+1:]...)
			break
		}
	}
	if deleteditem.ID == "" {
		c.JSON(404, gin.H{"error": "item not found"})
	}
	c.JSON(200, deleteditem)
}
