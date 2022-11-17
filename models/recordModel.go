package models

import (
	"time"

	"gorm.io/gorm"
)

/*
TODOs:
- Add private method to check if associated user can perform operation before saving. Return error if not
*/

type Record struct {
	gorm.Model
	OperationID       int
	Operation         Operation
	UserID            int `gorm:"index:idx_balance"`
	User              User
	Amount            int64
	UserBalance       int64 `gorm:"index:idx_balance,priority:1,sort:desc"`
	OperationResponse string
	Date              time.Time
}
