package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserList(c *gin.Context) {
	var res struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
	res.Message = `ok`
	res.Code = http.StatusOK
	c.JSON(http.StatusOK, res)
}
