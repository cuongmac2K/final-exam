package ginuserfollow

import (
	"finnal-exam/common"
	"finnal-exam/component"
	userfollowbiz "finnal-exam/modules/userfollow/userfollowbiz"
	"finnal-exam/modules/userfollow/userfollowmodel"
	userfollowstorage "finnal-exam/modules/userfollow/userfollowstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListUserFollowing(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		//
		filter := userfollowmodel.Filter{
			UserId: int(uid.GetLocalID()),
			//Id: id,
		}
		//request := c.MustGet(common.CurrentUser).(common.Requester)
		//id := request.GetUserId()

		//var filter userfollowmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := userfollowstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userfollowbiz.NewListUserFollowingBiz(store)

		result, err := biz.ListUsersFollowing(c.Request.Context(), &filter, &paging, int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		for i := range result {
			result[i].Mask(false)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
