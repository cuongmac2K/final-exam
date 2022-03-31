package postlikebiz

import (
	"context"
	"finnal-exam/component/asyncjob"
	"finnal-exam/modules/postlike/postlikemodel"
)

type UserLikePost interface {
	Create(ctx context.Context, data *postlikemodel.PostLike) error
}

type IncreaseLikeCount interface {
	IncreaseLikeCount(ctx context.Context, id int) error
	CheckPost(ctx context.Context, id int) bool
}

type userLikePostBiz struct {
	store    UserLikePost
	incStore IncreaseLikeCount
}

func NewUserLikePostBiz(store UserLikePost, incStore IncreaseLikeCount) *userLikePostBiz {
	return &userLikePostBiz{store: store, incStore: incStore}
}

func (biz *userLikePostBiz) LikePost(
	ctx context.Context,
	data *postlikemodel.PostLike,
) error {

	if biz.incStore.CheckPost(ctx, data.Id) == false {
		return postlikemodel.ErrPostNotExist
	}
	// log.Println(data.Id)
	err := biz.store.Create(ctx, data)

	if err != nil {
		return postlikemodel.ErrCannotLikePost(err)
	}

	// side effect
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.incStore.IncreaseLikeCount(ctx, data.Id)
	})

	_ = asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}
