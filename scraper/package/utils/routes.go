package utils

import (
	"log"
	"net/http"

	"github.com/patrickmn/go-cache"
)

func HandleRequests(c *cache.Cache) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RootRoute(w, r, c)
	})

	http.HandleFunc("/health", HealthRoute)

	// define port / log errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}
