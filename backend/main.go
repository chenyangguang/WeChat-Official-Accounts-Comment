package main

import (
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/router"
	"os"
	"github.com/gin-gonic/gin"
	"io"
	"log"
)

func initLogConfig () {
	logFile, err := os.Create("gin.log")
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	log.SetOutput(gin.DefaultWriter)
}

func main() {
	initLogConfig()
	r := router.InitRouter()
	r.Run()
}
