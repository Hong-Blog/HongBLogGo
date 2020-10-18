package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usercenter/handler/roleHandler"
	"usercenter/handler/userHandler"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/", index)
	router.GET("ping", ping)

	userRouter := router.Group(`user`)
	{
		userRouter.GET(`/`, userHandler.UserList)
		userRouter.GET(":id", userHandler.GetById)
	}

	roleRouter := router.Group("role")
	{
		roleRouter.GET("/", roleHandler.RoleList)
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
