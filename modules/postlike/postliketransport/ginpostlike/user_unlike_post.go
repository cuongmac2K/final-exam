package ginpostlike

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/post/poststorage"
	"finnal-exam/modules/postlike/postlikebiz"
	"finnal-exam/modules/postlike/postlikestorage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserUnLikePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		//uid, err := strconv.Atoi(c.Param("id"))
		log.Println("=========", uid)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		//data := postlikemodel.PostLike{
		//	Id:     int(uid.GetLocalID()),
		//	UserId: requester.GetUserId(),
		//}

		store := postlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		dncStore := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postlikebiz.NewUserUnLikePostBiz(store, dncStore)

		if err := biz.UnLikePost(c.Request.Context(), requester.GetUserId(), int(uid.GetLocalID())); err != nil {
			panic(err)
		}
		//int(uid.GetLocalID())
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
