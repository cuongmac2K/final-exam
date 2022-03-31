package commentmodel

import (
	"errors"
	"finnal-exam/common"
)

const EntityName = "Comment"

type Comment struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	PostId          int    `json:"post_id" gorm:"column:post_id;"`
	Content         string `json:"content" gorm:"column:content;"`
	ParentCmt       int    `json:"parents_cmt" gorm:"column:parents_cmt;"`
	LikesCount      int    `json:"likes_count" gorm:"column:likes_count;"`
	Level           int    `json:"level" gorm:"column:level;"`
}

func (Comment) TableName() string {
	return "comments"
}

type CommentCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int    `json:"-" gorm:"column:user_id;"`
	PostId          int    `json:"post_id" gorm:"column:post_id;"`
	Content         string `json:"content" gorm:"column:content;"`
	ParentCmt       int    `json:"parents_cmt" gorm:"column:parents_cmt;"`
	Level           int    `json:"-" gorm:"column:level;"`
}

func (CommentCreate) TableName() string {
	return Comment{}.TableName()
}

var (
	ErrCannotReplyComment = common.NewCustomError(
		errors.New("you can't reply right now"),
		"can't find parent comment",
		"ErrCannotReplyComment",
	)
	ErrCannotReplyCommentPlus = common.NewCustomError(
		errors.New("you can't reply right now"),
		"Comment over two level",
		"ErrCannotReplyComment",
	)
)
