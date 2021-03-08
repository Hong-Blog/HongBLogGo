package tagHandler

import (
	"github.com/gin-gonic/gin"
	"hong-blog/models"
	"hong-blog/models/bizTags"
	"hong-blog/validator"
	"net/http"
	"strconv"
)

func TagList(c *gin.Context) {
	var req models.PagedRequest
	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	req.PageIndex = pageIndex
	req.PageSize = pageSize
	list, count := bizTags.GetTagList(req)

	var response models.PagedResponse
	response.Data = list
	response.Total = count

	c.JSON(http.StatusOK, response)
}

func AddBizTag(c *gin.Context) {
	var req bizTags.AddBizTagRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: validator.Translate(err)})
		return
	}

	if err := bizTags.AddBizTag(req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.String(http.StatusOK, "")
}

func GetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tag, err := bizTags.GetById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}

func UpdateById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var req bizTags.BizTags
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: validator.Translate(err)})
		return
	}

	req.Id = id
	if err := req.Update(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.String(http.StatusOK, "")
}

func DeleteById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	req := bizTags.BizTags{Id: id}
	if err := req.Delete(); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.String(http.StatusOK, "")
}
