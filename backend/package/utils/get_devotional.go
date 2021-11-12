package utils

import (
	"context"

	def "backend/package/definitions"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDevotionalUtil(collection *mongo.Collection, id string) (def.Devotional, int) {
	var result def.Devotional

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return result, 400
	}

	// actual query
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		return result, 404
	} else if err != nil {
		return result, 500
	}

	return result, 200
}
