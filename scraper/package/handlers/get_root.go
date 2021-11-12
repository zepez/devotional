package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"

	define "scraper/package/definitions"
	u "scraper/package/utils"
)

func GetRoot(c *gin.Context, memcache *cache.Cache) {

	// instantiate new devotional struct
	devotional := define.Devotional{}
	devotional.Target_Date = u.GenTargetDate(c.Query("date"))

	// respond with cached devotional if exists
	cachedContent, found := memcache.Get(devotional.Target_Date)
	if found {
		c.JSON(http.StatusOK, cachedContent)
		return
	}

	// get dovotional metadata
	devotional.Created_at = time.Now()
	devotional.Updated_at = time.Now()
	devotional.Source = u.GenLink(devotional.Target_Date)

	// actual devotional content
	html, plain, name, image := u.Scrape(devotional.Source)

	devotional.Name = name
	devotional.Html = html
	devotional.Plain = plain
	devotional.Image = image

	// cache the newly scraped content
	memcache.Set(devotional.Target_Date, devotional, cache.DefaultExpiration)

	// response
	c.JSON(http.StatusOK, devotional)
	return
}
