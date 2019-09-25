package dao

import (
	"github.com/chenyangguang/WeChat-Official-Accounts-Comment/backend/load"
	"github.com/jinzhu/gorm"
	"strings"
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
		return nil, err
	} else {
		return comments, nil
	}
}

// GetCommentByID Get one comment by primary key
func (c Comment) GetCommentByID(id int64) (*Comment, error) {
	comment := &Comment{}
	if result := load.Conn.First(&comment, "id = ?", id); result.Error != nil || result.RecordNotFound() {
		return nil, result.Error
	} else {
		return comment, nil
	}
}

// Create add a new comment record
func (c Comment) Create(commentUid, articleId, content string, parentId int) error {
	comment := Comment{
		Content:    strings.TrimSpace(content),
		ArticleId:  articleId,
		CommentUid: commentUid,
		ParentId:   parentId,
	}
	if err := load.Conn.Create(&comment).Error; err != nil {
		return err
	}

	return nil
}

// Update update comment
func (c Comment) UpdateComment(comment *Comment) (err error) {
	if comment == nil {
		return
	}
	if err := load.Conn.Model(&Comment{}).Where("id = ?", comment.ID).Update(&comment).Error; err != nil {
		return err
	}
	return nil
}

// Delete delete comment soft
func (c Comment) Delete(comment *Comment) (err error) {
	if comment == nil {
		return
	}
	if err := load.Conn.Delete(&comment).Error; err != nil {
		return err
	}
	return nil
}
