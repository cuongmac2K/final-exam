package commentlikemodel

import (
	"errors"
	"finnal-exam/common"
	"fmt"
	"time"
)

const EntityName = "UserLikeComment"

type CommentLike struct {
	Id        int                `json:"id" gorm:"column:id;"`
	UserId    int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (CommentLike) TableName() string {
	return "comment_likes"
}

var (
	ErrUnlikeComment = common.NewCustomError(
		errors.New("you have already unliked this comment"),
		"you have already unliked this comment",
		"ErrUnlikeComment",
	)
)

func ErrCannotLikeComment(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this Comment"),
		fmt.Sprintf("ErrCannotLikeComment"),
	)
}
func ErrCommentDoesntexist(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("comment doesn't exist"),
		fmt.Sprintf("ErrCannotLikeComment"),
	)
}
