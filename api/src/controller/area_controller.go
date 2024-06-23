package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"gorm.io/gorm"
)

type Area struct {
	ID    uint   `json:”id”`
	Name string   `json:”name”`
	URL string   `json:”url”`
}

func GetArea(c *gin.Context, db *gorm.DB) {
	var area []Area
	if err := db.Find(&area).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, area)
		log.Println(area)
	}

	c.JSON(http.StatusOK, gin.H{
		"area": area,
	})
}