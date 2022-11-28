package controllers

import (
	"github.com/diego-arend/boilerplate_go_gin_gorm/database"
	"github.com/diego-arend/boilerplate_go_gin_gorm/models"
	"github.com/diego-arend/boilerplate_go_gin_gorm/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user models.User

	db := database.GetDatabase()

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	user.Password = services.SHA256Encoder(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create car: " + err.Error(),
		})

		return
	}

	c.Status(204)
}
