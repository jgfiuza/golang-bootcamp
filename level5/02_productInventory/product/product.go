package product

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

var inventory []ProductInventory

func init() {
	for i := 1; i < 10; i++ {
		iStr := strconv.Itoa(i)
		newProduct := Product{ID: iStr, Code: "100" + iStr, Name: "IPhone " + iStr, Price: float64(i) * 100}
		newProductInventory := ProductInventory{Product: newProduct, Quantity: i}
		inventory = append(inventory, newProductInventory)
	}
}

func findProductInventoryByID(id string) (int, ProductInventory, error) {
	for idx, singleProductInventory := range inventory {
		if id == singleProductInventory.Product.ID {
			return idx, singleProductInventory, nil
		}
	}
	return 0, ProductInventory{}, errors.New("Couldn't find ProductInventory matching ID")
}

func createProduct(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newProductInventory ProductInventory

	json.Unmarshal(reqBody, &newProductInventory)
	inventory = append(inventory, newProductInventory)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newProductInventory)
}

func getAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(inventory)
}

func getOneProductByID(w http.ResponseWriter, r *http.Request) {
	pathVariable := mux.Vars(r)
	productID := pathVariable["id"]
	_, foundProductInventory, err := findProductInventoryByID(productID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"Couldn't find ProductInventory with matching ID."}`))
		return
	}

	json.NewEncoder(w).Encode(foundProductInventory)
}

func updateProduct(w http.ResponseWriter, r *http.Request) {
	pathVariable := mux.Vars(r)
	productID := pathVariable["id"]
	var updatedProductInventory ProductInventory

	idx, _, err := findProductInventoryByID(productID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"Couldn't find ProductInventory with matching ID."}`))
		return
	}

	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedProductInventory)
	inventory[idx] = updatedProductInventory
	json.NewEncoder(w).Encode(updatedProductInventory)
}

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	pathVariable := mux.Vars(r)
	productID := pathVariable["id"]

	idx, _, err := findProductInventoryByID(productID)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error":"Couldn't find ProductInventory with matching ID."}`))
		return
	}

	inventory = append(inventory[:idx], inventory[idx+1:]...)
}

// LoadRoutes loads products resources into router
func LoadRoutes(router *mux.Router) {
	log.Printf("Loading Product routes...")
	router.HandleFunc("/product", getAllProducts).Methods(http.MethodGet)
	router.HandleFunc("/product", createProduct).Methods(http.MethodPost)
	router.HandleFunc("/product/{id}", getOneProductByID).Methods(http.MethodGet)
	router.HandleFunc("/product/{id}", updateProduct).Methods(http.MethodPut)
	router.HandleFunc("/product/{id}", deleteProduct).Methods(http.MethodDelete)
}
