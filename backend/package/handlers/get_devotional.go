package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDevotional(c *gin.Context) {
	c.String(http.StatusOK, "GET /devotional")
}
