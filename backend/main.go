package main

import (
	handler "backend/package/handlers"
	jobs "backend/package/jobs"
	"context"
	"fmt"
	"log"
	"os"
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
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("DEVOTIONAL_BACKEND_MONGO_URI")))
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
	job := cron.New()
	fmt.Printf("[devotional/backend/jobs] scrape | scheduling %s | %s \n", os.Getenv("DEVOTIONAL_BACKEND_SCRAPER_FREQ"), time.Now())
	job.AddFunc(os.Getenv("DEVOTIONAL_BACKEND_SCRAPER_FREQ"), func() { jobs.PutScraped(collection, ctx) })
	job.Start()
	// defer job.Stop()

	// create router
	router := gin.Default()

	// route definitions
	router.GET("/devotional/:id", func(c *gin.Context) { handler.GetDevotional(c, collection) })
	router.GET("/devotionals/:page", func(c *gin.Context) { handler.GetDevotionals(c, collection) })
	router.GET("/health", func(c *gin.Context) { handler.GetHealth(c) })

	// default port 8080 can be changed via env
	router.Run()
}
