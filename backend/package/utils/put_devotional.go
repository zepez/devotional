package utils

import (
	def "backend/package/definitions"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PutDevotionalUtil(collection *mongo.Collection, ctx context.Context, devotional def.Devotional) {
	// convert input devotional to bson
	bsonDevotional, err := bson.Marshal(devotional)
	if err != nil {
		fmt.Println(err)
		return
	}

	// // check if existing devotional already exists
	var existingDevotional def.Devotional
	err = collection.FindOne(ctx, bson.M{"target_date": devotional.Target_Date}).Decode(&existingDevotional)

	// log if existing is found
	if existingDevotional.Source != "" {
		fmt.Printf("[devotional/backend/utils] create | abort | %s \n", time.Now())
	}

	// only insert devotional if no document matches
	if err == mongo.ErrNoDocuments {
		fmt.Printf("[devotional/backend/utils] create | insert | %s \n", time.Now())
		_, err = collection.InsertOne(ctx, bsonDevotional)
	} else if err != nil {
		fmt.Println(err)
		return
	}

	return
}
