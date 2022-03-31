package ginuser

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/component/hasher"
	"finnal-exam/component/tokenprovider/jwt"
	"finnal-exam/modules/user/userbiz"
	"finnal-exam/modules/user/usermodel"
	"finnal-exam/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUserData usermodel.UserLogin
		if err := c.ShouldBind(&loginUserData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		tokenProvider := jwt.NewTokenJWTProvider(appCtx.SecretKey())
		md5 := hasher.NewMd5Hash()

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewLoginBiz(store, tokenProvider, md5, 3600*24)
		account, err := biz.Login(c.Request.Context(), &loginUserData)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(account))
	}
}
