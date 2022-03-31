package gincomment

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/comment/commentbiz"
	"finnal-exam/modules/comment/commentmodel"
	"finnal-exam/modules/comment/commentstorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var data commentmodel.CommentCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		res := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = res.GetUserId()

		store := commentstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := commentbiz.NewCreateCommentBiz(store)

		if err := biz.CreateComment(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypeComment)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
