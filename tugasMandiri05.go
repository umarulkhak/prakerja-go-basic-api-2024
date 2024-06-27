package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// Product struct represents a product entity
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

var (
	products []Product
	nextID   = 1
	mutex    sync.RWMutex
)

func main() {
	http.HandleFunc("/products", handleProducts)
	http.HandleFunc("/products/", handleProductByID)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getProducts(w, r)
	case http.MethodPost:
		createProduct(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	mutex.RLock()
	defer mutex.RUnlock()
	json.NewEncoder(w).Encode(products)
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	mutex.Lock()
	defer mutex.Unlock()
	product.ID = nextID
	nextID++
	products = append(products, product)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func handleProductByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodPut:
		updateProductByID(w, r, id)
	case http.MethodDelete:
		deleteProductByID(w, r, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func updateProductByID(w http.ResponseWriter, r *http.Request, id int) {
	var updatedProduct Product
	err := json.NewDecoder(r.Body).Decode(&updatedProduct)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	found := false
	for i, p := range products {
		if p.ID == id {
			products[i] = updatedProduct
			updatedProduct.ID = id
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedProduct)
}

func deleteProductByID(w http.ResponseWriter, r *http.Request, id int) {
	mutex.Lock()
	defer mutex.Unlock()
	found := false
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
