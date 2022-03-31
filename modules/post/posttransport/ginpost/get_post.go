package ginpost

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/post/postbiz"
	"finnal-exam/modules/post/poststorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetPost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewGetPostBiz(store)

		data, err := biz.GetPost(c.Request.Context(), int(uid.GetLocalID()))

		if err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(data))
	}
}
