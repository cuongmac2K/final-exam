package usermodel

import (
	"errors"
	"finnal-exam/common"
	"finnal-exam/component/tokenprovider"
	"time"
)

const EntityName = "User"

type User struct {
	common.SQLModel `json:",inline"`
	Email           string         `json:"email" gorm:"column:email;"`
	Password        string         `json:"-" gorm:"column:password;"`
	Salt            string         `json:"-" gorm:"column:salt;"`
	LastName        string         `json:"last_name" gorm:"column:last_name;"`
	FirstName       string         `json:"first_name" gorm:"column:first_name;"`
	Phone           string         `json:"phone" gorm:"column:phone;"`
	BirthDay        *time.Time     `json:"birthday" gorm:"column:birthday;"`
	Role            string         `json:"role" gorm:"column:role;"`
	Avatar          *common.Image  `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	IsVerified      int            `json:"-" gorm:"column:is_verified;"`
	Record          *common.Record `json:"-" gorm:"column:record;type:json"`
	FollowerCount   int            `json:"follower_count" gorm:"follower_count"`
	FollowingCount  int            `json:"following_count" gorm:"following_count"`
}

func (u *User) GetUserId() int {
	return u.Id
}

func (u *User) GetEmail() string {
	return u.Email
}

func (u *User) GetRole() string {
	return u.Role
}

func (User) TableName() string {
	return "users"
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserCreate struct {
	common.SQLModel `json:",inline"`
	Email           string         `json:"email" gorm:"column:email;"`
	Password        string         `json:"password" gorm:"column:password;"`
	LastName        string         `json:"last_name" gorm:"column:last_name;"`
	FirstName       string         `json:"first_name" gorm:"column:first_name;"`
	Role            string         `json:"-" gorm:"column:role;"`
	Phone           string         `json:"phone" gorm:"column:phone;"`
	BirthDay        *time.Time     `json:"birthday" gorm:"column:birthday;"`
	Salt            string         `json:"-" gorm:"column:salt;"`
	Avatar          *common.Image  `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
	Record          *common.Record `json:"-" gorm:"column:record;type:json"`
}
type UserInfo struct {
	Email     string        `json:"email" gorm:"column:email;"`
	LastName  string        `json:"last_name" gorm:"column:last_name;"`
	FirstName string        `json:"first_name" gorm:"column:first_name;"`
	Phone     string        `json:"phone" gorm:"column:phone;"`
	BirthDay  *time.Time    `json:"birthday" gorm:"column:birthday;"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserCreate) TableName() string {
	return User{}.TableName()
}

func (u *UserCreate) Mask(isAdmin bool) {
	u.GenUID(common.DbTypeUser)
}

type UserUpdate struct {
	common.SQLModel `json:",inline"`

	Password  *string       `json:"password" gorm:"column:password;"`
	FirstName *string       `json:"first_name" gorm:"column:first_name;"`
	LastName  *string       `json:"last_name" gorm:"column:last_name;"`
	Phone     *string       `json:"phone" gorm:"column:phone;"`
	Birthday  *time.Time    `json:"birthday" gorm:"birthday"`
	Avatar    *common.Image `json:"avatar,omitempty" gorm:"column:avatar;type:json"`
}

func (UserUpdate) TableName() string {
	return User{}.TableName()
}

type UserLogin struct {
	Email    string `json:"email" form:"email" gorm:"column:email;"`
	Password string `json:"password" form:"password" gorm:"column:password;"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}

type Account struct {
	AccessToken  *tokenprovider.Token `json:"access_token"`
	RefreshToken *tokenprovider.Token `json:"refresh_token"`
}

func NewAccount(at, rt *tokenprovider.Token) *Account {
	return &Account{
		AccessToken:  at,
		RefreshToken: rt,
	}
}

var (
	ErrUsernameOrPasswordInvalid = common.NewCustomError(
		errors.New("username or password invalid"),
		"username or password invalid",
		"ErrUsernameOrPasswordInvalid",
	)

	ErrEmailExisted = common.NewCustomError(
		errors.New("email has already existed"),
		"email has already existed",
		"ErrEmailExisted",
	)
	ErrOtpIncorrect = common.NewCustomError(
		errors.New("your otp code isn't correct"),
		"incorrect otp",
		"ErrOtpIncorrect",
	)
	ErrEmailNotVerified = common.NewCustomError(
		errors.New("your email isn't verified"),
		"email is not verified",
		"ErrEmailNotVerified",
	)
	ErrOtpExpired = common.NewCustomError(
		errors.New("expired otp code"),
		"your otp code is expired",
		"ErrOtpExpired",
	)
)
