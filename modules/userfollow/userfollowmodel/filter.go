package userfollowmodel

type Filter struct {
	Id     int `json:"-" form:"id"`
	UserId int `json:"-" form:"user_id"`
}
