package jobs

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"

	u "backend/package/utils"
)

func PutScraped(collection *mongo.Collection, ctx context.Context) {
	fmt.Println("[devotional/backend/jobs] running | scrape | 0 * * * * *")
	scraped := u.GetScrapedUtil()
	u.PutDevotionalUtil(collection, ctx, scraped)
	fmt.Println("[devotional/backend/jobs] finish | scrape | 0 * * * * *")
}
