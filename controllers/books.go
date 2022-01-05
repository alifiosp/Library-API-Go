package controllers

import (
	"net/http"

	"github.com/alifiosp/Library-API-Go/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type BooksInput struct {
	Code            string `json:"Code" binding:"required,number,gte=3"`
	Title           string `json:"Title" binding:"required"`
	Author          string `json:"Author" binding:"required"`
	Genre           string `json:"Genre" binding:"required"`
	PublicationYear string `json:"Publication Year" binding:"required,number,gte=4"`
	Placement       string `json:"Placement" binding:"required"`
}

//view data[GET]
func Show(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var bks []models.Books
	db.Find(&bks)
	c.JSON(200, gin.H{
		"data": bks,
	})
}

//add Data[POST]
func AddBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//validasi inputan/masukan
	var dataInput BooksInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message":   "Problem with Input: MUST FILL ALL LIST !",
			"Problem 1": "Code: must use number and must fill 3 number !, ex: 001",
			"Problem 2": "Publication Year: must use number and format year !, ex: 2004",
		})
		return
	}

	//input data
	bks := models.Books{
		Code:            dataInput.Code,
		Title:           dataInput.Title,
		Author:          dataInput.Author,
		Genre:           dataInput.Genre,
		PublicationYear: dataInput.PublicationYear,
		Placement:       dataInput.Placement,
	}

	//create data
	db.Create(&bks)

	//show result
	c.JSON(200, gin.H{
		"Books": bks,
	})
}

//change data[PUT]
func ChangeBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//checking data
	var bks models.Books
	if err := db.Where("Code = ?", c.Param("Code")).First(&bks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No Books/Problem with Input",
		})
		return
	}

	//validation
	var dataInput BooksInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "No Books/Problem with Input",
		})
		return
	}

	//change data
	db.Model(&bks).Update(dataInput)

	//show result
	c.JSON(200, gin.H{
		"data": bks,
	})
}

//delete data[DELETE]
func DeleteBooks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	//checking
	var bks models.Books
	if err := db.Where("Code = ?", c.Param("Code")).First(&bks).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No Books/Problem with Input",
		})
		return
	}

	//delete
	db.Delete(bks)

	//show result
	c.JSON(200, gin.H{
		"data": "Successfully Deleted",
	})
}
