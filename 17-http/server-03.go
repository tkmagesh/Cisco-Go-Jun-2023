package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float32 `json:"cost"`
}

var products []Product = []Product{
	{Id: 101, Name: "Pen", Cost: 5},
	{Id: 102, Name: "Pencil", Cost: 2},
	{Id: 103, Name: "Marker", Cost: 15},
}

// handlers
func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":8080", nil)
}
