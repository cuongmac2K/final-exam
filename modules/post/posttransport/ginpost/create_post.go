package ginpost

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/post/postbiz"
	"finnal-exam/modules/post/postmodel"
	"finnal-exam/modules/post/poststorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreatePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data postmodel.PostCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		res := c.MustGet(common.CurrentUser).(common.Requester)
		data.UserId = res.GetUserId()

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewCreatePostBiz(store)

		if err := biz.CreatePost(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.GenUID(common.DbTypePost)
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
