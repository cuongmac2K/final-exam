package commentlikebiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/component/asyncjob"
	"finnal-exam/modules/commentlike/commentlikemodel"
)

type UserUnlikeComment interface {
	DeleteLike(ctx context.Context,
		id, userId int) error
	Find(ctx context.Context, id, userId int) bool
}
type DecreaseLikeCountStore interface {
	DecreaseLikeCount(ctx context.Context,
		id int) error
}
type userUnlikeCommentBiz struct {
	store        UserUnlikeComment
	decLikeStore DecreaseLikeCountStore
}

func NewUserUnlikeCommentBiz(store UserUnlikeComment, decLikeStore DecreaseLikeCountStore) *userUnlikeCommentBiz {
	return &userUnlikeCommentBiz{store: store, decLikeStore: decLikeStore}
}
func (biz *userUnlikeCommentBiz) UnlikeComment(ctx context.Context, id, userId int) error {
	if !biz.store.Find(ctx, id, userId) {
		return commentlikemodel.ErrUnlikeComment
	}
	if err := biz.store.DeleteLike(ctx, id, userId); err != nil {
		return err
	}
	go func() {
		defer common.AppRecover()
		job := asyncjob.NewJob(func(ctx context.Context) error {
			return biz.decLikeStore.DecreaseLikeCount(ctx, id)
		})
		_ = asyncjob.NewGroup(true, job).Run(ctx)
	}()
	return nil
}
