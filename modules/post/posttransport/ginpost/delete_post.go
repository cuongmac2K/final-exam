package ginpost

import (
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/modules/post/postbiz"
	"finnal-exam/modules/post/poststorage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DeletePost(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		req := c.MustGet(common.CurrentUser).(common.Requester)
		userid := req.GetUserId()
		//id, err := strconv.Atoi(c.Param("id"))
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := poststorage.NewSQLStore(appCtx.GetMainDBConnection())
		biz := postbiz.NewDeletePostBiz(store)

		if err := biz.DeletePost(c.Request.Context(), int(uid.GetLocalID()), userid); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
