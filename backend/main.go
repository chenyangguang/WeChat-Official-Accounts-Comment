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
	Content    string
	ArticleId  string
	CommentUid string
	ParentId   int
	IsTop      uint
	Status     uint
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

	r := gin.Default()

	// todo 公众号信息

	// todo 文章
	r.GET("/article", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "",
		})
	})

	// 查询一篇文章中的留言列表
	r.GET("/comment", func(c *gin.Context) {
		articleId := c.Query("article_id")

		var comments = []Comment{}
		if err := db.Where(&Comment{ArticleId: articleId}).Find(&comments).Error; err != nil {
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

	// 新增留言
	r.POST("/comment", func(c *gin.Context) {
		commentUid := c.PostForm("comment_uid")
		articleId := c.PostForm("article_id")
		content := c.PostForm("content")
		//parentId := c.Query("parent_id")

		// do save comment
		comment := &Comment{
			Content:    content,
			ArticleId:  articleId,
			CommentUid: commentUid,
			//ParentId:   parentId,
		}
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

	// todo 留言置顶

	// 回复也作为

	// todo 点赞

	// todo 关注

	r.Run()
}
