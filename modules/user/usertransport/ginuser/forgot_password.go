package ginuser

import (
	"errors"
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/component/tokenprovider/jwt"
	"finnal-exam/modules/user/userbiz"
	"finnal-exam/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ForgotPassword(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		email := c.Param("email")
		if !common.Valid(email) {
			common.ErrEmailNotRightFormat(errors.New("wrong input"))
		}
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		emailProvider := common.NewEmailProvider(appCtx.EmailUsername(), appCtx.EmailPassword())
		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewForgotPasswordBiz(store, emailProvider, tokenProvider, 3600)
		if err := biz.ForgotPassword(c.Request.Context(), email); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
