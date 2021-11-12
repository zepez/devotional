package utils

import (
	def "backend/package/definitions"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func PutDevotionalUtil(collection *mongo.Collection, ctx context.Context, devotional def.Devotional) {
	// convert input devotional to bson
	bsonDevotional, err := bson.Marshal(devotional)
	if err != nil {
		log.Fatalln(err)
	}

	// check if existing devotional already exists
	var existingDevotional def.Devotional
	err = collection.FindOne(ctx, bsonDevotional).Decode(&existingDevotional)

	// log if existing is found
	if existingDevotional.Source != "" {
		fmt.Printf("[devotional/backend/utils] create | abort | %s \n", time.Now())
	}

	// only insert devotional if no document matches
	if err == mongo.ErrNoDocuments {
		fmt.Printf("[devotional/backend/utils] create | insert | %s \n", time.Now())
		_, err = collection.InsertOne(ctx, bsonDevotional)
	} else if err != nil {
		log.Fatal(err)
	}

	return
}
