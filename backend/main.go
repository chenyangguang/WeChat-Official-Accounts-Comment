package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"net/http"
)

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)
	r.GET("/", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	r.Run(":8080")
}
