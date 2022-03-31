package userfollowbiz

import (
	"context"
	"finnal-exam/common"
	userfollowmodel "finnal-exam/modules/userfollow/userfollowmodel"
)

type ListUserFollowStore interface {
	GetUsersFollow(ctx context.Context,
		conditions map[string]interface{},
		filter *userfollowmodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserFollowBiz struct {
	store ListUserFollowStore
}

func NewListUserFollowBiz(store ListUserFollowStore) *listUserFollowBiz {
	return &listUserFollowBiz{store: store}
}

func (biz *listUserFollowBiz) ListUsers(
	ctx context.Context,
	filter *userfollowmodel.Filter,
	paging *common.Paging,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersFollow(ctx, nil, filter, paging)

	if err != nil {
		return nil, common.ErrCannotListEntity(userfollowmodel.EntityName, err)
	}

	return users, nil
}
