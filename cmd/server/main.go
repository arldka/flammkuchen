package main

import (
	"github.com/arldka/flammkuchen/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HandleIndex)
  http.HandleFunc("/search", handlers.HandleSearch)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
