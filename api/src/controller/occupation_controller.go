package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"gorm.io/gorm"
)

type Occupation struct {
	ID int `json:”id”`
	Name string `json:”name”`
}

func GetOccupation(c *gin.Context, db *gorm.DB) {
	var occupation []Occupation
	if err := db.Find(&occupation).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, occupation)
		log.Println(occupation)
	}

	c.JSON(http.StatusOK, gin.H{
		"occupation": occupation,
	})
}