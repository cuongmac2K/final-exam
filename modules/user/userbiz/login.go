package userbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/component/tokenprovider"
	"finnal-exam/modules/user/usermodel"
)

type LoginStore interface {
	FindUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}
type loginBiz struct {
	appCtx        component.AppContext
	storeUser     LoginStore
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginBiz(storeUser LoginStore,
	tokenProvider tokenprovider.Provider,
	hasher Hasher, expiry int) *loginBiz {
	return &loginBiz{storeUser: storeUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}
func (biz *loginBiz) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := biz.storeUser.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	passHashed := biz.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}

	if user.IsVerified == 0 {
		return nil, usermodel.ErrEmailNotVerified
	}

	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, common.ErrInternal(err)
	}

	return accessToken, nil
}
