package component

import (
	"gorm.io/gorm"
)

type AppContext interface {
	GetMainDBConnection() *gorm.DB
	SecretKey() string
	EmailUsername() string
	EmailPassword() string
}

type appCtx struct {
	db            *gorm.DB
	secretKey     string
	emailUsername string
	emailPassword string
}

func NewAppContext(db *gorm.DB, secretKey string, emailUsername string, emailPassword string) *appCtx {
	return &appCtx{db: db, secretKey: secretKey, emailUsername: emailUsername, emailPassword: emailPassword}
}

func (ctx *appCtx) GetMainDBConnection() *gorm.DB {
	return ctx.db
}

func (ctx *appCtx) SecretKey() string     { return ctx.secretKey }
func (ctx *appCtx) EmailUsername() string { return ctx.emailUsername }
func (ctx *appCtx) EmailPassword() string { return ctx.emailPassword }
