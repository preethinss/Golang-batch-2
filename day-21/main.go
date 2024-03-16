package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Item struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var item []Item

func main() {
	http.HandleFunc("/items", getAllitems)
	http.HandleFunc("/items/add", addItem)
	http.HandleFunc("/items/delete", deleteItem)

	fmt.Println("server is started and listening on port 8080 ")
	http.ListenAndServe(":8080", nil)
}
func addItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	err := json.NewDecoder(r.Body).Decode(&newItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	item = append(item, newItem)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)

}

func getAllitems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	if id == "" {
		http.Error(w, "pleae provide an id", http.StatusBadRequest)
		return
	}
	var deleteItem Item
	for i, data := range item {
		if data.Id == id {
			deleteItem = item[i]
			item = append(item[:i], item[i+1:]...)
			break
		}
	}
	if deleteItem.Id == "" {
		http.Error(w, "item not found", http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(deleteItem)
}
