package v1

import (
	"blog/models"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context) {
	var article models.Article
	_ = c.ShouldBindJSON(&article)
	code := models.CreateArticle(&article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    article,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询文章
func GetArticle(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	articles, code, total := models.GetArticle(pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询分类下所有文章
func GetCateArt(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	id, _ := strconv.Atoi(c.Param("id"))
	articles, code, total := models.GetCateArt(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    articles,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询单个文章
func GetArtInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	artInfo, code := models.GetArtInfo(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    artInfo,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑文章
func EditArticle(c *gin.Context) {
	var article models.Article
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&article)
	code := models.EditArticle(id, &article)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := models.DeleteArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
