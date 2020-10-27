package app

import (
	"log"
	"net/http"
	"productInventory/product"

	"github.com/gorilla/mux"
)

var (
	port      = ":8000"
)

// Start the application and load resources
func Start() {
	log.Printf("Starting web api on port %s", port)
	log.Fatal(http.ListenAndServe(port, loadRouter()))
}

func loadRouter() *mux.Router {
	log.Printf("Creating router...")
	router := mux.NewRouter().StrictSlash(true)
	product.LoadRoutes(router)
	return router
}
