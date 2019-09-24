package controller

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllComments(ctx *gin.Context) {
	articleId := ctx.Query("article_id")

	//var comments = []Comment{}
	if comments, err := CommentDao.GetComments(articleId); err != nil {
		log.Println(err)
		ctx.JSON(1002, gin.H{
			"message": "Query comment fail",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message":       "Success",
			"comment_lists": comments,
		})
	}
}

func GetCommentById(ctx *gin.Context) {
	id := ctx.Param("id")
	idx, _ := strconv.Atoi(id)
	if comment, err := CommentDao.GetCommentByID(idx); err != nil || comment == nil {
		log.Println(err)
		//ctx.AbortWithStatus(404)
		ctx.AbortWithError(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"meesage": "Get record success.",
			"comment": comment,
		})
	}
}

/*
func (c *CommentController) AddComment(ctx *gin.Context) {
	commentUid := ctx.PostForm("comment_uid")
	articleId := ctx.PostForm("article_id")
	content := ctx.PostForm("content")
	parentId := ctx.GetInt("parent_id")

	// do save comment
	comment := &Comment{
		Content:    content,
		ArticleId:  articleId,
		CommentUid: commentUid,
		ParentId:   parentId,
	}

	// todo check parent comment id
	// if notExistsParent  condition
	if err := model.Create(&comment).Error; err != nil {
		log.Println("Create comment fail", err)
		ctx.JSON(1001, gin.H{
			"message": "create comment fail",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Add comment success.",
	})
}

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
