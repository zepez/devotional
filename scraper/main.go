package main

import (
	"fmt"
	"time"

	u "scraper/package/utils"

	"github.com/patrickmn/go-cache"
)

func main() {
	// instantiate cache with default expiration
	c := cache.New(24*time.Hour, 10*time.Minute)

	// start message
	fmt.Println("Scraping server started on http://localhost:8080")

	// handle req and pass cache
	u.HandleRequests(c)
}
