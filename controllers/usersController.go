package controllers

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/jcordoba95/lp-server/initializers"
	"github.com/jcordoba95/lp-server/models"
)

func CurrentUser(c *gin.Context) models.User {
	claims := jwt.ExtractClaims(c)
	username := claims["id"]

	var user models.User
	initializers.DB.Where("username = ?", username).First(&user)
	return user
}

func GetCurrentUser(c *gin.Context) {
	user := CurrentUser(c)
	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersCreate(c *gin.Context) {
	var body struct {
		Username string
		Password string
	}
	c.Bind(&body)

	user := models.User{Username: body.Username, Password: body.Password, Status: "active"}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersIndex(c *gin.Context) {
	var users []models.User
	initializers.DB.Find(&users)

	c.JSON(200, gin.H{
		"users": users,
	})
}

func UsersShow(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	initializers.DB.Find(&user, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Username string
		Password string
	}
	c.Bind(&body)

	var user models.User
	initializers.DB.Find(&user, id)

	initializers.DB.Model(&user).Updates(models.User{
		Username: body.Username,
		Password: body.Password,
	})

	c.JSON(200, gin.H{
		"user": user,
	})
}

func UsersDelete(c *gin.Context) {
	id := c.Param("id")

	// GORM internally changes "deleted_at" from NULL to timestamp instead of deleting a record. Other requests ignore rows where "deleted_at" is not NULL.
	// To follow instructions and use "status", here we update the user's status field to inactive.
	var user models.User
	initializers.DB.Find(&user, id)
	initializers.DB.Model(&user).Updates(models.User{
		Status: "inactive",
	})

	// Normal delete operation to change "deleted_at"
	initializers.DB.Delete(&models.User{}, id)

	c.JSON(200, gin.H{
		"user": user,
	})
}
