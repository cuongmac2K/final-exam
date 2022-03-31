package userbiz

import (
	"context"
	"errors"
	"finnal-exam/common"
	"finnal-exam/component/asyncjob"
	"finnal-exam/modules/user/usermodel"
	"strings"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
	emailProvider   common.SendEmailProvider
}

func NewRegisterBusiness(registerStorage RegisterStorage, hasher Hasher, emailProvider common.SendEmailProvider) *registerBusiness {
	return &registerBusiness{
		registerStorage: registerStorage,
		hasher:          hasher,
		emailProvider:   emailProvider,
	}
}

func (business *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := business.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if user != nil {
		return usermodel.ErrEmailExisted
	}

	data.Record = common.NewRecord()

	salt := common.GenSalt(50)
	if strings.TrimSpace(data.Email) == "" {
		panic(common.ErrInvalidRequest(errors.New("Column 'email' cannot be null\"")))
	}
	if strings.TrimSpace(data.Password) == "" {
		panic(common.ErrInvalidRequest(errors.New("Column 'password' cannot be null\"")))
	}
	if strings.TrimSpace(data.LastName) == "" {
		panic(common.ErrInvalidRequest(errors.New("Column 'last_name' cannot be null\"")))
	}
	if strings.TrimSpace(data.FirstName) == "" {
		panic(common.ErrInvalidRequest(errors.New("Column 'first_name' cannot be null\"")))
	}
	data.Password = business.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user" // hard code
	data.Status = 1

	if err := business.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(usermodel.EntityName, err)
	}
	go func() {
		defer common.AppRecover()
		sendEmail := asyncjob.NewJob(func(ctx context.Context) error {
			if err := business.emailProvider.SendVerifyCode(data.Record.GetOtp(), data.Email, common.VerifyEmail); err != nil {
				return err
			}
			return nil
		})
		_ = asyncjob.NewGroup(true, sendEmail).Run(ctx)
	}()

	return nil
}
