package controllers

// import (
// 	"io/ioutil"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/assert"
// )

// func SetUpRouter() *gin.Engine {
// 	router := gin.Default()
// 	return router
// }

// func TestRecordsIndex(t *testing.T) {
// 	mockResponse := `{"records":"example"}`
// 	r := SetUpRouter()
// 	r.GET("/records", RecordsIndex)
// 	req, _ := http.NewRequest("GET", "/records", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	responseData, _ := ioutil.ReadAll(w.Body)
// 	assert.Equal(t, mockResponse, string(responseData))
// 	assert.Equal(t, http.StatusOK, w.Code)
// }
