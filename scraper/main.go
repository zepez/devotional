package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	handle "scraper/package/handlers"

	"github.com/patrickmn/go-cache"
)

func main() {
	// instantiate cache with default expiration
	c := cache.New(24*time.Hour, 10*time.Minute)

	// start message
	fmt.Println("Scraping server started on http://localhost:8080")

	// root, scraping route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { handle.Root(w, r, c) })

	// healthcheck route
	http.HandleFunc("/health", handle.Health)

	// define port / log errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}
