package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	def "backend/package/definitions"
)

func GetDevotional(c *gin.Context, collection *mongo.Collection) {
	var result def.Devotional

	err := collection.FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}
