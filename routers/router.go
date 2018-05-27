package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaoheigou/GoFun/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	apiv1 := r.Group("/api/v1")
	{

		// //获取文章列表
		// apiv1.GET("/articles", v1.GetArticles)
		// //获取指定文章
		// apiv1.GET("/articles/:id", v1.GetArticle)
		// //新建文章
		// apiv1.POST("/articles", v1.AddArticle)
		// //更新指定文章
		// apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		apiv1.GET("/test", v1.Get)
	}

	return r
}