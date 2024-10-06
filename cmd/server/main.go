package main

import (
	"github.com/arldka/flammkuchen/handlers"
	"log"
	"net/http"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()
	r.HandleFunc("/", handlers.HandleIndex)
  r.HandleFunc("/search", handlers.HandleSearch)
  r.HandleFunc("/ks/{kustomizationName}/{namespace}", handlers.HandleKustomization).Methods("GET")
  http.Handle("/", r)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
