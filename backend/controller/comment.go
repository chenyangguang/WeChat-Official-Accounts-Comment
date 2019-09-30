package controller

import (
	"net/http"
	"strconv"

	"strings"

	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/dao"
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/load/log"
	"github.com/gin-gonic/gin"
)

// GetAllComments 获取所有留言
// todo 应该改成分页
func GetAllComments(ctx *gin.Context) {
	articleId := strings.TrimSpace(ctx.Query("article_id"))
	comments, err := CommentDao.GetComments(articleId)
	if err != nil {
		log.Logger.Info(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1002,
			"message": "Query comment fail",
			"data":    comments,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Success",
		"data":    comments,
	})
}

// GetCommentById 获取单条留言
func GetCommentById(ctx *gin.Context) {
	id := ctx.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)

	if comment, err := CommentDao.GetCommentByID(idx); err != nil || comment == nil {
		log.Logger.Info(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "Record not found ",
			"data":    nil,
		})
		return
	} else {
		log.Logger.Info(comment)

		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "Get record success.",
			"data":    comment,
		})
	}

}

// CreateComment 写留言
func CreateComment(ctx *gin.Context) {
	commentUid := ctx.PostForm("comment_uid")
	articleId := ctx.PostForm("article_id")
	content := ctx.PostForm("content")
	parentId := ctx.GetInt("parent_id")

	// todo check parent comment id
	if err := CommentDao.Create(commentUid, articleId, content, parentId); err != nil {
		log.Logger.Info(err)
		ctx.JSON(1001, gin.H{
			"message": "Create comment fail",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create comment success.",
	})
}

// UpdateComment 更新留言信息
func UpdateComment(ctx *gin.Context) {
	id := ctx.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)
	comment, err := CommentDao.GetCommentByID(idx)
	if err != nil || comment == nil {
		log.Logger.Info(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1003,
			"message": "Record not found",
		})
		return
	}
	ctx.BindJSON(&comment)

	err = CommentDao.UpdateComment(comment)
	if err != nil {
		log.Logger.Info(err)
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1005,
			"message": "Update comment fail",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "Update comment success",
	})
}

// DeleteComment 删除留言
func DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)
	var comment dao.Comment
	var err error
	response := gin.H{}
	defer ctx.JSON(http.StatusOK, response)

	if comment, err := CommentDao.GetCommentByID(idx); err != nil || comment == nil {
		log.Logger.Info(err)
		response["code"] = 1003
		response["message"] = "Record not found"
		return
	}
	err = CommentDao.Delete(&comment)
	if err != nil {
		log.Logger.Info(err)
		response["code"] = 1006
		response["message"] = "Delete comment fail."
		return
	}
	response["code"] = 0
	response["message"] = "Delete comment success."

}
