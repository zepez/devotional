package handlers

import (
	u "backend/package/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDevotional(c *gin.Context, collection *mongo.Collection) {
	id := c.Param("id")

	result, status := u.GetDevotionalUtil(collection, id)

	if status != 200 {
		c.Writer.WriteHeader(status)
		return
	}

	c.JSON(status, result)
}
