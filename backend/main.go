package main

import (
	handler "backend/package/handlers"
	jobs "backend/package/jobs"
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// create router
	router := gin.Default()

	// create database client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))
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

	// route definitions
	router.GET("/devotional", func(c *gin.Context) { handler.GetDevotional(c, collection) })
	router.GET("/devotionals", func(c *gin.Context) { handler.GetDevotionals(c, collection) })

	// start jobs
	job := cron.New()
	fmt.Println("[devotional/backend/jobs] scheduling | scrape | 0 * * * * *")
	job.AddFunc("@every 30s", func() { jobs.PutScraped(collection, ctx) })
	job.Start()
	// defer job.Stop()

	// default port 8080 can be changed via env
	router.Run()
}
