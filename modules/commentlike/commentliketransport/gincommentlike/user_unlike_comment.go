package gincommentlike

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/comment/commentstorage"
	"finnal-exam/modules/commentlike/commentlikebiz"
	"finnal-exam/modules/commentlike/commentlikestorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnlikeComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		res := c.MustGet(common.CurrentUser).(common.Requester)
		userId := res.GetUserId()

		db := appCtx.GetMainDBConnection()

		store := commentlikestorage.NewSQLStore(db)
		decLikeStore := commentstorage.NewSQLStore(db)
		biz := commentlikebiz.NewUserUnlikeCommentBiz(store, decLikeStore)

		if err := biz.UnlikeComment(c.Request.Context(), int(uid.GetLocalID()), userId); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
