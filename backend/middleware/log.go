package middleware

import (
	"fmt"
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/config"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

// LogMiddleware 日志中间件
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		Log2File()
	}
}

// Log2File 日志写文件
func Log2File() gin.HandlerFunc {
	fileName := path.Join(config.LogFilePath, config.LogFileName)
	out, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("err", err)
	}

	logger := logrus.New()
	logger.Out = out
	logger.SetLevel(logrus.DebugLevel)

	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}
}

// todo Log2ES
