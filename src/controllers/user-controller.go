package controllers

import (
	"strconv"

	"github.com/BrunoUemura/golang-rest-api/src/database"
	"github.com/BrunoUemura/golang-rest-api/src/models"
	"github.com/BrunoUemura/golang-rest-api/src/services"
	"github.com/gin-gonic/gin"
)

func ShowUsers(c *gin.Context) {
	db := database.GetDatabase()

	var users []models.User

	err := db.Find(&users).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot list users: " + err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func ShowUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var user models.User
	err = db.First(&user, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot find user: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func CreateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	user.Password, _ = services.HashPassword(user.Password)

	err = db.Create(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot create user: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func UpdateUser(c *gin.Context) {
	db := database.GetDatabase()

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot bind JSON: " + err.Error(),
		})
		return
	}

	err = db.Save(&user).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot update user: " + err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	newid, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "ID has to be integer",
		})
		return
	}

	db := database.GetDatabase()

	var user models.User
	err = db.Delete(&user, newid).Error
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Cannot delete user: " + err.Error(),
		})
		return
	}

	c.Status(204)
}