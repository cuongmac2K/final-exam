package common

type SimpleUser struct {
	SQLModel  `json:",inline"`
	Id        string `json:"-" gorm:"column:id;"`
	LastName  string `json:"last_name" gorm:"column:last_name;"`
	FirstName string `json:"first_name" gorm:"column:first_name;"`
	Avatar    *Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (SimpleUser) TableName() string {
	return "users"
}

func (u *SimpleUser) Mask(isAdmin bool) {
	u.GenUID(DbTypeUser)
}
