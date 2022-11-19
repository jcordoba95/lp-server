package models

import (
	"errors"
	"math"
	"strconv"

	"gorm.io/gorm"
)

/*
TODOs:
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

func typeIsValid(Type string) bool {
	switch Type {
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

func DoOperation(Type string, a int, b int) (string, error) {
	switch Type {
	case "addition":
		return strconv.Itoa(a + b), nil
	case "subtraction":
		return strconv.Itoa(a - b), nil
	case "multiplication":
		return strconv.Itoa(a * b), nil
	case "division":
		if b == 0 {
			return "", errors.New("Invalid operation: Cannot divide by zero")
		}
		return strconv.Itoa(a / b), nil
	case "square_root":
		if a < 0 {
			return "", errors.New("Invalid operation: Square root of negative number")
		}
		floatNumber := float64(a)
		res := math.Sqrt(floatNumber)
		stringRes := strconv.FormatFloat(res, 'f', -1, 64)
		return stringRes, nil
	case "random_string":
		// TODO: method to generate random string from API here
		return "abc123", nil // mock for now.
	default:
		return "", errors.New("Invalid Operation Type")
	}
}
