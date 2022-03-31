package userbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/component"
	"finnal-exam/component/asyncjob"
	"finnal-exam/component/tokenprovider"
	"finnal-exam/modules/user/usermodel"
)

type ForgotPasswordStore interface {
	FindUser(ctx context.Context,
		conditions map[string]interface{},
		moreInfo ...string) (*usermodel.User, error)
}
type forgotPasswordBiz struct {
	appCtx              component.AppContext
	storeForgotPassword ForgotPasswordStore
	emailProvider       common.SendEmailProvider
	tokenProvider       tokenprovider.Provider
	expiry              int
}

func NewForgotPasswordBiz(storeForgotPassword ForgotPasswordStore, emailProvider common.SendEmailProvider,
	tokenProvider tokenprovider.Provider,
	expiry int) *forgotPasswordBiz {
	return &forgotPasswordBiz{storeForgotPassword: storeForgotPassword,
		emailProvider: emailProvider,
		tokenProvider: tokenProvider,
		expiry:        expiry,
	}
}
func (biz *forgotPasswordBiz) ForgotPassword(ctx context.Context, email string) error {
	user, err := biz.storeForgotPassword.FindUser(ctx, map[string]interface{}{"email": email})

	if err != nil {
		return usermodel.ErrUsernameOrPasswordInvalid
	}

	if user.Status == 0 {
		return common.ErrEntityDeleted(usermodel.EntityName, err)
	}

	if user.IsVerified != 1 {
		return usermodel.ErrEmailNotVerified
	}
	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return common.ErrInternal(err)
	}

	go func() {
		defer common.AppRecover()
		sendEmail := asyncjob.NewJob(func(ctx context.Context) error {
			if err := biz.emailProvider.SendVerifyCode(accessToken.Token, email, common.ForgotPassword); err != nil {
				return err
			}
			return nil
		})
		_ = asyncjob.NewGroup(true, sendEmail).Run(ctx)
	}()

	return nil
}
