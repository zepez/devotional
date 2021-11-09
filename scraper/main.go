package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"

	u "scraper/package/utils"
)

func RootRoute(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	// set headers
	w.Header().Set("Content-Type", "application/json")

	// instantiate new devotional struct
	devotional := u.Devotional{}
	devotional.Target_Date = u.GenTargetDate(r.URL.Query())

	// respond with cached devotional if exists
	cachedContent, found := c.Get(devotional.Target_Date)
	if found {
		w.Write(u.GenJson(cachedContent))
		return
	}

	// get dovotional metadata
	devotional.Id = uuid.Must(uuid.NewRandom()).String()
	devotional.Created_at = time.Now()
	devotional.Updated_at = time.Now()
	devotional.Source = u.GenLink(devotional.Target_Date)

	// actual devotional content
	html, plain, name := u.Scrape(devotional.Source)

	devotional.Name = name
	devotional.Html = html
	devotional.Plain = plain

	// cache the newly scraped content
	c.Set(devotional.Target_Date, devotional, cache.DefaultExpiration)

	// response
	w.Write(u.GenJson(devotional))
	return
}

func handleRequests(c *cache.Cache) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		RootRoute(w, r, c)
	})

	// define port / log errors
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	// instantiate cache with default expiration
	c := cache.New(24*time.Hour, 10*time.Minute)

	// start message
	fmt.Println("Scraping server started on http://localhost:8080")

	// handle req and pass cache
	handleRequests(c)
}
