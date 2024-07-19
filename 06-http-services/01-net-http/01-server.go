package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Product struct {
	Id   int     `json:"id"`
	Name string  `json:"name"`
	Cost float64 `json:"cost"`
}

// in-memory sample data
var products []Product = []Product{
	{Id: 101, Name: "Pen", Cost: 10},
	{Id: 102, Name: "Pencil", Cost: 5},
	{Id: 103, Name: "Marker", Cost: 50},
}

type AppServer struct {
}

// http.Handler interface implementation
func (as *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL.Path)
	switch r.URL.Path {
	case "/":
		// fmt.Fprintln(w, "Processing request for /")
		w.Write([]byte("Processing request for /"))
	case "/products":
		switch r.Method {
		case http.MethodGet:
			if err := json.NewEncoder(w).Encode(products); err != nil {
				http.Error(w, "Error processing the request", http.StatusInternalServerError)
			}
		case http.MethodPost:
			var newProduct Product
			if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
				http.Error(w, "invalid payload", http.StatusBadRequest)
				return
			}
			products = append(products, newProduct)
			http.Error(w, "Object Created", http.StatusCreated)
		default:
			http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		}

	case "/customers":
		fmt.Fprintln(w, "All the customers data will be served /")
	default:
		// w.WriteHeader(http.StatusNotFound)
		http.Error(w, "Resource not found", http.StatusNotFound)
	}

}

func main() {
	appServer := &AppServer{}
	http.ListenAndServe(":8080", appServer)
}
