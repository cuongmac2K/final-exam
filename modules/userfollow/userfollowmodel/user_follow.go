package userfollowmodel

import (
	"errors"
	"finnal-exam/common"
	"time"
)

const EntityName = "UserFollowUser"

type UserFollow struct {
	Id        int                `json:"id" gorm:"column:id;"`
	UserId    int                `json:"user_id" gorm:"column:user_id;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (uf UserFollow) TableName() string {
	return "user_follows"
}

type UserFollowing struct {
	Id        int                `json:"user_id" gorm:"column:user_id;"`
	UserId    int                `json:"id" gorm:"column:id;"`
	CreatedAt *time.Time         `json:"created_at" gorm:"column:created_at;"`
	User      *common.SimpleUser `json:"user" gorm:"preload:false;"`
}

func (uf UserFollowing) TableName() string {
	return "user_follows"
}

var (
	ErrFollowExisted = common.NewCustomError(
		errors.New("you have already followed this user"),
		"you have already followed this user",
		"ErrFollowExisted",
	)
	ErrUnFollowExisted = common.NewCustomError(
		errors.New("you can't unfollow this user"),
		"you haven't followed this user yet",
		"ErrUnFollowExisted",
	)
	ErrUserNotAvailable = common.NewCustomError(
		errors.New("can't find this user"),
		"this user is not available",
		"ErrUserNotAvailable",
	)
	ErrUserFollowMyself = common.NewCustomError(
		errors.New("can't follow yourself"),
		"you can't follow yourself",
		"ErrUserFollowMyself",
	)
)
