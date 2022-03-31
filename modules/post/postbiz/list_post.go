package postbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/post/postmodel"
)

type ListPostStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		id int,
		filter *postmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.Post, error)
}

type listPostBiz struct {
	store ListPostStore
}

func NewListPostBiz(store ListPostStore) *listPostBiz {
	return &listPostBiz{store: store}
}

func (biz *listPostBiz) ListPost(
	ctx context.Context,
	filter *postmodel.Filter,
	paging *common.Paging,
	id int,
) ([]postmodel.Post, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, id, filter, paging, "User", "Comments")

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return result, err
}
