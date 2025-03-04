package main

import (
	"github.com/arldka/flammkuchen/handlers"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", handlers.HandleIndex)
	r.HandleFunc("/search", handlers.HandleSearch)
	r.HandleFunc("/ks/{kustomizationName}/{namespace}", handlers.HandleKustomization)
	r.HandleFunc("/hr/{helmReleaseName}/{namespace}", handlers.HandleHelmRelease)
	http.Handle("/", r)
	log.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
