package userbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

type UpdateUserStore interface {
	FindUser(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*usermodel.User, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *usermodel.UserUpdate,
	) error
}

type updateUserBiz struct {
	store UpdateUserStore
}

func NewUpdateUserBiz(store UpdateUserStore) *updateUserBiz {
	return &updateUserBiz{store: store}
}

func (biz *updateUserBiz) UpdateUser(ctx context.Context, id int, data *usermodel.UserUpdate) error {
	oldData, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrCannotGetEntity(usermodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(usermodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(usermodel.EntityName, err)
	}

	return nil
}
