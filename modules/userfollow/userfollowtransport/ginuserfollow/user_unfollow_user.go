package ginuserfollow

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/user/userstorage"
	"finnal-exam/modules/userfollow/userfollowbiz"
	"finnal-exam/modules/userfollow/userfollowstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnFollowUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		res := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMainDBConnection()
		store := userfollowstorage.NewSQLStore(appCtx.GetMainDBConnection())
		deincStore := userstorage.NewSQLStore(db)
		biz := userfollowbiz.NewUserUnFollowUserBiz(store, deincStore)

		if err := biz.UnFollowUser(c.Request.Context(), res.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
