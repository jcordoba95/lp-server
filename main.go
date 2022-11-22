package main

import (
	"fmt"
	"log"
	"os"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jcordoba95/lp-server/controllers"
	"github.com/jcordoba95/lp-server/initializers"
	"github.com/jcordoba95/lp-server/models"
)

type login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var identityKey = "id"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)
	config := cors.DefaultConfig()
	config.AddAllowHeaders("Content-Type,access-control-allow-origin, access-control-allow-headers, Authorization")
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       fmt.Sprintf("lp-jcordoba95-%s", os.Getenv("ENVIRONMENT")),
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &models.User{
				Username: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			var user models.User
			username := loginVals.Username
			password := loginVals.Password

			initializers.DB.Where("username = ? AND password = ?", username, password).First(&user)
			if &user != nil {
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*models.User); ok && v.Username != "" {
				return true
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.POST("/login", authMiddleware.LoginHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/v1")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/users", controllers.UsersIndex)
		auth.GET("/me", controllers.GetCurrentUser)
		auth.GET("/records", controllers.RecordsIndex)
		auth.POST("/records", controllers.RecordsCreate)
		auth.DELETE("records/:id", controllers.RecordsDelete)
	}

	r.Run(os.Getenv("PORT"))
}
