package usermodel

import "finnal-exam/common"

type FindUser struct {
	Email     string `json:"email" form:"email"`
	LastName  string `json:"last_name" form:"last_name"`
	FirstName string `json:"first_name" form:"first_name"`
	Phone     string `json:"phone" form:"phone"`
}
type Trending struct {
	Day   string `json:"day" form:"day"`
	Month string `json:"month "form:"month"`
	Year  string `json:"year" form:"year"`
	Limit int    `json:"limit" form:"limit"`
}
type UserTren struct {
	Id              int           `json:"id" gorm:"column:id;"`
	LastName        string        `json:"last_name" gorm:"column:last_name;"`
	FirstName       string        `json:"first_name" gorm:"column:first_name;"`
	Avatar          *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	common.SQLModel `json:",inline"`
}
type VerifyUser struct {
	Email string `json:"email" form:"email"`
	Otp   string `json:"otp" form:"otp"`
}
