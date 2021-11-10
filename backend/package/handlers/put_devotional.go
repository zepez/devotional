package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PutDevotional(c *gin.Context) {
	c.String(http.StatusOK, "PUT /devotional")
}
