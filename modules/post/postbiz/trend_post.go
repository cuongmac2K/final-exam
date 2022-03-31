package postbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/post/postmodel"
)

type TrendPostStore interface {
	ListTrendByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *postmodel.FilterTrend,
		paging *common.Paging,
		moreKeys ...string,
	) ([]postmodel.PostTrend, error)
}

type trendPostBiz struct {
	store TrendPostStore
}

func NewTrendPostBiz(store TrendPostStore) *trendPostBiz {
	return &trendPostBiz{store: store}
}

func (biz *trendPostBiz) ListTrendPost(
	ctx context.Context,
	filter *postmodel.FilterTrend,
	paging *common.Paging,
) ([]postmodel.PostTrend, error) {
	result, err := biz.store.ListTrendByCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(postmodel.EntityName, err)
	}

	return result, err
}
