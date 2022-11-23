package models

import (
	"errors"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"
	"strings"

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
		response, err := http.Get("https://www.random.org/strings/?num=1&len=10&digits=on&loweralpha=on&unique=off&format=plain&rnd=new")
		if err != nil {
			return "", err
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", err
		}
		// remove new line at the end of the response and return random string
		return strings.TrimSuffix(string(responseData), "\n"), nil
	default:
		return "", errors.New("Invalid Operation Type")
	}
}
