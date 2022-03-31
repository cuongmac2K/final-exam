package commentbiz

import (
	"context"
	"finnal-exam/modules/comment/commentmodel"
)

type CreateCommentStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*commentmodel.Comment, error)
	CreateData(ctx context.Context, data *commentmodel.CommentCreate) error
}
type createCommentBiz struct {
	store CreateCommentStore
}

func NewCreateCommentBiz(store CreateCommentStore) *createCommentBiz {
	return &createCommentBiz{store: store}
}
func (biz *createCommentBiz) CreateComment(ctx context.Context, data *commentmodel.CommentCreate) error {

	if data.ParentCmt != 0 {
		parent, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": data.ParentCmt})
		if err != nil {
			return commentmodel.ErrCannotReplyComment
		}
		if parent.Level != 0 {
			return commentmodel.ErrCannotReplyCommentPlus
		}
		data.Level = 1

	}
	if err := biz.store.CreateData(ctx, data); err != nil {
		return err
	}
	return nil
}
