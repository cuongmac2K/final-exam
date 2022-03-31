package ginuser

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/user/userbiz"
	"finnal-exam/modules/user/userstorage"
	"github.com/gin-gonic/gin"

	"net/http"
)

func GetUserbyID(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewGetUserBiz(store)

		data, err := biz.GetUser(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
