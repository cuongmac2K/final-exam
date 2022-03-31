package postbiz

import (
	"context"
	"finnal-exam/modules/post/postmodel"
)

type CreatePostStore interface {
	CreateData(ctx context.Context, data *postmodel.PostCreate) error
}
type createPostBiz struct {
	store CreatePostStore
}

func NewCreatePostBiz(store CreatePostStore) *createPostBiz {
	return &createPostBiz{store: store}
}
func (biz *createPostBiz) CreatePost(ctx context.Context, data *postmodel.PostCreate) error {
	if err := biz.store.CreateData(ctx, data); err != nil {
		return err
	}
	return nil
}
