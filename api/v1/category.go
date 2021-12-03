package v1

import (
	"blog/models"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

//添加分类
func AddCategory(c *gin.Context) {
	var category models.Category
	_ = c.ShouldBindJSON(&category)
	code := models.CheckCategory(category.Name)
	if code == errmsg.SUCCESS {
		models.CreateCategory(&category)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    category,
		"message": errmsg.GetErrMsg(code),
	})
}

//查询分类
func GetCategory(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	if pageSize == 0 {
		pageSize = -1
	}
	if pageNum == 0 {
		pageNum = -1
	}
	users, total := models.GetCategory(pageSize, pageNum)
	code := errmsg.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    users,
		"total":   total,
		"message": errmsg.GetErrMsg(code),
	})
}

//编辑分类
func EditCategory(c *gin.Context) {
	var category models.Category
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&category)
	code := models.CheckUser(category.Name)
	if code == errmsg.SUCCESS {
		models.EditCategory(id, &category)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

//删除分类
func DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := models.DeleteCategory(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
