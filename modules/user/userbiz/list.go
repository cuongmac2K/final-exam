package userbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

type User interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *usermodel.FindUser,
		paging *common.Paging,
		moreKeys ...string,
	) ([]usermodel.UserInfo, error)
}

type userBiz struct {
	user User
}

func NewUserBiz(user User) *userBiz {
	return &userBiz{user: user}
}

func (biz *userBiz) ListUserByCondition(
	ctx context.Context,
	filter *usermodel.FindUser,
	paging *common.Paging,
) ([]usermodel.UserInfo, error) {
	result, err := biz.user.ListDataByCondition(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity("users", err)
	}

	return result, err
}
