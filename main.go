package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Product struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Desc string `json:"desc"`
}
type Products []Product

var cnt = 0

func Count() int {
	cnt++
	return cnt
}

var productsStore = Products{
	dummyProduct(),
}

func dummyProduct() Product {
	return Product{Name: "Dummy Product", Type: "Dummy Type", Desc: "Test Description"}
}
func allProducts(w http.ResponseWriter, r *http.Request) {


	fmt.Fprintf(w, "All Products\n")
	for _, product := range productsStore {
		json.NewEncoder(w).Encode(product)
	}
}
func singleProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Fprintf(w, "ID must be a valid Integer\n")
		return
	}
	fmt.Println("Data for ID ", id)
	for i, v := range productsStore {
		if i == id {
			json.NewEncoder(w).Encode(v)
			return
		}
	}
	fmt.Fprintf(w, "Id is not founded")
}

func createProduct(w http.ResponseWriter, r *http.Request) {

	var product Product
	json.NewDecoder(r.Body).Decode(&product)
	productsStore = append(productsStore, product)
	fmt.Println("created a product successfully")
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Fprintf(w, "Id must be a valid Integer\n")
		return
	}
	fmt.Println("Updating User", id)
	for i := range productsStore {
		if i == id {
			var product Product
			json.NewDecoder(r.Body).Decode(&product)
			productsStore[i] = product
			json.NewEncoder(w).Encode(product)
			fmt.Println("Update Product Successfully")
			return
		}
	}
	fmt.Fprintf(w, "Id is not founded")
}
func deleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		fmt.Fprintf(w, "Id must be a valid Integer\n")
		return
	}
	for i := range productsStore {
		if i == id {
			productsStore = append(productsStore[:i], productsStore[i + 1:]...)
			fmt.Println("Delete Product Successfully")
			return
		}
	}
	fmt.Println("Deleting User", id)

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}

func handleRquests() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.HandleFunc("/", homePage)
	r.Get("/products", allProducts)
	r.Post("/products", createProduct)
	r.Get("/products/{id}", singleProduct)
	r.Put("/products/{id}", updateProduct)
	r.Delete("/products/{id}", deleteProduct)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {
	handleRquests()
}
