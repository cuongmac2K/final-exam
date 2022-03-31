package userfollowbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/component/asyncjob"
	"finnal-exam/modules/userfollow/userfollowmodel"
)

type UserUnFollowUserStore interface {
	Delete(ctx context.Context, userId, id int) error
	Find(ctx context.Context, id, userId int) bool
}
type DeIncreaseFollowStore interface {
	DeIncreaseFollowerCount(ctx context.Context, id int) error
	DeIncreaseFollowingCount(ctx context.Context, id int) error
}

type userUnFollowUserBiz struct {
	store      UserUnFollowUserStore
	deincStore DeIncreaseFollowStore
}

func NewUserUnFollowUserBiz(store UserUnFollowUserStore,
	deincStore DeIncreaseFollowStore) *userUnFollowUserBiz {
	return &userUnFollowUserBiz{store: store, deincStore: deincStore}
}

func (biz *userUnFollowUserBiz) UnFollowUser(ctx context.Context, userId,
	id int) error {
	if !biz.store.Find(ctx, id, userId) {
		return userfollowmodel.ErrUnFollowExisted
	}

	if err := biz.store.Delete(ctx, userId, id); err != nil {
		return err
	}

	go func() {
		defer common.AppRecover()
		job1 := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.deincStore.DeIncreaseFollowingCount(ctx, userId)
		})
		job2 := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.deincStore.DeIncreaseFollowerCount(ctx, id)
		})

		_ = asyncjob.NewGroup(true, job1, job2).Run(ctx)
	}()
	return nil

}
