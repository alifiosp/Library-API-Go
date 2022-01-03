package main

import (
	"net/http"

	"github.com/alifiosp/Library-API-Go/controllers"
	"github.com/alifiosp/Library-API-Go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Data": "Home Page",
		})
	})

	v1 := r.Group("/Library")
	v2 := r.Group("/LibraryUsers")

	db := models.SetupModelsBooks()

	v1.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	v1.GET("/Books", controllers.Show)
	v1.POST("/Books", controllers.AddBooks)
	v1.PUT("/Books/:Code", controllers.ChangeBooks)
	v1.DELETE("/Books/:Code", controllers.DeleteBooks)

	db1 := models.SetupModelsUsers()

	v2.Use(func(c *gin.Context) {
		c.Set("db1", db1)
		c.Next()
	})

	v2.GET("/Users", controllers.ShowUsers)
	v2.POST("/Users", controllers.AddUsers)
	v2.PUT("/Users/:Id", controllers.ChangeUsers)
	v2.DELETE("/Users/:Id", controllers.DeleteUsers)

	r.Run()
}
