package jobs

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	u "backend/package/utils"
)

func PutScraped(collection *mongo.Collection, ctx context.Context) {
	fmt.Printf("[devotional/backend/jobs] scrape | start | %s \n", time.Now())
	scraped := u.GetScrapedUtil()
	u.PutDevotionalUtil(collection, ctx, scraped)
	fmt.Printf("[devotional/backend/jobs] scrape | finish | %s \n", time.Now())
}
