package ginpostlike

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/post/poststorage"
	"finnal-exam/modules/postlike/postlikebiz"
	"finnal-exam/modules/postlike/postlikemodel"
	"finnal-exam/modules/postlike/postlikestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLikePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := postlikemodel.PostLike{
			Id:     int(uid.GetLocalID()),
			UserId: requester.GetUserId(),
		}

		store := postlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postlikebiz.NewUserLikePostBiz(store, incStore)

		if err := biz.LikePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
