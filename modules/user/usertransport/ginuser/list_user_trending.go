package ginuser

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/user/userbiz"
	"finnal-exam/modules/user/usermodel"
	"finnal-exam/modules/user/userstorage"
	"github.com/gin-gonic/gin"

	"net/http"
)

func ListUserTrend(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var filter usermodel.Trending

		if err := c.ShouldBind(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		var paging common.Paging

		if err := c.ShouldBind(&paging); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		paging.Fulfill()

		db := userstorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := userbiz.NewUserTrendingBiz(db)

		result, err := biz.ListUserTrending(c.Request.Context(), &filter, &paging)

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(result))
	}
}
