package userfollowbiz

import (
	"context"
	"finnal-exam/common"
	userfollowmodel "finnal-exam/modules/userfollow/userfollowmodel"
)

type ListUserFollowingStore interface {
	GetUsersFollowing(ctx context.Context,
		conditions map[string]interface{},
		filter *userfollowmodel.Filter,
		paging *common.Paging,
		userid int,
		moreKeys ...string,
	) ([]common.SimpleUser, error)
}

type listUserFollowingBiz struct {
	store ListUserFollowingStore
}

func NewListUserFollowingBiz(store ListUserFollowingStore) *listUserFollowingBiz {
	return &listUserFollowingBiz{store: store}
}

func (biz *listUserFollowingBiz) ListUsersFollowing(
	ctx context.Context,
	filter *userfollowmodel.Filter,
	paging *common.Paging,
	userid int,
) ([]common.SimpleUser, error) {
	users, err := biz.store.GetUsersFollowing(ctx, nil, filter, paging, userid)

	if err != nil {
		return nil, common.ErrCannotListEntity(userfollowmodel.EntityName, err)
	}

	return users, nil
}
