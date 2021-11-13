package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"

	handler "scraper/package/handlers"
	u "scraper/package/utils"
)

func main() {
	// start message
	fmt.Printf("[devotional/backend/server] server | starting | %s \n", time.Now())

	// instantiate cache with default expiration
	cache := cache.New(24*time.Hour, 10*time.Minute)
	fmt.Printf("[devotional/backend/server] cache | created | %s \n", time.Now())

	// create router
	router := gin.Default()

	// define routes
	router.GET("/health", func(c *gin.Context) { handler.GetHealth(c) })
	router.GET("/", func(c *gin.Context) { handler.GetRoot(c, cache) })

	// start server
	router.Run(":" + u.GetEnvUtil("PORT", "8081"))
}
