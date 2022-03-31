package userfollowstorage

import (
	"context"
	"finnal-exam/common"
	"finnal-exam/modules/userfollow/userfollowmodel"
)

func (s *sqlStore) CreateLike(ctx context.Context,
	data *userfollowmodel.UserFollow) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
