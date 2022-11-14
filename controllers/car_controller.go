package controllers

import (
	"strconv"

	"github.com/diego-arend/boilerplate_go_gin_gorm/database"
	"github.com/diego-arend/boilerplate_go_gin_gorm/models"
	"github.com/gin-gonic/gin"
)

func ShowCar(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	var car models.Car
	err = db.First(&car, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find car id: " + err.Error(),
		})

		return
	}

	c.JSON(200, car)
}

func ListCars(c *gin.Context) {
	var cars []models.Car

	db := database.GetDatabase()

	err := db.Find(&cars).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list books: " + err.Error(),
		})

		return
	}

	c.JSON(200, cars)
}

func CreateCar(c *gin.Context) {
	var car models.Car

	db := database.GetDatabase()

	err := c.ShouldBindJSON(&car)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Create(&car).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create car: " + err.Error(),
		})

		return
	}

	c.JSON(200, car)
}

func UpdateCar(c *gin.Context) {
	var car models.Car

	db := database.GetDatabase()

	err := c.ShouldBindJSON(&car)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})

		return
	}

	err = db.Save(&car).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update car: " + err.Error(),
		})

		return
	}

	c.JSON(200, car)
}

func DeleteCar(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})

		return
	}

	db := database.GetDatabase()

	err = db.Delete(&models.Car{}, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete car id: " + err.Error(),
		})

		return
	}

	c.Status(204)
}
