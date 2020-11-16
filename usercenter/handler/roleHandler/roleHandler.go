package roleHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"usercenter/models"
	"usercenter/models/sysRole"
)

// 角色列表
// @Summary 角色列表
// @Description 角色列表
// @Tags 角色
// @Success 200 {array} sysRole.SysRole
// @Router /roles [get]
func RoleList(c *gin.Context) {
	var req = sysRole.GetAllRoleRequest{}
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.KeyWord = c.DefaultQuery("keyWord", "")

	list, count := sysRole.GetAllRole(req)
	var res models.PagedResponse
	res.Total = count
	res.Data = list

	c.JSON(http.StatusOK, res)
}
