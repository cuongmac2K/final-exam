package userbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

type GetUserStore interface {
	GetUserByID(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
}

type getUserbiz struct {
	store GetUserStore
}

func NewGetUserBiz(store GetUserStore) *getUserbiz {
	return &getUserbiz{store: store}
}

func (biz *getUserbiz) GetUser(ctx context.Context, id int) (*usermodel.User, error) {
	data, err := biz.store.GetUserByID(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return nil, common.ErrCannotGetEntity(usermodel.EntityName, err)
	}
	return data, err
}
