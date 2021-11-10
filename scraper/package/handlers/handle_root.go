package handlers

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/patrickmn/go-cache"

	u "scraper/package/utils"
)

func Root(w http.ResponseWriter, r *http.Request, c *cache.Cache) {
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
