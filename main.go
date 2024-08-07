package main

import (
	"log"
	"net/http"

	"github.com/javier-elizaga/go-api/handlers"
)

func main() {
	http.HandleFunc("/live", handlers.GetLive)
	http.HandleFunc("/api/v1/users", handlers.GetUsers)
	http.HandleFunc("/api/v1/nearby", handlers.NearbyUser)
	log.Println("Starting server on 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
