package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	"time"
	"api/src/controller"
)

type HandlerWithDB struct {
    DB *gorm.DB
}

func (h *HandlerWithDB) GetUserWrapper(c *gin.Context) {
    controller.GetUser(c, h.DB)
}

func (h *HandlerWithDB) GetAreaWrapper(c *gin.Context) {
    controller.GetArea(c, h.DB)
}

func (h *HandlerWithDB) GetOccupationWrapper(c *gin.Context) {
    controller.GetOccupation(c, h.DB)
}

func main() {
	r := gin.Default()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"greet": "hello, world!",
		})
	})

	r.GET("/users", func(c *gin.Context) {
        handler := &HandlerWithDB{DB: InitDB()}
        handler.GetUserWrapper(c)
    })

	r.GET("/areas", func(c *gin.Context) {
        handler := &HandlerWithDB{DB: InitDB()}
        handler.GetAreaWrapper(c)
    })

	r.GET("/occupations", func(c *gin.Context) {
		handler := &HandlerWithDB{DB: InitDB()}
		handler.GetOccupationWrapper(c)
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}

func InitDB() *gorm.DB {
	dsn := "host=172.28.1.3 user=postgres password=postgres dbname=pickup_analytics port=5432 sslmode=disable TimeZone=Asia/Tokyo"

	var db *gorm.DB
	var err error

	retries := 3
	for retries > 0 {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break // 接続成功したらループを抜ける
		}

		log.Printf("failed to connect database: %v\n", err)
		log.Println("retrying connection...")

		retries--
		time.Sleep(5 * time.Second) // リトライの間隔を設定する場合、ここで調整する
	}

	if err != nil {
		log.Fatalf("failed to connect database after retries: %v", err)
	}

	return db
}
