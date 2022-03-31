package userfollowbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/component/asyncjob"
	"finnal-exam/modules/userfollow/userfollowmodel"
)

type UserFollowUserStore interface {
	CreateLike(ctx context.Context, data *userfollowmodel.UserFollow) error
	Find(ctx context.Context, id, userId int) bool
}
type IncreaseFollowStore interface {
	IncreaseFollowerCount(ctx context.Context, id int) error
	IncreaseFollowingCount(ctx context.Context, id int) error
	CheckUser(ctx context.Context, id int) bool
}

type userFollowUserBiz struct {
	store    UserFollowUserStore
	incStore IncreaseFollowStore
}

func NewUserFollowUserBiz(store UserFollowUserStore,
	incStore IncreaseFollowStore) *userFollowUserBiz {
	return &userFollowUserBiz{store: store, incStore: incStore}
}

func (biz *userFollowUserBiz) FollowUser(ctx context.Context,
	data *userfollowmodel.UserFollow) error {

	if biz.store.Find(ctx, data.Id, data.UserId) {
		return userfollowmodel.ErrFollowExisted
	}

	if data.Id == data.UserId {
		return userfollowmodel.ErrUserFollowMyself
	}

	if !biz.incStore.CheckUser(ctx, data.Id) {
		return userfollowmodel.ErrUserNotAvailable
	}

	if err := biz.store.CreateLike(ctx, data); err != nil {
		return err
	}

	go func() {
		defer common.AppRecover()
		job1 := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseFollowingCount(ctx, data.UserId)
		})
		job2 := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.incStore.IncreaseFollowerCount(ctx, data.Id)
		})

		_ = asyncjob.NewGroup(true, job1, job2).Run(ctx)
	}()
	return nil

}
