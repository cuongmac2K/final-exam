package ginuserfollow

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/user/userstorage"
	"finnal-exam/modules/userfollow/userfollowbiz"
	"finnal-exam/modules/userfollow/userfollowmodel"
	"finnal-exam/modules/userfollow/userfollowstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowUser(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		res := c.MustGet(common.CurrentUser).(common.Requester)

		data := userfollowmodel.UserFollow{
			Id:     int(uid.GetLocalID()),
			UserId: res.GetUserId(),
		}
		db := appCtx.GetMainDBConnection()
		store := userfollowstorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := userstorage.NewSQLStore(db)
		biz := userfollowbiz.NewUserFollowUserBiz(store, incStore)

		if err := biz.FollowUser(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
