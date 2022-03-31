package postlikebiz

import (
	"context"
	"finnal-exam/component/asyncjob"
	"finnal-exam/modules/postlike/postlikemodel"
)

type UserUnLikePost interface {
	Delete(ctx context.Context, userId, id int) error
	CheckPostLike(ctx context.Context, userId, id int) bool
}

type DecreaseLikeCount interface {
	DecreaseLikeCount(ctx context.Context, id int) error
}

type userUnLikePostBiz struct {
	store    UserUnLikePost
	dncStore DecreaseLikeCount
}

func NewUserUnLikePostBiz(store UserUnLikePost, dncStore DecreaseLikeCount) *userUnLikePostBiz {
	return &userUnLikePostBiz{store: store, dncStore: dncStore}
}

func (biz *userUnLikePostBiz) UnLikePost(
	ctx context.Context,
	userId,
	id int,
) error {
	var err error
	if biz.store.CheckPostLike(ctx, userId, id) {

		err = biz.store.Delete(ctx, userId, id)
		if err != nil {
			return postlikemodel.ErrCannotUnlikePost(err)
		}

	} else {
		return postlikemodel.ErrDoesNotExistPostLike(err)
	}

	// side effect
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.dncStore.DecreaseLikeCount(ctx, id)
	})

	_ = asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}
