package utils

import (
	def "backend/package/definitions"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PutDevotionalUtil(collection *mongo.Collection, ctx context.Context, devotional def.Devotional) *mongo.InsertOneResult {
	bson, err := bson.Marshal(devotional)
	if err != nil {
		log.Fatalln(err)
	}

	response, err := collection.InsertOne(ctx, bson)
	if err != nil {
		log.Fatalln(err)
	}

	return response
}
