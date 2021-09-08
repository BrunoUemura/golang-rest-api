package controllers

import (
	"strconv"

	"github.com/BrunoUemura/golang-rest-api/src/database"
	"github.com/BrunoUemura/golang-rest-api/src/models"
	"github.com/gin-gonic/gin"
)

func ShowOrders(c *gin.Context) {
	db := database.GetDatabase()

	var orders []models.Order

	err := db.Find(&orders).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list orders: " + err.Error(),
		})
		return
	}

	c.JSON(200, orders)
}

func ShowOrder(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var order models.Order
	err = db.First(&order, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find order: " + err.Error(),
		})
		return
	}

	c.JSON(200, order)
}

func CreateOrder(c *gin.Context) {
	db := database.GetDatabase()

	var order models.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Create(&order).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create order: " + err.Error(),
		})
		return
	}

	c.JSON(200, order)
}

func UpdateOrder(c *gin.Context) {
	db := database.GetDatabase()

	var order models.Order

	err := c.ShouldBindJSON(&order)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&order).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update order: " + err.Error(),
		})
		return
	}

	c.JSON(200, order)
}

func DeleteOrder(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var order models.Order
	err = db.Delete(&order, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete order: " + err.Error(),
		})
		return
	}

	c.Status(204)
}