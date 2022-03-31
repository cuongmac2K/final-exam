package userbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

type UserTrending interface {
	ListUserByFollwer(ctx context.Context,
		conditions map[string]interface{},
		filter *usermodel.Trending,
		paging *common.Paging,
		moreKeys ...string,
	) ([]usermodel.UserTren, error)
}

type userTrendingBiz struct {
	user UserTrending
}

func NewUserTrendingBiz(user UserTrending) *userTrendingBiz {
	return &userTrendingBiz{user: user}
}

func (biz *userTrendingBiz) ListUserTrending(
	ctx context.Context,
	filter *usermodel.Trending,
	paging *common.Paging,
) ([]usermodel.UserTren, error) {
	result, err := biz.user.ListUserByFollwer(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListUserTrendEntity("users", err)
	}
	return result, err
}
