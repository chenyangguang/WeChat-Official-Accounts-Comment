package router

import (
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	// todo 公众号信息

	// todo 文章
	//r.GET("/article", controller.GetArticleById)

	// curl localhost:8080/comments?article_id=article_1
	// 查询一篇文章中的留言列表
	r.GET("/comments", controller.GetAllComments)

	// 单条留言
	// curl localhost:8080/comment/4 -X GET
	r.GET("/comment/:id", controller.GetCommentById)

	/*
		// curl localhost:8080/comment -X POST -d "content=3content&article_id=article_3&comment_uid=uid_3"
		// 新增留言
		r.POST("/comment", controller.AddComment)

		// 更新留言，比如置顶, 不显示等
		// curl localhost:8080/comment/8 -X PUT -d '{"is_top":1}'
		r.PUT("/comment/:id", controller.UpdateComment)

		r.DELETE("/comment/:id", controller.DeleteCommentByID)

		// 回复也作为

		// todo 点赞

		// todo 关注
	*/

	return r

}
