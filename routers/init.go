package routers

import (
	v1 "blog/api/v1"
	"blog/middleware"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())
	Auth := r.Group("api/v1")
	Auth.Use(middleware.JwtToken())
	{
		//user模块路由接口
		Auth.POST("user/add", v1.AddUser)
		Auth.PUT("user/:id", v1.EditUser)
		Auth.DELETE("user/:id", v1.DeleteUser)
		//分类模块路由接口
		Auth.POST("category/add", v1.AddCategory)
		Auth.PUT("category/:id", v1.EditCategory)
		Auth.DELETE("category/:id", v1.DeleteCategory)
		//文章模块路由接口
		Auth.POST("article/add", v1.AddArticle)
		Auth.PUT("article/:id", v1.EditArticle)
		Auth.DELETE("article/:id", v1.DeleteArticle)
		//上传文件
		Auth.POST("upload", v1.Upload)
	}
	router := r.Group("api/v1")
	{
		router.GET("users", v1.GetUsers)
		router.GET("category", v1.GetCategory)
		router.GET("article", v1.GetArticle)
		router.GET("article/info/:id", v1.GetArtInfo)
		router.GET("article/list/:id", v1.GetCateArt)
		router.POST("login", v1.Login)
	}
	return r
}
