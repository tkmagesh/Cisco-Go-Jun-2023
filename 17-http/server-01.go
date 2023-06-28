package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	Id       int     `json:"id"`
	Name     string  `json:"name"`
	Cost     float32 `json:"cost"`
	Units    int     `json:"-"`
	Category string  `json:"category"`
}

type Products []Product

var products = Products{
	Product{105, "Pen", 5, 50, "Stationary"},
	Product{107, "Pencil", 2, 100, "Stationary"},
	Product{103, "Marker", 50, 20, "Utencil"},
	Product{102, "Stove", 5000, 5, "Utencil"},
	Product{101, "Kettle", 2500, 10, "Utencil"},
	Product{104, "Scribble Pad", 20, 20, "Stationary"},
	Product{109, "Golden Pen", 2000, 20, "Stationary"},
}

type AppServer struct {
}

// http.Handler interface implementation
func (appServer *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintln(w, "Hello World!")
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}

	default:
		http.Error(w, "Resource not found!", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
