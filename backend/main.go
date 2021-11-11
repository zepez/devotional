package main

import (
	handler "backend/package/handlers"
	jobs "backend/package/jobs"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// create router
	router := gin.Default()

	// create database client
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	// create context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

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

	jobs.PutScraped(collection, ctx)

	// default port 8080 can be changed via env
	router.Run()
}
