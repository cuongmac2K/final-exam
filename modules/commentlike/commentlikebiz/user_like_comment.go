package commentlikebiz

import (
	"context"
	"finnal-exam/component/asyncjob"
	"finnal-exam/modules/commentlike/commentlikemodel"
)

type UserLikeComment interface {
	Create(ctx context.Context, data *commentlikemodel.CommentLike) error
	CheckComment(ctx context.Context, id int) bool
}

type IncreaseLikeCount interface {
	IncreaseLikeCount(ctx context.Context, id int) error
}

type userLikeCommentBiz struct {
	store    UserLikeComment
	incStore IncreaseLikeCount
}

func NewUserLikeCommentBiz(store UserLikeComment, incStore IncreaseLikeCount) *userLikeCommentBiz {
	return &userLikeCommentBiz{store: store, incStore: incStore}
}

func (biz *userLikeCommentBiz) LikeComment(
	ctx context.Context,
	data *commentlikemodel.CommentLike,
) error {
	var err error
	if biz.store.CheckComment(ctx, data.Id) {
		err = biz.store.Create(ctx, data)
	} else {
		return commentlikemodel.ErrCommentDoesntexist(err)
	}

	if err != nil {
		return commentlikemodel.ErrCannotLikeComment(err)
	}

	// side effect
	job := asyncjob.NewJob(func(ctx context.Context) error {
		return biz.incStore.IncreaseLikeCount(ctx, data.Id)
	})

	_ = asyncjob.NewGroup(true, job).Run(ctx)

	return nil
}
