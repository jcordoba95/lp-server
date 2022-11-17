package models

import (
	"errors"

	"gorm.io/gorm"
)

/*
TODOs:
- Change type from string to enum
- Create util folder or class to request random string generation and use that client inside a method to generate random string
*/

type Operation struct {
	gorm.Model
	Type string `gorm:"uniqueIndex"`
	Cost int64
}

func (u *Operation) BeforeSave(tx *gorm.DB) (err error) {
	if !typeIsValid(u.Type) {
		err = errors.New("Invalid type")
	}
	return
}

func typeIsValid(status string) bool {
	switch status {
	case
		"addition",
		"subtraction",
		"multiplication",
		"division",
		"square_root",
		"random_string":
		return true
	}
	return false
}
