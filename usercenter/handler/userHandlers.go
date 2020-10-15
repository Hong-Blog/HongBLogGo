package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"usercenter/models"
	"usercenter/models/sysUser"
)

func UserList(c *gin.Context) {
	var req = sysUser.GetAllUserRequest{}
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.KeyWord = c.DefaultQuery("keyWord", "")

	users, count := sysUser.GetAllUser(req)

	var res models.PagedResponse
	res.Total = count
	res.Data = users

	c.JSON(http.StatusOK, res)
}
