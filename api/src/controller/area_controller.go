package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetArea(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"area": "area",
	})
}