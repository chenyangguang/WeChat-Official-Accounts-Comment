package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/dao"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetAllComments(ctx *gin.Context) {
	articleId := strings.TrimSpace(ctx.Query("article_id"))
	comments, err := CommentDao.GetComments(articleId)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1002,
			"message": "Query comment fail",
			"data":    comments,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":          0,
		"message":       "Success",
		"comment_lists": comments,
	})
}

func GetCommentById(ctx *gin.Context) {
	id := ctx.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)

	var comment dao.Comment
	if comment, err := CommentDao.GetCommentByID(idx); err != nil || comment == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "Record not found ",
			"data":    nil,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"meesage": "Get record success.",
		"data":    comment,
	})

}

func AddComment(ctx *gin.Context) {
	commentUid := ctx.PostForm("comment_uid")
	articleId := ctx.PostForm("article_id")
	content := ctx.PostForm("content")
	parentId := ctx.GetInt("parent_id")

	// todo check parent comment id
	if err := CommentDao.AddComment(commentUid, articleId, content, parentId); err != nil {
		log.Println("Create comment fail", err)
		ctx.JSON(1001, gin.H{
			"message": "Create comment fail",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create comment success.",
	})
}

/*
	func (c *CommentController) UpdateComment(ctx *gin.Context) {
		id := ctx.Param("id")
		var comment Comment
		if result := model.Where("id = ? ", id).First(&comment); result.Error != nil || !result.RecordNotFound() {
			ctx.AbortWithStatus(404)
		}

		ctx.BindJSON(&comment)
		if err := model.Save(&comment).Error; err != nil {
			ctx.JSON(1005, gin.H{
				"message": "Update comment fail.",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Update comment success",
			})
		}
	}

	func (c *CommentController) DeleteCommentByID(ctx *gin.Context) {
		id := ctx.Param("id")
		var comment Comment
		if result := model.Where("id = ? ", id).First(&comment); result.Error != nil || !result.RecordNotFound() {
			ctx.AbortWithStatus(404)
		}

		if err := model.Delete(&comment).Error; err != nil {
			ctx.JSON(1005, gin.H{
				"message": "Delete comment fail.",
			})
		} else {
			ctx.JSON(http.StatusOK, gin.H{
				"message": "Delete comment success",
			})
		}

	}

	/*


		r.DELETE("/comment/:id", func(c *gin.Context) {
			id := c.Param("id")
			var comment Comment
			if result := model.Where("id = ? ", id).First(&comment); result.Error != nil || !result.RecordNotFound() {
				c.AbortWithStatus(404)
			}

			if err := model.Delete(&comment).Error; err != nil {
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

*/
