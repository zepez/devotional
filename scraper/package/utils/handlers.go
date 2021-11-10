package utils

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"
)

func HealthRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	return
}

func RootRoute(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
	// set headers
	w.Header().Set("Content-Type", "application/json")

	// instantiate new devotional struct
	devotional := Devotional{}
	devotional.Target_Date = GenTargetDate(r.URL.Query())

	// respond with cached devotional if exists
	cachedContent, found := c.Get(devotional.Target_Date)
	if found {
		w.Write(GenJson(cachedContent))
		return
	}

	// get dovotional metadata
	devotional.Id = uuid.Must(uuid.NewRandom()).String()
	devotional.Created_at = time.Now()
	devotional.Updated_at = time.Now()
	devotional.Source = GenLink(devotional.Target_Date)

	// actual devotional content
	html, plain, name := Scrape(devotional.Source)

	devotional.Name = name
	devotional.Html = html
	devotional.Plain = plain

	// cache the newly scraped content
	c.Set(devotional.Target_Date, devotional, cache.DefaultExpiration)

	// response
	w.Write(GenJson(devotional))
	return
}
