package app

import (
	"log"
	"net/http"
)

func Start() {
	//define routes
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", main.greet)
	mux.HandleFunc("/customers", main.getAllCustomers)

	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}
