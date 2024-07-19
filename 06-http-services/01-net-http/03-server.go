/*
CURL
curl http://localhost:8080/
curl http://localhost:8080/products
curl http://localhost:8080/products -X POST -H "Content-Type: application/json" -d '{"id" : 104, "name" : "scribble pad", "cost" : 20}'
curl http://localhost:8080/customers
curl http://localhost:8080/users
*/
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

type Middleware func(http.HandlerFunc) http.HandlerFunc

type AppServer struct {
	routes      map[string]http.HandlerFunc
	middlewares []Middleware
}

func (as *AppServer) AddRoute(url string, handler http.HandlerFunc) {
	for _, middleware := range as.middlewares {
		handler = middleware(handler)
	}
	as.routes[url] = handler
}

func (as *AppServer) AddMiddleware(middleware Middleware) {
	as.middlewares = append(as.middlewares, middleware)
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
		routes: make(map[string]http.HandlerFunc),
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

// custom middlewares
func logMiddleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		handler(w, r)
	}
}

func main() {
	appServer := NewAppServer()
	appServer.AddMiddleware(logMiddleware)
	appServer.AddRoute("/", IndexHandler)
	appServer.AddRoute("/products", ProductsHandler)
	appServer.AddRoute("/customers", CustomersHandler)
	http.ListenAndServe(":8080", appServer)
}
