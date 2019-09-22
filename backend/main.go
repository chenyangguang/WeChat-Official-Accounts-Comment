package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Comment 留言
type Comment struct {
	gorm.Model
	Content    string `json:"content" form:"content"`
	ArticleId  string `json:"article_id" form:"article_id"`
	CommentUid string `json:"comment_uid" form:"comment_uid"`
	ParentId   int    `json:"parent_id" form:"parent_id"`
	IsTop      int    `json:"is_top" form:"is_top"`
	Status     int    `json:"status" form:"status"`
}

// User 用户表
type User struct {
	Uid       int64 `grom:"primary_key"`
	Openid    string
	Nickname  string
	Gender    int
	Avatar    string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

func main() {

	db, err := gorm.Open("mysql", "git:Git0618@/wechat_comments?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("imit connect database fail:", err)
	}
	defer db.Close()

	db.AutoMigrate(&Comment{})

	r := gin.Default()

	// todo 公众号信息

	// todo 文章
	r.GET("/article", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "",
		})
	})

	// curl localhost:8080/comments?article_id=article_1
	// 查询一篇文章中的留言列表
	r.GET("/comments", func(c *gin.Context) {
		articleId := c.Query("article_id")
		fmt.Println(articleId)

		var comments = []Comment{}
		if err := db.Where("article_id = ? ", articleId).Find(&comments).Error; err != nil {
			fmt.Println(err)
			c.JSON(1002, gin.H{
				"message": "Query comment fail",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message":       "Success",
			"comment_lists": comments,
		})
	})

	// 单挑留言
	// curl localhost:8080/comment/4 -X GET
	r.GET("/comment/:id", func(c *gin.Context) {
		id := c.Param("id")
		var comment Comment
		if result := db.Where("id = ? ", id).First(&comment); result.Error != nil || !result.RecordNotFound() {
			c.AbortWithStatus(404)
		}

		c.JSON(http.StatusOK, gin.H{
			"meesage": "Get record success.",
			"comment": comment,
		})
	})

	// curl localhost:8080/comment -X POST -d "content=3content&article_id=article_3&comment_uid=uid_3"
	// 新增留言
	r.POST("/comment", func(c *gin.Context) {
		commentUid := c.PostForm("comment_uid")
		articleId := c.PostForm("article_id")
		content := c.PostForm("content")
		parentId := c.GetInt("parent_id")

		// do save comment
		comment := &Comment{
			Content:    content,
			ArticleId:  articleId,
			CommentUid: commentUid,
			ParentId:   parentId,
		}

		// todo check parent comment id
		// if notExistsParent  condition
		if err := db.Create(&comment).Error; err != nil {
			fmt.Println("Create comment fail", err)
			c.JSON(1001, gin.H{
				"message": "create comment fail",
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Add comment success.",
		})
	})

	// 更新留言，比如置顶, 不显示等
	// curl localhost:8080/comment/8 -X PUT -d '{"is_top":1}'
	r.PUT("/comment/:id", func(c *gin.Context) {
		id := c.Param("id")
		var comment Comment
		if result := db.Where("id = ? ", id).First(&comment); result.Error != nil || !result.RecordNotFound() {
			c.AbortWithStatus(404)
		}

		c.BindJSON(&comment)
		if err := db.Save(&comment).Error; err != nil {
			c.JSON(1005, gin.H{
				"message": "Update comment fail.",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Update comment success",
			})
		}
	})

	r.DELETE("/comment/:id", func(c *gin.Context) {
		id := c.Param("id")
		var comment Comment
		if result := db.Where("id = ? ", id).First(&comment); result.Error != nil || !result.RecordNotFound() {
			c.AbortWithStatus(404)
		}

		if err := db.Delete(&comment).Error; err != nil {
			c.JSON(1005, gin.H{
				"message": "Delete comment fail.",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "Delete comment success",
			})
		}

	})
	// todo 留言置顶

	// 回复也作为

	// todo 点赞

	// todo 关注

	r.Run()
}
