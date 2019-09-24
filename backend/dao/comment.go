package dao

import (
	"log"

	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/load"
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

// GetComments Search comment by article_id
func (c Comment) GetComments(articleId string) (comments []*Comment, err error) {
	if err := load.Conn.Where("article_id = ? ", articleId).Find(&comments).Error; err != nil {
		return comments, nil
	}
	log.Println("++++", err)
	return nil, err
}

func (c Comment) GetCommentByID(id int) (*Comment, error) {
	println("=--===", id)
	var comment *Comment
	if result := load.Conn.Where("id = ? ", id).First(&comment); result.Error != nil || result.RecordNotFound() {
		log.Println("====", result, comment)
		return nil, result.Error
	} else {
		return comment, nil
	}
}
