package postlikemodel

import (
	"errors"
	"finnal-exam/common"
	"fmt"
	"time"
)

const EntityName = "UserLikePost"

type PostLike struct {
	Id        int                `json:"id" gorm:"column:id;"`
	UserId    int                `json:"-" gorm:"column:user_id;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func ErrCannotLikePost(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot like this Post"),
		fmt.Sprintf("ErrCannotLikePost"),
	)
}

func ErrCannotUnlikePost(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("Cannot unlike this Post"),
		fmt.Sprintf("ErrCannotUnlikePost"),
	)
}
func ErrLikeDifferentOne(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("LikeCount  more than Zero "),
		fmt.Sprintf("ErrCannotUnlikePost"),
	)
}
func ErrDoesNotExistPostLike(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("PostLike does not exist "),
		fmt.Sprintf("ErrDoesNotExistPostLike"),
	)
}

var ErrPostNotExist = common.NewCustomError(
	errors.New("you can't like this post because it doesn't exist"),
	"post does not exist",
	"ErrPostNotExist",
)
