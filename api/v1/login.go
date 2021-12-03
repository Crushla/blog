package v1

import (
	"blog/middleware"
	"blog/models"
	"blog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var user models.User
	var token string
	c.ShouldBindJSON(&user)
	code := models.CheckLogin(user.Username, user.Password)
	if code == errmsg.SUCCESS {
		token, code = middleware.GenerateToken(user.Username)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"token":   token,
	})
}
