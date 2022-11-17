package models

import (
	"errors"

	"gorm.io/gorm"
)

/*
TODOs:
- Add email validation for username
- Add password encryption
- Add user auth for api with jwt
- Create method to return current user balance
*/

type User struct {
	gorm.Model
	Username string
	Password string
	Status   string
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if !statusIsValid(u.Status) {
		err = errors.New("Invalid status")
	}
	return
}

func statusIsValid(status string) bool {
	switch status {
	case
		"active",
		"inactive":
		return true
	}
	return false
}
