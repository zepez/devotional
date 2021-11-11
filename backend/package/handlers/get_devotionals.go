package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"

	u "backend/package/utils"
)

func GetDevotionals(c *gin.Context, collection *mongo.Collection) {

	results := u.GetDevotionalsUtil(collection)

	c.JSON(http.StatusOK, results)
}
