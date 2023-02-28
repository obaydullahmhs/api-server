package api

import (
	"api-server/auth"
	"api-server/data"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func allSales(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "All Sales Order\n")
	for _, product := range data.SalesStore {
		json.NewEncoder(w).Encode(product)
	}
}

func saleForID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	fmt.Println("Data for ID ", id)
	ok := false
	for _, v := range data.SalesStore {
		if v.ID == id {
			json.NewEncoder(w).Encode(v)
			ok = true
		}
	}
	if !ok {
		fmt.Fprintf(w, "Id is not founded")
	}
	w.WriteHeader(http.StatusOK)
}

func createSale(w http.ResponseWriter, r *http.Request) {

	var sale data.Sale
	json.NewDecoder(r.Body).Decode(&sale)
	data.
	SalesStore = append(data.SalesStore, sale)
	fmt.Println("created a Sale successfully")
	w.WriteHeader(http.StatusCreated)
}

func updateSale(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	fmt.Println("Updating ID", id)
	for i, v := range data.SalesStore {
		if v.ID == id {
			var sale data.Sale
			json.NewDecoder(r.Body).Decode(&sale)
			data.SalesStore[i] = sale
			json.NewEncoder(w).Encode(sale)
			fmt.Println("Updating a Sale Successfully")
			return
		}
	}
	fmt.Fprintf(w, "Id is not founded")
	w.WriteHeader(http.StatusOK)
}

func deleteSale(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	for i, v := range data.SalesStore {
		if v.ID == id {
			data.SalesStore = append(data.SalesStore[:i], data.SalesStore[i+1:]...)
			fmt.Println("Delete Product Successfully")
			return
		}
	}
	fmt.Println("Deleting User", id)
	w.WriteHeader(http.StatusAccepted)

}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage")
}


var jwtsecretkey = []byte("abcdef")

func StartServer(port int) {
	fmt.Printf("Starting server at %v\n", port)
	fmt.Printf("Sever Link: http://localhost:%v\n", port)
	fmt.Println(auth.GenerateJWT(jwtsecretkey))
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.HandleFunc("/", auth.IsAuthorized(homePage))
	r.Get("/sales", auth.IsAuthorized(allSales))
	r.Post("/sales", auth.IsAuthorized(createSale))
	r.Get("/sales/{id}", auth.IsAuthorized(saleForID))
	r.Put("/sales/{id}", auth.IsAuthorized(updateSale))
	r.Delete("/sales/{id}", auth.IsAuthorized(deleteSale))
	log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), r))

}