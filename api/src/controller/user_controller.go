package controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"gorm.io/gorm"
	"time"
)

type User struct {
    ID              int       `db:"id"`
    Name            string    `db:"name"`
    Mail            string    `db:"mail"`
    Pass            string    `db:"pass"`
    Age             int       `db:"age"`
    ExperienceYears int       `db:"experience_years"`
    IsPublished     bool      `db:"is_published"`
    CreatedAt       time.Time `db:"created_at"`
    UpdatedAt       time.Time `db:"updated_at"`
}

func GetUser(c *gin.Context, db *gorm.DB) {
	var user []User
	if err := db.Find(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
		log.Println(user)
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}