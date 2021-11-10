package main

import (
	handler "backend/package/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/devotional", handler.GetDevotional)
	router.GET("/devotionals", handler.GetDevotionals)

	router.PUT("/devotional", handler.PutDevotional)

	// default port 8080
	// can be changed via env
	router.Run()
}
