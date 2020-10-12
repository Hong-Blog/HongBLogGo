package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usercenter/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", index)
	router.GET("ping", ping)

	userRouter := router.Group(`user`)
	{
		userRouter.GET(`/`, handler.UserList)
		userRouter.GET(`/:id`, handler.UserList)
	}

	return router
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func index(c *gin.Context) {
	c.String(http.StatusOK, "hello gin")
}
