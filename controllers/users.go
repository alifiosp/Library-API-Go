package controllers

import (
	"net/http"

	"github.com/alifiosp/Library-API-Go/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type UsersInput struct {
	Id      string `json:"Id" binding:"required,number,gte=5"`
	Name    string `json:"Name" binding:"required"`
	Email   string `json:"Email" binding:"required,email"`
	Address string `json:"Address" binding:"required"`
	Status  string `json:"Status" binding:"required"`
}

//view data[GET]
func ShowUsers(c *gin.Context) {
	db1 := c.MustGet("db1").(*gorm.DB)

	var usr []models.Users
	db1.Find(&usr)
	c.JSON(200, gin.H{
		"data": usr,
	})
}

//add Data[POST]
func AddUsers(c *gin.Context) {
	db1 := c.MustGet("db1").(*gorm.DB)

	//validasi inputan/masukan
	var dataInput UsersInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message":   "Problem with Input: MUST FILL ALL LIST !",
			"Problem 1": "Id: must use number and must fill 5 number !, ex: 01234",
			"Problem 2": "Email: must use email format !",
		})
		return
	}

	//input data
	usr := models.Users{
		Id:      dataInput.Id,
		Name:    dataInput.Name,
		Email:   dataInput.Email,
		Address: dataInput.Address,
		Status:  dataInput.Status,
	}

	//create data
	db1.Create(&usr)

	//show result
	c.JSON(200, gin.H{
		"Users": usr,
	})
}

//change data[PUT]
func ChangeUsers(c *gin.Context) {
	db1 := c.MustGet("db1").(*gorm.DB)

	//checking data
	var usr models.Users
	if err := db1.Where("Id = ?", c.Param("Id")).First(&usr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No User/Problem with Input",
		})
		return
	}

	//validation
	var dataInput UsersInput
	if err := c.ShouldBindJSON(&dataInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Message": "No User/Problem with Input",
		})
		return
	}

	//change data
	db1.Model(&usr).Update(dataInput)

	//show result
	c.JSON(200, gin.H{
		"data": usr,
	})
}

//delete data[DELETE]
func DeleteUsers(c *gin.Context) {
	db1 := c.MustGet("db1").(*gorm.DB)

	//checking
	var usr models.Users
	if err := db1.Where("Id = ?", c.Param("Id")).First(&usr).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "No User/Problem with Input",
		})
		return
	}

	//delete
	db1.Delete(usr)

	//show result
	c.JSON(200, gin.H{
		"data": "Successfully Deleted",
	})
}
