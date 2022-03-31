package userfollowstorage

import (
	"context"
	"finnal-exam/modules/userfollow/userfollowmodel"
)

func (s *sqlStore) Find(ctx context.Context, id, userId int) bool {
	var data userfollowmodel.UserFollow
	if err := s.db.Table(userfollowmodel.UserFollow{}.TableName()).
		Where("id = ? and user_id = ?", id, userId).
		First(&data).Error; err != nil {
		return false
	}
	return true
}
