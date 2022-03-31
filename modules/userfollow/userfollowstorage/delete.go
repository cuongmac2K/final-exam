package userfollowstorage

import (
	"context"
	"finnal-exam/common"
)

func (s *sqlStore) Delete(ctx context.Context, userId, followId int) error {
	db := s.db

	if err := db.Table("user_follows").
		Where("user_id = ? and id = ?", userId, followId).
		Delete(nil).
		Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
