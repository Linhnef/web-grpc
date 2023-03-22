package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"web/gin-gonic/database"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

var db = database.Connect()

func Create(c *gin.Context) {

	var user User
	c.ShouldBind(&user)

	if err := db.Table("users").Create(&user).Error; err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
	})
}

func Get(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user User
	if err := db.Where("id = ?", id).Table("users").First(&user).Error; err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}

func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	var user User

	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.ShouldBind(&user)
	db.Save(&user)
	c.JSON(200, gin.H{
		"data": user,
	})

}

func Delete(c *gin.Context) {

	id, err := strconv.Atoi(c.Params.ByName("id"))

	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	}
	var user User
	if err := db.Where("id = ?", id).Table("users").Delete(&user).Error; err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(200, gin.H{
			"data": user,
		})
	}
}
