package userbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/user/usermodel"
)

type UpdatePasswordStore interface {
	FindUser(ctx context.Context,
		conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	UpdatePassword(
		ctx context.Context,
		id int,
		user *usermodel.User,
	) error
}
type updatePasswordBiz struct {
	store  UpdatePasswordStore
	hasher Hasher
}

func NewUpdatePasswordBiz(store UpdatePasswordStore, hasher Hasher) *updatePasswordBiz {
	return &updatePasswordBiz{store: store, hasher: hasher}
}

func (biz *updatePasswordBiz) ResetPassword(ctx context.Context, id int, password string) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	salt := common.GenSalt(50)
	user.Password = biz.hasher.Hash(password + salt)
	user.Salt = salt

	if err := biz.store.UpdatePassword(ctx, user.Id, user); err != nil {
		return err
	}

	return nil
}
