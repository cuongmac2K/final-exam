package userbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/component/asyncjob"
	"finnal-exam/modules/user/usermodel"
)

type SendOtpCodeStore interface {
	FindUser(ctx context.Context,
		conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	UpdateRecord(
		ctx context.Context,
		id int,
		record *common.Record,
	) error
}
type sendOtpBiz struct {
	store         SendOtpCodeStore
	emailProvider common.SendEmailProvider
}

func NewSendOtpBiz(store SendOtpCodeStore, emailProvider common.SendEmailProvider) *sendOtpBiz {
	return &sendOtpBiz{store: store, emailProvider: emailProvider}
}

func (biz *sendOtpBiz) SendOtpCode(ctx context.Context, email string) error {
	user, err := biz.store.FindUser(ctx, map[string]interface{}{"email": email})

	if err != nil {
		return err
	}
	user.Record = common.NewRecord()
	if err := biz.store.UpdateRecord(ctx, user.Id, user.Record); err != nil {
		return err
	}

	go func() {
		defer common.AppRecover()
		sendEmail := asyncjob.NewJob(func(ctx context.Context) error {
			if err := biz.emailProvider.SendVerifyCode(user.Record.GetOtp(), user.Email, common.VerifyEmail); err != nil {
				return err
			}
			return nil
		})
		_ = asyncjob.NewGroup(true, sendEmail).Run(ctx)
	}()

	return nil
}
