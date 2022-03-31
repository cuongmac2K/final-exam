package ginuser

import (
	"errors"
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/component/hasher"
	"finnal-exam/modules/user/userbiz"
	"finnal-exam/modules/user/usermodel"
	"finnal-exam/modules/user/userstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(appCtx component.AppContext) func(*gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMainDBConnection()
		var data usermodel.UserCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		if !common.Valid(data.Email) {
			panic(common.ErrInvalidRequest(errors.New("Email is invalid")))
		}
		emailProvider := common.NewEmailProvider(appCtx.EmailUsername(), appCtx.EmailPassword())

		store := userstorage.NewSQLStore(db)
		md5 := hasher.NewMd5Hash()
		biz := userbiz.NewRegisterBusiness(store, md5, emailProvider)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mask(false)

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
