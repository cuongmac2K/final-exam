package common

type SimpleComment struct {
	SQLModel   `json:",inline"`
	UserId     int         `json:"-" gorm:"column:user_id;"`
	Content    string      `json:"content" gorm:"column:content;"`
	SimpleUser *SimpleUser `json:"user" gorm:"preload:false;foreignKey:UserId"`
}

func (SimpleComment) TableName() string {
	return "comments"
}

type SimpleComments []SimpleComment
