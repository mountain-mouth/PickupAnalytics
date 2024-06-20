package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type Creature struct {
	ID    uint   `json:”id”`
	Code  string `json:”code”`
	Name string   `json:”name”`
}

func main() {

	r := gin.Default()

	// テストデータ作成
	InsertTestData()

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"greet": "hello, world!",
		})
	})

	r.GET("/echo/:echo", func(c *gin.Context) {
		echo := c.Param("echo")
		c.JSON(http.StatusOK, gin.H{
			"echo": echo,
		})
	})

	r.POST("/upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// Upload the file to specific dst.
			// c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, gin.H{
			"uploaded": len(files),
		})
	})

	r.GET("/creature", GetCreature)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func InitDB() *gorm.DB {
	dsn := "host=172.28.1.3 user=postgres password=postgres dbname=pickup_analytics port=5432 sslmode=disable TimeZone=Asia/Tokyo"
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect database: %v", err)
    }
	return db
}

func InsertTestData() {
	db := InitDB()
	db.AutoMigrate(&Creature{})

	testCreaturetData := []Creature{
		{Code: "L1212", Name: "うみねこ"},
		{Code: "L1213", Name: "うり"},
		{Code: "L1214", Name: "ねこ"},
	}

	for _, creature := range testCreaturetData {
		result := db.Create(&creature)

		if result.Error != nil {
			fmt.Println("エラー:", result.Error)
		} else {
			fmt.Println("新しい商品のID:", creature.Name)
		}
	} 
}

func GetCreature(c *gin.Context) {
	db := InitDB()

	var creature []Creature

	if err := db.Find(&creature).Error; err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, creature)
		log.Println(creature)
	}

}
