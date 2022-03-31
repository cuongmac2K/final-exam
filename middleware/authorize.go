package middleware

import (
	"context"
	"errors"
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/component/tokenprovider/jwt"
	"finnal-exam/modules/user/usermodel"
	"finnal-exam/modules/user/userstorage"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type AuthenStore interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authen header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use user_id to find from DB
func RequiredAuth(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}

func RequiredAuthEmail(appCtx component.AppContext) func(c *gin.Context) {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())

	return func(c *gin.Context) {
		type Filter struct {
			Token string `json:"token" form:"code"`
		}
		var token Filter
		if err := c.ShouldBind(&token); err != nil {
			panic(err)
		}
		db := appCtx.GetMainDBConnection()
		store := userstorage.NewSQLStore(db)
		fmt.Println("token", token.Token)
		payload, err := tokenProvider.Validate(token.Token)
		if err != nil {
			panic(err)
		}

		user, err := store.FindUser(c.Request.Context(), map[string]interface{}{"id": payload.UserId})

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(common.ErrNoPermission(errors.New("user has been deleted or banned")))
		}

		user.Mask(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}

}
