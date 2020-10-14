package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"usercenter/models/sysUser"
)

func UserList(c *gin.Context) {
	users := sysUser.GetAllUser()
	c.JSON(http.StatusOK, users)
}
