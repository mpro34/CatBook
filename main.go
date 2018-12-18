package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Globals for db access
var db *gorm.DB
var err error

type Cat struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Breed string `json:"breed"`
}

func main() {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close() //When db is finished with, close connection (asynchronous)

	db.AutoMigrate(&Cat{}) // Migrate the schema

	c1 := Cat{Name: "Zoey", Breed: "Tabby"}
	db.Create(&c1) // Create object and store in db

	// Create a gin router
	r := gin.Default()

	// Endpoints
	r.GET("/cats", GetCats)
	r.GET("/cats/:id", GetCat)
	r.POST("/cats", CreateCat)
	r.PUT("/cats/:id", UpdateCat)
	r.DELETE("/cats/:id", DeleteCat)

	r.Run(":8080")
}

//List
func GetCats(c *gin.Context) {
	var cats []Cat
	if err := db.Find(&cats).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cats)
	}
}

// Retrieve
func GetCat(c *gin.Context) {
	id := c.Params.ByName("id")
	var cat Cat
	if err := db.Where("id = ?", id).First(&cat).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, cat)
	}
}

// Create
func CreateCat(c *gin.Context) {
	var cat Cat
	c.BindJSON(&cat) //TODO: What does BindJSON() do?

	db.Create(&cat)
	c.JSON(200, cat)
}

// Update
func UpdateCat(c *gin.Context) {
	var cat Cat
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&cat).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&cat)

	db.Save(&cat)
	c.JSON(200, cat)
}

// Delete
func DeleteCat(c *gin.Context) {
	var cat Cat
	id := c.Params.ByName("id")
	d := db.Where("id = ?", id).Delete(&cat)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}