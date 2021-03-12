package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"hong-blog/middlewares"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()

	authMiddleware, err := middlewares.JwtMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// When you use jwt_middleware.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := authMiddleware.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/authorization", authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	// todo 需要更新redis
	//r.GET("/refresh_token", authMiddleware.RefreshHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	addr := ":80"
	if gin.Mode() == gin.DebugMode {
		addr = ":8082"
	}

	_ = r.Run(addr)
}
