package postmodel

import (
	"errors"
	"finnal-exam/common"
)

const EntityName = "Post"

type Post struct {
	common.SQLModel `json:",inline"`
	Id              int                    `json:"id" gorm:"column:id;"`
	UserId          int                    `json:"-" gorm:"column:user_id;"`
	Caption         string                 `json:"caption" gorm:"column:caption;"`
	Images          *common.Images         `json:"images" gorm:"column:images;"`
	User            *common.SimpleUser     `json:"user" gorm:"preload:false;"`
	Comments        *common.SimpleComments `json:"comments" gorm:"preload:false;foreignKey:Id"`
	LikesCount      int                    `json:"likes_count" gorm:"column:likes_count;"` //computed field
}

type PostDetails struct {
	common.SQLModel `json:",inline"`
	Id              int                    `json:"id" gorm:"column:id;"`
	UserId          int                    `json:"-" gorm:"column:user_id;"`
	Caption         string                 `json:"caption" gorm:"column:caption;"`
	Images          *common.Images         `json:"images" gorm:"column:images;"`
	User            *common.SimpleUser     `json:"user" gorm:"preload:false;"`
	Comments        *common.SimpleComments `json:"comments" gorm:"preload:false;foreignKey:Id"`
	LikesCount      int                    `json:"likes_count" gorm:"column:likes_count;"` //computed field
}
type PostTrend struct {
	Caption    string         `json:"caption" gorm:"column:caption;"`
	Images     *common.Images `json:"images" gorm:"column:images;"`
	LikesCount int            `json:"likes_count" gorm:"column:likes_count;"` //computed field
}

func (Post) TableName() string {
	return "posts"
}

type PostCreate struct {
	common.SQLModel `json:",inline"`
	UserId          int                `json:"-" gorm:"column:user_id;"`
	Caption         string             `json:"caption" gorm:"column:caption;"`
	Images          *common.Images     `json:"images" gorm:"column:images;"`
	User            *common.SimpleUser `json:"user" gorm:"preload:false;"`
	LikesCount      int                `json:"likes_count" gorm:"column:likes_count;"` // computed field
}

func (PostCreate) TableName() string {
	return Post{}.TableName()
}

type PostUpdate struct {
	common.SQLModel `json:",inline"`
	Caption         *string        `json:"caption,omitempty" gorm:"column:caption;"`
	Images          *common.Images `json:"images,omitempty" gorm:"column:images;"`
}

func (PostUpdate) TableName() string {
	return Post{}.TableName()
}

func (data *Post) Mask(isAdminOrOwner bool) {
	data.GenUID(common.DbTypePost)

	if u := data.User; u != nil {
		u.Mask(isAdminOrOwner)
	}
}

type PostCheck struct {
	common.SQLModel `json:",inline"`
	UserId          int `json:"-" gorm:"column:user_id;"`
	LikesCount      int `json:"likes_count" gorm:"column:likes_count;"` // computed field
}

var (
	ErrNoDeletePermission = common.NewCustomError(
		errors.New("you don't have permission to delete this post"),
		"you don't have permission to delete this post",
		"ErrNoDeletePermission",
	)
	ErrNoUpdatePermission = common.NewCustomError(
		errors.New("you don't have permission to update this post"),
		"you don't have permission to update this post",
		"ErrNoUpdatePermission",
	)
)
