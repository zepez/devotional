package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDevotionals(c *gin.Context) {
	c.String(http.StatusOK, "GET /devotionals")
}
