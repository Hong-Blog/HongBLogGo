package postHandler

import (
	"github.com/gin-gonic/gin"
	"hong-blog/models"
	"hong-blog/models/bizArticle"
	"hong-blog/validator"
	"net/http"
	"strconv"
)

func GetArticleList(c *gin.Context) {
	var req bizArticle.GetArticleListRequest

	pageIndex, _ := strconv.Atoi(c.DefaultQuery("pageIndex", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	typeId, _ := strconv.Atoi(c.DefaultQuery("typeId", "0"))
	keyword := c.DefaultQuery("keyword", "")

	req.PageIndex = pageIndex
	req.PageSize = pageSize
	req.TypeId = typeId
	req.Keyword = keyword

	list, count := bizArticle.GetArticleList(req)

	var response models.PagedResponse
	response.Data = list
	response.Total = count
	c.JSON(http.StatusOK, response)
}

func GetDetailById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	detail, err := bizArticle.GetDetailById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, detail)
}

func DeleteById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	article := bizArticle.BizArticle{Id: id}
	err := article.Delete()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.String(http.StatusOK, "")
}

func AddArticle(c *gin.Context) {
	var req bizArticle.AddArticleModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: validator.Translate(err)})
		return
	}

	if err := bizArticle.AddArticle(req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.String(http.StatusOK, "")
}

func UpdateById(c *gin.Context) {
	var req bizArticle.EditArticleModel
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: validator.Translate(err)})
		return
	}

	id, _ := strconv.Atoi(c.Param("id"))
	req.Id = id

	if err := bizArticle.EditArticle(req); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	c.String(http.StatusOK, "")
}
