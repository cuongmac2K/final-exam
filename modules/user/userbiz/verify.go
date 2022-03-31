package userbiz

import (
	"context"
	"finnal-exam/modules/user/usermodel"
)

type VerifyEmailStore interface {
	UpdateVerify(ctx context.Context, id int) error
	FindUser(ctx context.Context,
		conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}
type verifyEmailBiz struct {
	store VerifyEmailStore
}

func NewVerifyEmailBiz(store VerifyEmailStore) *verifyEmailBiz {
	return &verifyEmailBiz{store: store}
}
func (biz *verifyEmailBiz) VerifyEmail(ctx context.Context, email, otp string) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": email})

	if err != nil {
		return err
	}
	if user.Record.CheckExpired() {
		return usermodel.ErrOtpExpired
	}
	if user.Record.GetOtp() != otp {
		return usermodel.ErrOtpIncorrect
	}
	if err = biz.store.UpdateVerify(ctx, user.Id); err != nil {
		return err
	}
	return nil
}
