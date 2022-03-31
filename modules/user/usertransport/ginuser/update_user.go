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

func UpdateUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		req := c.MustGet(common.CurrentUser).(common.Requester)
		userid := req.GetUserId()

		var data usermodel.UserUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewUpdateUserBiz(store)

		if err := biz.UpdateUser(c.Request.Context(), userid, &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
