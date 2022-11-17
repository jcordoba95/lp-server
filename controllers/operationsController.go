package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jcordoba95/lp-server/initializers"
	"github.com/jcordoba95/lp-server/models"
)

func OperationsCreate(c *gin.Context) {
	var body struct {
		Type string
		Cost int64
	}
	c.Bind(&body)

	operation := models.Operation{Type: body.Type, Cost: body.Cost}
	result := initializers.DB.Create(&operation)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"operation": operation,
	})
}

func OperationsIndex(c *gin.Context) {
	var operations []models.Operation
	initializers.DB.Find(&operations)

	c.JSON(200, gin.H{
		"operations": operations,
	})
}

func OperationsShow(c *gin.Context) {
	id := c.Param("id")

	var operation models.Operation
	initializers.DB.Find(&operation, id)

	c.JSON(200, gin.H{
		"operation": operation,
	})
}

func OperationsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Type string
		Cost int64
	}
	c.Bind(&body)

	var operation models.Operation
	initializers.DB.Find(&operation, id)

	initializers.DB.Model(&operation).Updates(models.Operation{
		Type: body.Type,
		Cost: body.Cost,
	})

	c.JSON(200, gin.H{
		"operation": operation,
	})
}

func OperationsDelete(c *gin.Context) {
	id := c.Param("id")

	initializers.DB.Delete(&models.Operation{}, id)

	c.Status(200)
}
