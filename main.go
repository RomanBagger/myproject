package main

import (
	"encoding/json"
	"log"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

type Product struct {
	Id       string
	Name     string
	Quantity int
	Price    float64
}

var Products []Product

func homepage(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/homepage.html")
	product := Product{"1", "Mango", 14, 0.99}
	tmpl.Execute(w, product)
	log.Println("Endpoint hit: homepage")
	//fmt.Fprintf(w, "Welcome to homepage")
}

func returnALLproduct(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit: returnALLproduct")
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	for _, product := range Products {
		if product.Id == key {
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(product); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}
	http.Error(w, "Product not found", http.StatusNotFound)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/products/", returnALLproduct).Methods("GET")
	myRouter.HandleFunc("/product/{id}", getProduct).Methods("GET")
	myRouter.HandleFunc("/", homepage).Methods("GET")
	log.Println("Starting server on :10000")
	if err := http.ListenAndServe("localhost:10000", myRouter); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}

func main() {
	Products = []Product{
		{Id: "1", Name: "Chair", Quantity: 100, Price: 100.00},
	}

	handleRequests()
}
