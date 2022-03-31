package postbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/post/postmodel"
)

type DeletePostStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
	DeleteData(
		ctx context.Context,
		id int,
	) error
}

type deletePostBiz struct {
	store DeletePostStore
}

func NewDeletePostBiz(store DeletePostStore) *deletePostBiz {
	return &deletePostBiz{store: store}
}

func (biz *deletePostBiz) DeletePost(ctx context.Context, id int, userid int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if oldData.UserId != userid {
		return postmodel.ErrNoDeletePermission
	}

	if err != nil {
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	if err := biz.store.DeleteData(ctx, id); err != nil {
		return common.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}

	return nil
}
