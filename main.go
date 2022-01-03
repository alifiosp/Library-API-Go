package main

import (
	"net/http"

	"github.com/alifiosp/Library-API-Go/controllers"
	"github.com/alifiosp/Library-API-Go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := models.SetupModelsbooks()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"Data": "Home Page",
		})
	})

	r.GET("/Books", controllers.Show)
	r.POST("/Books", controllers.AddBooks)
	r.PUT("/Books/:Code", controllers.ChangeBooks)
	r.DELETE("/Books/:Code", controllers.DeleteBooks)

	r.Run()
}
