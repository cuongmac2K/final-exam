package gincommentlike

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/comment/commentstorage"
	"finnal-exam/modules/commentlike/commentlikebiz"
	"finnal-exam/modules/commentlike/commentlikemodel"
	"finnal-exam/modules/commentlike/commentlikestorage"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func UserLikeComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		log.Println("==========", uid)
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		requester := c.MustGet(common.CurrentUser).(common.Requester)

		data := commentlikemodel.CommentLike{
			Id:     int(uid.GetLocalID()),
			UserId: requester.GetUserId(),
		}
		log.Println("====", data)
		store := commentlikestorage.NewSQLStore(appCtx.GetMainDBConnection())
		incStore := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentlikebiz.NewUserLikeCommentBiz(store, incStore)

		if err := biz.LikeComment(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
