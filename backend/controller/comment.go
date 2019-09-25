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
	log.Println("===")

	log.Println(articleId)
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
		"code":    0,
		"message": "Success",
		"data":    comments,
	})
}

func GetCommentById(ctx *gin.Context) {
	id := ctx.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)

	if comment, err := CommentDao.GetCommentByID(idx); err != nil || comment == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "Record not found ",
			"data":    nil,
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    0,
			"meesage": "Get record success.",
			"data":    comment,
		})
	}

}

func CreateComment(ctx *gin.Context) {
	commentUid := ctx.PostForm("comment_uid")
	articleId := ctx.PostForm("article_id")
	content := ctx.PostForm("content")
	parentId := ctx.GetInt("parent_id")

	// todo check parent comment id
	if err := CommentDao.Create(commentUid, articleId, content, parentId); err != nil {
		log.Println("Create comment fail", err)
		ctx.JSON(1001, gin.H{
			"message": "Create comment fail",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create comment success.",
	})
}

func UpdateComment(ctx *gin.Context) {
	id := ctx.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)
	isTop := ctx.Query("is_top")

	comment, err := CommentDao.GetCommentByID(idx)
	if err != nil || comment == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    1003,
			"message": "Record not found",
		})
		return
	}
	istop, _ := strconv.Atoi(isTop)
	comment.IsTop = istop
	err = CommentDao.UpdateComment(comment)
	if err != nil {
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

func DeleteComment(ctx *gin.Context) {
	id := ctx.Param("id")
	idx, _ := strconv.ParseInt(id, 10, 64)
	var comment dao.Comment
	var err error
	response := gin.H{}
	defer ctx.JSON(http.StatusOK, response)

	if comment, err := CommentDao.GetCommentByID(idx); err != nil || comment == nil {
		response["code"] = 1003
		response["message"] = "Record not found"
		return
	}
	err = CommentDao.Delete(&comment)
	if err != nil {
		response["code"] = 1006
		response["message"] = "Delete comment fail."
		return
	}
	response["code"] = 0
	response["message"] = "Delete comment success."

}
