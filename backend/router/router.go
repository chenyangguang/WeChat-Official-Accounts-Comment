package router

import (
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/controller"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
)

func InitRouter() *gin.Engine {

	logFile, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter)

	r := gin.Default()

	r.GET("/comments", controller.GetAllComments)
	r.GET("/comment/:id", controller.GetCommentById)
	r.POST("/comment", controller.CreateComment)
	r.PUT("/comment/:id", controller.UpdateComment)
	r.DELETE("/comment/:id", controller.DeleteComment)

	return r

}
