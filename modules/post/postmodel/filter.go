package postmodel

type Filter struct {
	Username    string `json:"username,omitempty" form:"username"`
	Email       string `json:"email,omitempty" form:"email"`
	Caption     string `json:"caption,omitempty" form:"caption"`
	DateFrom    string `json:"date_from,omitempty" form:"date_from"`
	DateTo      string `json:"date_to,omitempty" form:"date_to"`
	IsFollowing bool   `json:"is_following,omitempty" form:"is_following"`
	ImagesId    []int  `json:"images_id,omitempty" form:"images_id"`
}
type FilterTrend struct {
	Month string `json:"month,omitempty" form:"month"`
	Year  string `json:"year,omitempty" form:"year"`
	Date  string `json:"date,omitempty" form:"date"`
}
