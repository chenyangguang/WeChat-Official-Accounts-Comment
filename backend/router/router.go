package router

import (
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/middleware"
	"github.com/gin-gonic/gin"

	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/controller"
	"io"
	"log"
	"os"
)

func initLogfile() {
	logFile, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter)

}
func InitRouter() *gin.Engine {
	initLogfile()

	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/comments", controller.GetAllComments)
		v1.GET("/comment/:id", controller.GetCommentById)
		v1.POST("/comment", controller.CreateComment)
		v1.PUT("/comment/:id", controller.UpdateComment)
		v1.DELETE("/comment/:id", controller.DeleteComment)
	}
	v1.Use(middleware.AuthMiddleware(),middleware.RequestIdMiddleware())

	return r

}
