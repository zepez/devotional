package utils

import (
	"context"
	"log"

	def "backend/package/definitions"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDevotionalUtil(collection *mongo.Collection) def.Devotional {
	var result def.Devotional

	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}
