package product

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// Product holdes the information of products in the inventory
type Product struct {
	ID    string  `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// ProductInventory holdes the inventory of products
type ProductInventory struct {
	Product  Product
	Quantity int `json:"qty"`
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	// implement
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	// implement
}

func getOneProductById(w http.ResponseWriter, r *http.Request) {
	//pathVariable := mux.Vars(r)
	//productID := pathVariable["id"]
	// implement
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	// implement
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// implement
}

// LoadRoutes loads products resources into router
func LoadRoutes(router *mux.Router) {
	log.Printf("Loading Product routes...")
	router.HandleFunc("/product", getAllProducts).Methods("GET")
	router.HandleFunc("/product", createProduct).Methods("POST")
	router.HandleFunc("/product/:id", getOneProductById).Methods("GET")
	router.HandleFunc("/product/:id", updateProduct).Methods("PUT")
	router.HandleFunc("/product/:id", deleteProduct).Methods("DELETE")
}
