package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"loginsrv/jwt_middleware"
	"net/http"
)

func main() {
	r := gin.Default()

	authMiddleware, err := jwt_middleware.JwtMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt_middleware.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
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

	groupAuth := r.Group("/auth")
	{
		groupAuth.GET("refresh_token", authMiddleware.RefreshHandler)
		groupAuth.Use(authMiddleware.MiddlewareFunc())
		{
			groupAuth.GET("hello", func(c *gin.Context) {
				claims := jwt.ExtractClaims(c)
				user, _ := c.Get("id")

				c.JSON(200, gin.H{
					"userID":   claims["id"],
					"userName": user.(*jwt_middleware.User).UserName,
					"text":     "Hello World.",
				})
			})
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.Run(":8090")
}
