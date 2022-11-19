package controllers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jcordoba95/lp-server/initializers"
	"github.com/jcordoba95/lp-server/models"
)

func RecordsIndex(c *gin.Context) {
	// Return paginated index: https://articles.wesionary.team/implement-pagination-in-golang-using-gorm-and-gin-b4ad8e2932a6
}

func RecordsCreate(c *gin.Context) {
	var body struct {
		OperationID int
		Val1        int
		Val2        int
	}
	c.Bind(&body)
	user := CurrentUser(c)
	var operation models.Operation
	initializers.DB.First(&operation, body.OperationID)

	// Check for user balance or assign new user balance for new user
	var userBalance int64
	var lastUserRecord models.Record
	if res := initializers.DB.Where("user_id = ?", user.ID).Limit(1).Order("created_at desc").Find(&lastUserRecord); res.Error != nil {
		userBalance = int64(1000)
	} else {
		userBalance = lastUserRecord.UserBalance
	}

	tx := initializers.DB.Begin()
	// Can the user pay for the operation?
	if userBalance < operation.Cost {
		tx.Rollback()
		c.JSON(406, gin.H{
			"error": "User balance is lower than operation's cost.",
		})
		return
	}

	// Perform operation
	operationResponse, err := models.DoOperation(operation.Type, body.Val1, body.Val2)
	if err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Save record
	amount := operation.Cost
	newUserBalance := userBalance - amount
	newRecord := models.Record{
		OperationID:       body.OperationID,
		Operation:         operation,
		User:              user,
		UserID:            int(user.ID),
		Amount:            amount,
		UserBalance:       newUserBalance,
		OperationResponse: operationResponse,
		Date:              time.Now(),
	}
	if err := tx.Create(&newRecord).Error; err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	// Commit transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"record": newRecord,
	})
}

func RecordsDelete(c *gin.Context) {
	id := c.Param("id")

	var record models.Record
	initializers.DB.First(&record, id)

	err, result := deleteRecord(record)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"record": result,
	})
}

func deleteRecord(record models.Record) (error, models.Record) {
	var userBalance int64
	var lastUserRecord models.Record
	if res := initializers.DB.Where("user_id = ?", record.UserID).Limit(1).Order("created_at desc").Find(&lastUserRecord); res.Error != nil {
		return res.Error, models.Record{}
	} else {
		userBalance = lastUserRecord.UserBalance
	}
	// Adjust Balance: Add back what was subtracted
	newUserBalance := userBalance + record.Amount
	newRecord := models.Record{
		UserID:      record.UserID,
		OperationID: record.OperationID,
		Amount:      (-1 * record.Amount),
		UserBalance: newUserBalance,
		Date:        time.Now(),
	}

	tx := initializers.DB.Begin()
	if err := tx.Create(&newRecord).Error; err != nil {
		tx.Rollback()
		return err, models.Record{}
	}
	if err := tx.Delete(&record).Error; err != nil {
		tx.Rollback()
		return err, models.Record{}
	}

	return tx.Commit().Error, newRecord
}
