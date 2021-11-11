package jobs

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	u "backend/package/utils"
)

func PutScraped(collection *mongo.Collection, ctx context.Context) {

	scraped := u.GetScrapedUtil()

	results := u.PutDevotionalUtil(collection, ctx, scraped)

	fmt.Printf("New devotional created with id %s \n", results.Id)
}
