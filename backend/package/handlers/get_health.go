package handlers

import (
	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.Writer.WriteHeader(200)
}
