package commentlikemodel

type Filter struct {
	Id     int `json:"_" form:"id"`
	UserId int `json:"-" form:"user_id"`
}
