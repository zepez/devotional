package utils

import (
	def "backend/package/definitions"
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetDevotionalsUtil(collection *mongo.Collection, page string) ([]def.Devotional, int) {
	var results []def.Devotional

	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return results, 400
	}

	findOptions := options.Find()
	findOptions.SetLimit(3)
	findOptions.SetSkip(pageInt)

	cursor, err := collection.Find(context.TODO(), bson.D{}, findOptions)
	if err != nil {
		return results, 500
	}

	for cursor.Next(context.TODO()) {
		var res def.Devotional
		err := cursor.Decode(&res)
		if err != nil {
			return results, 500
		}

		results = append(results, res)
	}

	if results == nil {
		return results, 404
	}

	return results, 200
}
