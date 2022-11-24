package models

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeIsValid(t *testing.T) {
	validTypes := []string{"addition", "subtraction", "multiplication", "division", "square_root", "random_string"}
	invalidTypes := []string{"module", "power_of", "factorial", "imaginary", "inverse", "sum"}
	for _, val := range validTypes {
		assert.Equal(t, typeIsValid(val), true, "should return true")
	}
	for _, val := range invalidTypes {
		assert.Equal(t, typeIsValid(val), false, "should return false")
	}
}

func TestDoOperationAddition(t *testing.T) {
	val1, _ := DoOperation("addition", 5, 10)
	val2, _ := DoOperation("addition", 1927465, 282749572)
	assert.Equal(t, val1, "15", "should add numbers properly")
	assert.Equal(t, val2, "284677037", "should add numbers properly")
}

func TestDoOperationSubtraction(t *testing.T) {
	val1, _ := DoOperation("subtraction", 5, 10)
	val2, _ := DoOperation("subtraction", 1927465, 282749572)
	assert.Equal(t, val1, "-5", "should subtract numbers properly")
	assert.Equal(t, val2, "-280822107", "should subtract numbers properly")
}

func TestDoOperationMultiplication(t *testing.T) {
	val1, _ := DoOperation("multiplication", 5, 10)
	val2, _ := DoOperation("multiplication", 300, 123)
	assert.Equal(t, val1, "50", "should multiply numbers properly")
	assert.Equal(t, val2, "36900", "should multiply numbers properly")
}

func TestDoOperationDivision(t *testing.T) {
	val1, _ := DoOperation("division", 10, 2)
	val2, _ := DoOperation("division", 300, 100)
	assert.Equal(t, val1, "5", "should divide numbers properly")
	assert.Equal(t, val2, "3", "should divide numbers properly")

	_, err := DoOperation("division", 100, 0)
	assert.ErrorContains(t, err, "Invalid operation: Cannot divide by zero")
}

func TestDoOperationSquareRoot(t *testing.T) {
	val1, _ := DoOperation("square_root", 144, 0)
	val2, _ := DoOperation("square_root", 225, 0)
	assert.Equal(t, val1, "12", "should divide numbers properly")
	assert.Equal(t, val2, "15", "should divide numbers properly")

	_, err := DoOperation("square_root", -1, 0)
	assert.ErrorContains(t, err, "Invalid operation: Square root of negative number")
}

func TestDoOperationRandomString(t *testing.T) {
	val, _ := DoOperation("random_string", 0, 0)
	assert.Equal(t, reflect.TypeOf(val).Kind(), reflect.TypeOf("").Kind(), "should be a string")
	assert.Equal(t, len(val), 10, "should be length 10")
}
