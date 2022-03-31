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

func ListPost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := c.MustGet(common.CurrentUser).(common.Requester)
		id := request.GetUserId()

		var filter postmodel.Filter

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewListPostBiz(store)

		result, err := biz.ListPost(c.Request.Context(), &filter, &paging, id)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.NewSuccessResponse(result, paging, filter))
	}
}
