package main

import (
	handler "backend/package/handlers"
	jobs "backend/package/jobs"
	u "backend/package/utils"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// start message
	fmt.Printf("[devotional/backend/server] server | starting | %s \n", time.Now())

	// create database client
	client, err := mongo.NewClient(options.Client().ApplyURI(u.GetEnvUtil("DEVOTIONAL_BACKEND_MONGO_URI", "mongodb://127.0.0.1:27017/")))
	if err != nil {
		log.Fatal(err)
	}

	// create context
	ctx := context.Background()

	// connect client
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// define collection
	collection := client.Database("devotional").Collection("devotional")

	// start jobs
	crons := cron.New()
	freq := u.GetEnvUtil("DEVOTIONAL_BACKEND_SCRAPER_FREQ", "@every 30s")
	fmt.Printf("[devotional/backend/jobs] scrape | scheduling %s | %s \n", freq, time.Now())
	crons.AddFunc(freq, func() { jobs.PutScraped(collection, ctx) })
	crons.Start()
	defer crons.Stop()

	// create router
	router := gin.Default()

	// route definitions
	router.GET("/devotional/:id", func(c *gin.Context) { handler.GetDevotional(c, collection) })
	router.GET("/devotionals/:page", func(c *gin.Context) { handler.GetDevotionals(c, collection) })
	router.GET("/health", func(c *gin.Context) { handler.GetHealth(c) })

	// default port 8080 can be changed via env
	router.Run(":" + u.GetEnvUtil("PORT", "8080"))
}
