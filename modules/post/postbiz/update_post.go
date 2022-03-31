package postbiz

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/post/postmodel"
)

type UpdatePostStore interface {
	FindDataByCondition(
		ctx context.Context,
		conditions map[string]interface{},
		moreKeys ...string,
	) (*postmodel.Post, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *postmodel.PostUpdate,
	) error
}

type updatePostBiz struct {
	store UpdatePostStore
}

func NewUpdatePostBiz(store UpdatePostStore) *updatePostBiz {
	return &updatePostBiz{store: store}
}

func (biz *updatePostBiz) UpdatePost(ctx context.Context, id int, data *postmodel.PostUpdate, userid int) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if oldData.UserId != userid {
		return postmodel.ErrNoUpdatePermission
	}

	if err != nil {
		return common.ErrCannotGetEntity(postmodel.EntityName, err)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(postmodel.EntityName, nil)
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return common.ErrCannotUpdateEntity(postmodel.EntityName, err)
	}

	return nil
}
