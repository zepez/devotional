package handlers

import (
	"net/http"

	u "backend/package/utils"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDevotional(c *gin.Context, collection *mongo.Collection) {

	result := u.GetDevotionalUtil(collection)

	c.JSON(http.StatusOK, result)
}
