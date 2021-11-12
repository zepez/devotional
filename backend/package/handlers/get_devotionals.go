package handlers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	u "backend/package/utils"
)

func GetDevotionals(c *gin.Context, collection *mongo.Collection) {
	page := c.Param("page")

	results, status := u.GetDevotionalsUtil(collection, page)

	if status != 200 {
		c.Writer.WriteHeader(status)
		return
	}

	c.JSON(status, results)
}
