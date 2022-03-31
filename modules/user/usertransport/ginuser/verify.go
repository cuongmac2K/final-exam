package ginuser

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/user/userbiz"
	"finnal-exam/modules/user/usermodel"
	"finnal-exam/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func VerifyEmail(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter usermodel.VerifyUser
		if err := c.ShouldBind(&filter); err != nil {
			common.ErrInvalidRequest(err)
		}
		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewVerifyEmailBiz(store)

		if err := biz.VerifyEmail(c.Request.Context(), filter.Email, filter.Otp); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
