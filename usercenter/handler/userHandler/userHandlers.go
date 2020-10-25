package userHandler

import (
	"github.com/gin-gonic/gin"
	"log"
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

func GetById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Panicln("get user by id no found id", err.Error())
	}
	user := sysUser.GetById(id)
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	req := sysUser.UpdateUserRequest{}
	req.Id, _ = strconv.Atoi(c.Param("id"))

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}

	success := sysUser.UpdateUser(req)
	if !success {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "更新失败"})
		return
	}
	c.String(http.StatusOK, "")
}

func AddUser(c *gin.Context) {
	req := sysUser.AddUserRequest{}
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	req.RegIp = c.ClientIP()
	success, err := sysUser.AddUser(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	if !success {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "添加用户失败"})
		return
	}
	c.String(http.StatusOK, "")
}

func DeleteUser(c *gin.Context) {
	user := sysUser.SysUser{}
	user.Id, _ = strconv.Atoi(c.Param("id"))

	if err := user.LogicalDeleteById(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.String(http.StatusOK, "")
}

func UpdatePassword(c *gin.Context) {
	request := sysUser.UpdatePasswordByIdRequest{}
	request.Id, _ = strconv.Atoi(c.Param("id"))
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: err.Error()})
		return
	}
	if err := sysUser.UpdatePasswordById(request); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}
	c.String(http.StatusOK, "")
}
