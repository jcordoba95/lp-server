package controllers

import (
	"github.com/gin-gonic/gin"
)

func RecordsIndex(c *gin.Context) {
	// Return paginated index: https://articles.wesionary.team/implement-pagination-in-golang-using-gorm-and-gin-b4ad8e2932a6
}

func RecordsCreate(c *gin.Context) {
	var body struct {
		OperationID int
	}
	c.Bind(&body)

	// Extract current user from jwt token in request
	// Get last user balance
	// Get amount to be subtracted from user balance
	// Validate user can "pay" for operation
	// Perform operation
	// Save record
}

func RecordsDelete(c *gin.Context) {
	// id := c.Param("id")

	// Begin transaction
	// Create new record with last_user_balance + inverse of the amount of the record to be deleted <- this is to adjust balance
	// Delete record
	// Commit transaction

	c.Status(200)
}
