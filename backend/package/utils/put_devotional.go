package utils

import (
	def "backend/package/definitions"
	"context"
	"fmt"
	"log"

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
	if existingDevotional.Id != "" {
		fmt.Println("[devotional/backend/utils] put | existing devotional found, aborting")
	}

	// only insert devotional if no document matches
	if err == mongo.ErrNoDocuments {
		fmt.Println("[devotional/backend/utils] put | no existing devotional found, inserting")
		_, err = collection.InsertOne(ctx, bsonDevotional)
	} else if err != nil {
		log.Fatal(err)
	}

	return
}
