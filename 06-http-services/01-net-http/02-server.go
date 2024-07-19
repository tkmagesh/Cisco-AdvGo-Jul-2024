package main

import (
	"encoding/json"
	"fmt"
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
	routes map[string]func(http.ResponseWriter, *http.Request)
}

func (as *AppServer) AddRoute(url string, handler func(http.ResponseWriter, *http.Request)) {
	as.routes[url] = handler
}

// http.Handler interface implementation
func (as *AppServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler := as.routes[r.URL.Path]; handler == nil {
		http.Error(w, "Resource not found", http.StatusNotFound)
		return
	} else {
		handler(w, r)
	}
}

// App Server factory
func NewAppServer() *AppServer {
	return &AppServer{
		routes: make(map[string]func(http.ResponseWriter, *http.Request)),
	}
}

// Application specific handler functions
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Processing request for /"))
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
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
}

func CustomersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "All the customers data will be served /")
}

func main() {
	appServer := NewAppServer()
	appServer.AddRoute("/", IndexHandler)
	appServer.AddRoute("/products", ProductsHandler)
	appServer.AddRoute("/customers", CustomersHandler)
	http.ListenAndServe(":8080", appServer)
}