package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	def "backend/package/definitions"
)

func GetDevotionals(c *gin.Context, collection *mongo.Collection) {
	var results []def.Devotional

	findOptions := options.Find()
	findOptions.SetLimit(5)

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var res def.Devotional
		err := cursor.Decode(&res)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, res)
	}

	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, results)
}
