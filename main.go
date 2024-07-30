package main

import (
	"log"
	"net/http"

	handlers "github.com/sawilkhan/analytics/handlers/geo"
)

func main() {
	http.HandleFunc("/geo", handlers.GeoHandler)
	log.Println("starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
