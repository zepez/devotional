package utils

import (
	def "backend/package/definitions"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDevotionalsUtil(collection *mongo.Collection) []def.Devotional {
	var results []def.Devotional

	findOptions := options.Find()
	findOptions.SetLimit(5)

	cursor, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.TODO()) {
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

	return results
}
