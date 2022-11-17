package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jcordoba95/lp-server/controllers"
	"github.com/jcordoba95/lp-server/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// User Routes
	r.POST("/users", controllers.UsersCreate)
	r.PUT("/users/:id", controllers.UsersUpdate)
	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.UsersShow)
	r.DELETE("/users/:id", controllers.UsersDelete)

	// Operation Routes
	r.POST("/operations", controllers.OperationsCreate)
	r.PUT("/operations/:id", controllers.OperationsUpdate)
	r.GET("/operations", controllers.OperationsIndex)
	r.GET("/operations/:id", controllers.OperationsShow)
	r.DELETE("/operations/:id", controllers.OperationsDelete)

	// Record Routes

	r.Run() // listen and serve on 0.0.0.0:8080
}
