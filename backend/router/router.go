package router

import (
	"github.com/gin-gonic/gin"

	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/controller"
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/middleware"
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/load/log"
)

func InitRouter() *gin.Engine {
	log.InitLog()

	r := gin.Default()
	r.Use(middleware.LogMiddleware())
	r.Use(middleware.RequestIdMiddleware())
	v1 := r.Group("/v1")
	v1.Use(middleware.AuthMiddleware())
	{
		v1.GET("/comments", controller.GetAllComments)
		v1.GET("/comment/:id", controller.GetCommentById)
		v1.POST("/comment", controller.CreateComment)
		v1.PUT("/comment/:id", controller.UpdateComment)
		v1.DELETE("/comment/:id", controller.DeleteComment)
	}

	return r
}
